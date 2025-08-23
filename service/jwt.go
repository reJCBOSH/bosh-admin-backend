package service

import (
	"slices"
	"strings"
	"time"

	"bosh-admin/core/ctx"
	"bosh-admin/core/exception"
	"bosh-admin/dao/model"
	"bosh-admin/global"

	"github.com/duke-git/lancet/v2/random"
	"github.com/golang-jwt/jwt/v5"
)

const (
	JwtSubjectAccess  = "access"
	JwtSubjectRefresh = "refresh"
	JwtAudienceAll    = "all"
	JwtAudienceApi    = "api"
	JwtAudienceStatic = "static"
)

type JWTSvc struct {
	accessSecret    []byte // access token密钥
	refreshSecret   []byte // refresh token密钥
	accessDuration  int64  // access token有效时长
	refreshDuration int64  // refresh token有效时长
	bufferDuration  int64  // 缓冲时长
}

func NewJWTSvc() *JWTSvc {
	return &JWTSvc{
		accessSecret:    []byte(global.Config.JWT.AccessSecret),
		refreshSecret:   []byte(global.Config.JWT.RefreshSecret),
		accessDuration:  global.Config.JWT.AccessDuration,
		refreshDuration: global.Config.JWT.RefreshDuration,
		bufferDuration:  global.Config.JWT.BufferDuration,
	}
}

type CustomClaims struct {
	jwt.RegisteredClaims
	User *UserClaims
}

type UserClaims struct {
	UserId   uint   `json:"userId"`   // 用户id
	Username string `json:"username"` // 用户名
	Nickname string `json:"nickname"` // 昵称
	RoleId   uint   `json:"roleId"`   // 角色id
	RoleCode string `json:"roleCode"` // 角色标识
	DataAuth int    `json:"dataAuth"` // 数据权限
	DeptId   uint   `json:"deptId"`   // 部门id
	DeptCode string `json:"deptCode"` // 部门标识
	DeptPath string `json:"deptPath"` // 部门路径
}

func (svc *JWTSvc) GetAccessToken(c *ctx.Context) (string, error) {
	// 从请求头中获取Authorization
	headerAuth := c.Request.Header.Get("Authorization")
	if headerAuth == "" {
		return "", exception.NewException("请求头中未携带Authorization")
	}
	// 分割Authorization
	authParts := strings.Split(headerAuth, " ")
	if len(authParts) != 2 || authParts[0] != "Bearer" {
		return "", exception.NewException("请求头中Authorization格式有误")
	}
	return authParts[1], nil
}

// improveClaims 完善补充Claims
func improveClaims(claims *CustomClaims, duration int64) (*CustomClaims, int64) {
	nowTime := time.Now().Local()
	expiresAt := nowTime.Add(time.Duration(duration) * time.Second)
	claims.Issuer = global.Config.Server.Name
	claims.IssuedAt = jwt.NewNumericDate(nowTime)
	claims.NotBefore = jwt.NewNumericDate(nowTime)
	claims.ExpiresAt = jwt.NewNumericDate(expiresAt)
	return claims, expiresAt.UnixMilli()
}

func (svc *JWTSvc) GenerateToken(claims *CustomClaims) (string, int64, error) {
	var duration int64
	var secret []byte
	var err error
	switch claims.Subject {
	case JwtSubjectAccess:
		duration = svc.accessDuration
		secret = svc.accessSecret
	case JwtSubjectRefresh:
		duration = svc.refreshDuration
		secret = svc.refreshSecret
	default:
		err = exception.NewException("无效主题")
	}
	if err != nil {
		return "", 0, err
	}
	claims, expires := improveClaims(claims, duration)
	// 使用指定的签名方法签名token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 加密生成token字符串
	tokenStr, err := token.SignedString(secret)
	return tokenStr, expires, err
}

