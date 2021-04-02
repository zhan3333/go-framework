package app

type RegisterReq struct {
	Name     string `json:"name" form:"name" binding:"required" faker:"first_name" example:"zhan" comment:"姓名"`
	Password string `json:"password" form:"password" binding:"required" example:"123456" comment:"密码"`
	Email    string `json:"email" form:"email" binding:"required" faker:"email" example:"admin@go-framework.com" comment:"邮箱"`
}

type LoginReq struct {
	Email    string `json:"email" form:"email" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type LoginResp struct {
	AccessToken string `json:"access_token" comment:"凭据"`
	Type        string `json:"type" comment:"凭据类型"`
	ExpiresAt   int64  `json:"expires_at" comment:"到期时间"`
}
