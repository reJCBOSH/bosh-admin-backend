package api

import (
	"bosh-admin/core/ctx"
	"bosh-admin/dao/dto"
	"bosh-admin/service"
)

type SysUser struct {
	svc    *service.SysUserSvc
	jwtSvc *service.JWTSvc
}

func NewSysUserHandler() *SysUser {
	return &SysUser{svc: service.NewSysUserSvc(), jwtSvc: service.NewJWTSvc()}
}

func (h *SysUser) Login(c *ctx.Context) {
	var req dto.LoginRequest
	msg, err := c.ValidateParams(&req)
	if c.HandlerError(err, msg) {
		return
	}
	user, err := h.svc.Login(req.Username, req.Password, req.Captcha, req.CaptchaId)
	if c.HandlerError(err) {
		return
	}
	accessToken, refreshToken, expires, err := h.jwtSvc.UserLogin(user)
	if c.HandlerError(err) {
		return
	}
	c.SuccessWithData(dto.LoginResponse{
		Avatar:       user.Avatar,
		Username:     user.Username,
		Nickname:     user.Nickname,
		PwdUpdatedAt: user.PwdUpdatedAt.String(),
		Roles:        []string{user.Role.RoleCode},
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Expires:      expires,
	})
}

func (h *SysUser) RefreshToken(c *ctx.Context) {
	var req dto.RefreshTokenRequest
	msg, err := c.ValidateParams(&req)
	if c.HandlerError(err, msg) {
		return
	}
	accessToken, refreshToken, expires, err := h.jwtSvc.RefreshToken(req.RefreshToken)
	if c.HandlerError(err) {
		return
	}
	c.SuccessWithData(dto.RefreshTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Expires:      expires,
	})
}

func (h *SysUser) GetUserList(c *ctx.Context) {
	var req dto.GetUserListRequest
	msg, err := c.ValidateParams(&req)
	if c.HandlerError(err, msg) {
		return
	}
	data, total, err := h.svc.GetUserList(req.Username, req.Nickname, req.Gender, req.Status, req.RoleId, req.DeptId, req.PageNo, req.PageSize)
	if c.HandlerError(err) {
		return
	}
	var list []dto.GetUserListItem
	for _, v := range data {
		list = append(list, dto.GetUserListItem{
			Id:       v.Id,
			Username: v.Username,
			Avatar:   v.Avatar,
			Nickname: v.Nickname,
			Gender:   v.Gender,
			Status:   v.Status,
			RoleId:   v.RoleId,
			DeptId:   v.DeptId,
			Remark:   v.Remark,
			DeptName: v.Dept.DeptName,
			RoleName: v.Role.RoleName,
		})
	}
	c.SuccessWithData(dto.GetUserListResponse{
		List:  list,
		Total: total,
	})
}

func (h *SysUser) GetUserInfo(c *ctx.Context) {
	var req dto.IdRequest
	msg, err := c.ValidateParams(&req)
	if c.HandlerError(err, msg) {
		return
	}
	user, err := h.svc.GetUserById(req.Id)
	if c.HandlerError(err) {
		return
	}
	c.SuccessWithData(user)
}

func (h *SysUser) AddUser(c *ctx.Context) {
	var req dto.AddUserRequest
	msg, err := c.ValidateParams(&req)
	if c.HandlerError(err, msg) {
		return
	}
	err = h.svc.AddUser(req)
	if c.HandlerError(err) {
		return
	}
	c.Success()
}

func (h *SysUser) EditUser(c *ctx.Context) {
	var req dto.EditUserRequest
	msg, err := c.ValidateParams(&req)
	if c.HandlerError(err, msg) {
		return
	}
	err = h.svc.EditUser(req)
	if c.HandlerError(err) {
		return
	}
	c.Success()
}

func (h *SysUser) DelUser(c *ctx.Context) {
	var req dto.IdRequest
	msg, err := c.ValidateParams(&req)
	if c.HandlerError(err, msg) {
		return
	}
	userClaims := h.jwtSvc.GetUserClaims(c)
	if userClaims == nil {
		c.Fail(ctx.ServerError)
		return
	}
	err = h.svc.DelUser(userClaims.UserId, req.Id)
	if c.HandlerError(err) {
		return
	}
	c.Success()
}
