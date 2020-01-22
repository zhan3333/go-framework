package Request

type UserStoreRequest struct {
	Name     string `json:"name" from:"name" binding:"required"`
	Password string `json:"password" from:"password" binding:"required"`
	Email    string `json:"email" from:"email" binding:"required"`
}
