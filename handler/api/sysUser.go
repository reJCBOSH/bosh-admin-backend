package api

import (
    "bosh-admin/core/ctx"
    "bosh-admin/core/log"
    "bosh-admin/dao/dto"
    "bosh-admin/service"
)

type SysUser struct {
    svc            *service.SysUserSvc
    jwtSvc         *service.JWTSvc
    loginRecordSvc *service.SysLoginRecordSvc
}

func NewSysUserHandler() *SysUser {
    return &SysUser{svc: service.NewSysUserSvc(), jwtSvc: service.NewJWTSvc(), loginRecordSvc: service.NewSysLoginRecordSvc()}
}

func (h *SysUser) Login(c *ctx.Context) {
    var req dto.LoginReq
    msg, err := c.ValidateParams(&req)
    if c.HandlerError(err, msg) {
        return
    }
    loginIP := c.ClientIP()
    userAgent := c.Request.UserAgent()
    user, err := h.svc.Login(req.Username, req.Password, req.Captcha, req.CaptchaId)
    if c.HandlerError(err) {
        if user != nil {
            if err = h.loginRecordSvc.AddLoginRecord(user.Id, req.Username, loginIP, userAgent, 0); err != nil {
                log.Error("添加登录记录失败:", err.Error())
            }
        }
        return
    }
    accessToken, refreshToken, expires, err := h.jwtSvc.UserLogin(user)
    if c.HandlerError(err) {
        if err = h.loginRecordSvc.AddLoginRecord(user.Id, req.Username, loginIP, userAgent, 0); err != nil {
            log.Error("添加登录记录失败:", err.Error())
        }
        return
    }
    if err = h.loginRecordSvc.AddLoginRecord(user.Id, req.Username, loginIP, userAgent, 1); err != nil {
        log.Error("添加登录记录失败:", err.Error())
    }
    c.SuccessWithData(dto.LoginResp{
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
    var req dto.RefreshTokenReq
    msg, err := c.ValidateParams(&req)
    if c.HandlerError(err, msg) {
        return
    }
    accessToken, refreshToken, expires, err := h.jwtSvc.RefreshToken(req.RefreshToken)
    if c.HandlerError(err) {
        return
    }
    c.SuccessWithData(dto.RefreshTokenResp{
        AccessToken:  accessToken,
        RefreshToken: refreshToken,
        Expires:      expires,
    })
}

func (h *SysUser) GetUserList(c *ctx.Context) {
    var req dto.GetUserListReq
    msg, err := c.ValidateParams(&req)
    if c.HandlerError(err, msg) {
        return
    }
    data, total, err := h.svc.GetUserList(req.Username, req.Nickname, req.Gender, req.Status, req.RoleId, req.DeptId, req.PageNo, req.PageSize)
    if c.HandlerError(err) {
        return
    }
    var list []dto.UserListItem
    for _, v := range data {
        list = append(list, dto.UserListItem{
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
            DeptCode: v.Dept.DeptCode,
            RoleName: v.Role.RoleName,
            RoleCode: v.Role.RoleCode,
        })
    }
    c.SuccessWithList(list, total)
}

func (h *SysUser) GetUserInfo(c *ctx.Context) {
    var req dto.IdReq
    msg, err := c.ValidateParams(&req)
    if c.HandlerError(err, msg) {
        return
    }
    user, err := h.svc.GetUserById(req.Id)
    if c.HandlerError(err) {
        return
    }
    data := dto.UserListItem{
        Id:       user.Id,
        Username: user.Username,
        Avatar:   user.Avatar,
        Nickname: user.Nickname,
        Gender:   user.Gender,
        Status:   user.Status,
        RoleId:   user.RoleId,
        DeptId:   user.DeptId,
        Remark:   user.Remark,
        RoleName: user.Role.RoleName,
        DeptName: user.Dept.DeptName,
    }
    c.SuccessWithData(data)
}

func (h *SysUser) AddUser(c *ctx.Context) {
    var req dto.AddUserReq
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
    var req dto.EditUserReq
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
    var req dto.IdReq
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

func (h *SysUser) ResetPassword(c *ctx.Context) {
    var req dto.IdReq
    msg, err := c.ValidateParams(&req)
    if c.HandlerError(err, msg) {
        return
    }
    userClaims := h.jwtSvc.GetUserClaims(c)
    if userClaims == nil {
        c.Fail(ctx.ServerError)
        return
    }
    err = h.svc.ResetPassword(userClaims.UserId, req.Id)
    if c.HandlerError(err) {
        return
    }
    c.Success()
}

func (h *SysUser) SetUserStatus(c *ctx.Context) {
    var req dto.SetUserStatusReq
    msg, err := c.ValidateParams(&req)
    if c.HandlerError(err, msg) {
        return
    }
    userClaims := h.jwtSvc.GetUserClaims(c)
    if userClaims == nil {
        c.Fail(ctx.ServerError)
        return
    }
    err = h.svc.SetUserStatus(userClaims.UserId, req.Id, req.Status)
    if c.HandlerError(err) {
        return
    }
    c.Success()
}

func (h *SysUser) GetSelfInfo(c *ctx.Context) {
    userClaims := h.jwtSvc.GetUserClaims(c)
    if userClaims == nil {
        c.Fail(ctx.ServerError)
        return
    }
    user, err := h.svc.GetUserById(userClaims.UserId)
    if c.HandlerError(err) {
        return
    }
    data := dto.SelfInfo{
        Id:        user.Id,
        Username:  user.Username,
        Avatar:    user.Avatar,
        Nickname:  user.Nickname,
        Gender:    user.Gender,
        Birthday:  user.Birthday.String(),
        Email:     user.Email,
        Mobile:    user.Mobile,
        Introduce: user.Introduce,
    }
    c.SuccessWithData(data)
}

func (h *SysUser) EditSelfInfo(c *ctx.Context) {
    var req dto.EditSelfInfoReq
    msg, err := c.ValidateParams(&req)
    if c.HandlerError(err, msg) {
        return
    }
    userClaims := h.jwtSvc.GetUserClaims(c)
    if userClaims == nil {
        c.Fail(ctx.ServerError)
        return
    }
    err = h.svc.EditSelfInfo(userClaims.UserId, req)
    if c.HandlerError(err) {
        return
    }
    c.Success()
}

func (h *SysUser) EditSelfPassword(c *ctx.Context) {
    var req dto.EditSelfPasswordReq
    msg, err := c.ValidateParams(&req)
    if c.HandlerError(err, msg) {
        return
    }
    userClaims := h.jwtSvc.GetUserClaims(c)
    if userClaims == nil {
        c.Fail(ctx.ServerError)
        return
    }
    err = h.svc.EditSelfPassword(userClaims.UserId, req)
    if c.HandlerError(err) {
        return
    }
    c.Success()
}
