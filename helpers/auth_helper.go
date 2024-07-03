package helpers

import (
	"errors"
	"github.com/gin-gonic/gin"
)

func CheckUserRole(c *gin.Context, role string) (err error) {
	userRole := c.GetString("user_role")
	err = nil
	if userRole != role {
		err = errors.New("unauthorized to access this resource")
		return err
	}
	return err
}

func MatchUserRoleToUid(c *gin.Context, userId string) (err error) {
	userRole := c.GetString("user_role")
	uid := c.GetString("uid")
	err = nil

	if userRole == "USER" && uid != userId {
		err = errors.New("unauthorized to access this resource")
		return err
	}

	err = CheckUserRole(c, userRole)
	return err
}
