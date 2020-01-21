package Services

import (
	"go-framework/app/Models"
)

func EmailExists(email string) bool {
	user := Models.User{}
	Models.DB.Conn.Select([]string{"id"}).Where(&Models.User{Email: email}).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