// ParseToken 解析token
func (svc *JWTSvc) ParseToken(tokenStr, subject string) (*CustomClaims, error) {
	var keyFunc jwt.Keyfunc
	var err error
	switch subject {
	case JwtSubjectAccess:
		keyFunc = func(token *jwt.Token) (any, error) {
			return svc.accessSecret, nil
		}
	case JwtSubjectRefresh:
		keyFunc = func(token *jwt.Token) (any, error) {
			return svc.refreshSecret, nil
		}
	default:
		err = exception.NewException("无效主题")
	}
	if err != nil {
		return nil, err
	}
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, keyFunc)
	if err != nil {
		return nil, err
	}
	if token != nil && token.Valid {
		if claims, ok := token.Claims.(*CustomClaims); ok {
			return claims, nil
		}
	}
	return nil, exception.NewException("无效token")
}

// TokenValidate token校验
func (svc *JWTSvc) TokenValidate(token, subject, audience string) (*CustomClaims, error) {
	// 解析token
	claims, err := svc.ParseToken(token, subject)
	if err != nil {
		return nil, err
	}
	// claims注册验证
	if claims.ID == "" {
		return nil, exception.NewException("无效token")
	}
	if claims.Issuer != global.Config.Server.Name {
		return nil, exception.NewException("无效token")
	}
	if claims.Subject != subject {
		return nil, exception.NewException("无效token")
	}
	if !(slices.Contains(claims.Audience, "all") || slices.Contains(claims.Audience, audience)) {
		return nil, exception.NewException("无效token")
	}
	return claims, nil
}

// GetClaims 获取Claims
func (svc *JWTSvc) GetClaims(c *ctx.Context) (*CustomClaims, error) {
	accessToken, err := svc.GetAccessToken(c)
	if err != nil {
		return nil, err
	}
	return svc.ParseToken(accessToken, JwtSubjectAccess)
}

// GetUserClaims 获取UserClaims
func (svc *JWTSvc) GetUserClaims(c *ctx.Context) *UserClaims {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := svc.GetClaims(c); err != nil {
			return nil
		} else {
			return cl.User
		}
	} else {
		waitUse := claims.(*CustomClaims)
		return waitUse.User
	}
}

func (svc *JWTSvc) UserLogin(user *model.SysUser) (string, string, int64, error) {
	var claims = &CustomClaims{
		User: &UserClaims{
			UserId:   user.Id,
			Username: user.Username,
			Nickname: user.Nickname,
			RoleId:   user.RoleId,
			RoleCode: user.Role.RoleCode,
			DataAuth: user.Role.DataAuth,
			DeptId:   user.DeptId,
			DeptCode: user.Dept.DeptCode,
			DeptPath: user.Dept.DeptPath,
		},
	}
	uuid, err := random.UUIdV4()
	if err != nil {
		return "", "", 0, err
	}
	claims.ID = uuid
	var audience []string
	if user.Role.RoleCode == global.SuperAdmin && user.Dept.DeptCode == global.SystemAdmin {
		audience = []string{JwtAudienceAll}
	} else {
		audience = []string{JwtAudienceApi, JwtAudienceStatic}
	}
	claims.Audience = audience
	claims.Subject = JwtSubjectAccess
	accessToken, expires, err := svc.GenerateToken(claims)
	if err != nil {
		return "", "", 0, err
	}
	claims.Subject = JwtSubjectRefresh
	refreshToken, _, err := svc.GenerateToken(claims)
	if err != nil {
		return "", "", 0, err
	}
	return accessToken, refreshToken, expires, nil
}

func (svc *JWTSvc) RefreshToken(refreshToken string) (string, string, int64, error) {
	// 验证refresh token
	claims, err := svc.TokenValidate(refreshToken, JwtSubjectRefresh, JwtAudienceApi)
	if err != nil {
		return "", "", 0, err
	}
	// refresh token 临近过期则刷新
	if claims.ExpiresAt.Unix()-time.Now().Local().Unix() <= svc.bufferDuration {
		refreshToken, _, err = svc.GenerateToken(claims)
		if err != nil {
			return "", "", 0, err
		}
	}
	// 刷新access token
	claims.Subject = JwtSubjectAccess
	accessToken, expires, err := svc.GenerateToken(claims)
	return accessToken, refreshToken, expires, err
}
