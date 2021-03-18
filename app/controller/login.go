package controller

import (
	"fmt"
	"log"
	"net/http"
	"prestasi/app/models"
	"prestasi/app/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func Login(auth models.Auth) (bool, error, string) {
	var admin models.Admin

	if err := models.DB.Where(&models.Admin{Username: auth.Username}).First(&admin).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, errors.Errorf("Account not found"), ""
		}
	}
	err := utils.HashComparator([]byte(admin.Password), []byte(auth.Password))
	if err != nil {
		return false, errors.Errorf("Incorrect Password"), ""
	} else {
		{

			sign := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"username": auth.Username,
				//"account_number": account.AccountNumber,
					
			})

			token, err := sign.SignedString([]byte("secret"))
			if err != nil {
				return false, err, ""
			}
			return true, nil, token
		}
	}
}

func Sign(c *gin.Context) {
	var auth models.Auth
	//var admin models.Admin
	if err := c.Bind(&auth); err != nil {
		utils.WrapAPIError(c, err.Error(), http.StatusBadRequest)
		return
	}

	log.Println("LOGIN")
	fmt.Println(auth.Username)
	fmt.Println(auth.Password)

	flag, err, token := Login(auth)
	if flag {
		utils.WrapAPIData(c, map[string]interface{}{
			"token":    token,
			"username": auth.Username,
			//"ID":       account1.ID,
		}, http.StatusOK, "success")
	} else {
		utils.WrapAPIError(c, err.Error(), http.StatusBadRequest)
	}
}
