package middleware

import (
    "errors"

    "bosh-admin/core/ctx"
    "bosh-admin/core/log"
    "bosh-admin/dao"
    "bosh-admin/service"

    "github.com/gin-gonic/gin"
)

func JWTApiAuth() gin.HandlerFunc {
    return ctx.Handler(func(c *ctx.Context) {
        jwtSvc := service.NewJWTSvc()
        // 获取access token
        accessToken, err := jwtSvc.GetAccessToken(c)
        if err != nil {
            c.UnAuthorized(err.Error())
            c.Abort()
            return
        }
        // 验证token
        claims, err := jwtSvc.TokenValidate(accessToken, service.JwtSubjectAccess, service.JwtAudienceApi)
        if err != nil {
            c.UnAuthorized(err.Error())
            c.Abort()
            return
        }
        // 查找用户
        sysUserSvc := service.NewSysUserSvc()
        user, err := sysUserSvc.GetUserById(claims.User.UserId)
        if err != nil {
            if errors.Is(err, dao.NotFound) {
                c.UnAuthorized("账号不存在")
            } else {
                log.Error(err)
                c.UnAuthorized(ctx.ServerError)
            }
            c.Abort()
            return
        }
        if user.Status == 0 {
            c.UnAuthorized("账号已被禁用")
            c.Abort()
            return
        }
        c.Set("claims", claims)
        c.Next()
    })
}
