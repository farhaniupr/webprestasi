package controller

import (
	"net/http"
	"prestasi/app/models"
	"prestasi/app/utils"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func SignUp(c *gin.Context) {
	var admin models.Admin

	if err := c.Bind(&admin); err != nil {
		utils.WrapAPIError(c, err.Error(), http.StatusBadRequest)
		return
	}
	pass, err := utils.HashGenerator(admin.Password)
	if err != nil {
		utils.WrapAPIError(c, err.Error(), http.StatusBadRequest)
		return
	}
	admin.Password = pass

	flag, err := InsertNewUser(admin)

	if flag {
		utils.WrapAPIData(c, map[string]interface{}{
			"Username": admin.Username,
		}, http.StatusOK, "success")
		return
	} else {
		utils.WrapAPIError(c, err.Error(), http.StatusBadRequest)
		return
	}
}

func InsertNewUser(admin models.Admin) (bool, error) {

	if err := models.DB.Create(&admin).Error; err != nil {
		return false, errors.Errorf("invalid prepare statement :%+v\n", err)
	}
	return true, nil
}
