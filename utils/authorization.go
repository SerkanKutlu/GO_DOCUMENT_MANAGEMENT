package utils

import (
	customerror "documentService/customError"
	"github.com/golang-jwt/jwt"
)

func Authorize(authUser *jwt.Token, roles ...string) *customerror.CustomError {
	userRole := GetUserRole(authUser)

	for _, role := range roles {
		if role == userRole {
			return nil
		}
	}
	return customerror.UnForbiddenError

}

func GetUserRole(user *jwt.Token) string {
	userClaims := user.Claims.(jwt.MapClaims)
	userRole := userClaims["Role"]
	return userRole.(string)
}

func GetUserId(user *jwt.Token) string {
	userClaims := user.Claims.(jwt.MapClaims)
	userRole := userClaims["UserId"]
	return userRole.(string)
}
