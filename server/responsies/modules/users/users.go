package usersResponsiesModules

import "time"

type QueryUsersParams struct {
	Limit      *string `form:"limit" binding:"required"`
	Offset     *string `form:"offset" binding:"required"`
	ID         string  `form:"id"`
	Nickname   string  `form:"nickname"`
	Account    string  `form:"account"`
	Role       string  `form:"role"`
	Gender     string  `form:"gender"`
	Status     string  `form:"status"`
	CreateTime string  `form:"createTime"`
	UpdateTime string  `form:"updateTime"`
	StartTime  string  `form:"startTime"`
	EndTime    string  `form:"endTime"`
}

type CreateUserParams struct {
	Account  string `form:"account" json:"account" binding:"required"`
	Nickname string `form:"nickname" json:"nickname" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type UpdateUserParams struct {
	Nickname string `form:"nickname"`
	Avatar   string `form:"avatar"`
	Password string `form:"password"`
	Status   string `form:"status"`
}

type Page struct {
	Limit  *int `form:"limit" binding:"required"`
	Offset *int `form:"offset" binding:"required"`
}

type SingleUserResponse struct {
	ID         string     `json:"id"`
	Account    string     `json:"account"` // 账号
	Nickname   string     `json:"nickname"`
	Avatar     *string    `json:"avatar"`
	CreateTime *time.Time `json:"createTime"`
	UpdateTime *time.Time `json:"updateTime"`
	DeleteTime *time.Time `json:"deleteTime"`
	Status     *string    `json:"status"`
}