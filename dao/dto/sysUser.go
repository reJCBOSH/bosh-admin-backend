package dto

import "bosh-admin/dao"

type LoginRequest struct {
	Username  string `json:"username" form:"username" binding:"required"`   // 用户名
	Password  string `json:"password" form:"password" binding:"required"`   // 密码
	Captcha   string `json:"captcha" form:"captcha" binding:"required"`     // 验证码
	CaptchaId string `json:"captchaId" form:"captchaId" binding:"required"` // 验证码id
}

type LoginResponse struct {
	Avatar       string   `json:"avatar"`       // 头像
	Username     string   `json:"username"`     // 用户名
	Nickname     string   `json:"nickname"`     // 昵称
	PwdUpdatedAt string   `json:"pwdUpdatedAt"` // 密码更新时间
	Roles        []string `json:"roles"`        // 当前登录用户的角色
	AccessToken  string   `json:"accessToken"`  // access token
	RefreshToken string   `json:"refreshToken"` // refresh token
	Expires      int64    `json:"expires"`      // access token过期时间戳
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refreshToken" form:"refreshToken" binding:"required"` // refresh token
}

type RefreshTokenResponse struct {
	AccessToken  string `json:"accessToken"`  // access token
	RefreshToken string `json:"refreshToken"` // refresh token
	Expires      int64  `json:"expires"`      // access token过期时间戳
}

type GetUserListRequest struct {
	Pagination
	Username string `json:"username" form:"username"`                              // 用户名
	Nickname string `json:"nickname" form:"nickname"`                              // 昵称
	Gender   *int   `json:"gender" form:"gender" validate:"omitempty,oneof=0 1 2"` // 性别
	Status   *int   `json:"status" form:"status" validate:"omitempty,oneof=0 1"`   // 状态
	RoleId   *uint  `json:"roleId" form:"roleId" validate:"omitempty,gt=0"`        // 角色id
	DeptId   *uint  `json:"deptId" form:"deptId" validate:"omitempty,gt=0"`        // 部门id
}

type GetUserListResponse struct {
	List  []GetUserListItem `json:"list"`  // 用户列表
	Total int64             `json:"total"` // 总数
}

type GetUserListItem struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	Nickname string `json:"nickname"`
	Gender   int    `json:"gender"`
	Status   int    `json:"status"`
	RoleId   uint   `json:"roleId"`
	DeptId   uint   `json:"deptId"`
	Remark   string `json:"remark"`
	RoleName string `json:"roleName"`
	DeptName string `json:"deptName"`
}

type AddUserRequest struct {
	dao.AddBasicModel
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Nickname string `json:"nickname" validate:"required"`
	Gender   int    `json:"gender" validate:"oneof=0 1 2"`
	Status   int    `json:"status" validate:"oneof=0 1"`
	RoleId   uint   `json:"roleId" validate:"required,gt=0"`
	DeptId   uint   `json:"deptId" validate:"required,gt=0"`
	Remark   string `json:"remark"`
}

type EditUserRequest struct {
	dao.EditBasicModel
	Username string `json:"username" validate:"required"`
	Nickname string `json:"nickname" validate:"required"`
	Gender   int    `json:"gender" validate:"oneof=0 1 2"`
	Status   int    `json:"status" validate:"oneof=0 1"`
	RoleId   uint   `json:"roleId" validate:"required,gt=0"`
	DeptId   uint   `json:"deptId" validate:"required,gt=0"`
	Remark   string `json:"remark"`
}
