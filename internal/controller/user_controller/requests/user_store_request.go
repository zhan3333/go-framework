package requests

type UserStoreRequest struct {
	Name     string `json:"name" form:"name" binding:"required" faker:"first_name" example:"zhan"`
	Password string `json:"password" form:"password" binding:"required" example:"123456"`
	Email    string `json:"email" form:"email" binding:"required" faker:"email" example:"admin@go-framework.com" comment:"邮箱"`
}
