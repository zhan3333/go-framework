package requests

type UserStoreRequest struct {
	Name     string `json:"name" from:"name" binding:"required" faker:"name"`
	Password string `json:"password" from:"password" binding:"required" faker:"password"`
	Email    string `json:"email" from:"email" binding:"required" faker:"email"`
}
