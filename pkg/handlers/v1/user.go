package v1

import (
	http "inwinstack/cgmh/apiserver/pkg/httpwrapper"
	"inwinstack/cgmh/apiserver/pkg/models"

	"github.com/gin-gonic/gin"
)

var userDAO = &models.User{}

func ListUser(c *gin.Context) {
	users, err := userDAO.FindAll()
	if err != nil {
		http.JSON(c, http.StatusInternalServerError, nil, err)
		return
	}
	http.JSON(c, http.StatusSuccess, users, nil)
}

func UpdateUser(c *gin.Context) {
	user := &models.User{}
	err := c.ShouldBindJSON(&user)
	if err != nil || user.UUID == "" {
		http.JSON(c, http.StatusBadRequest, nil, http.ErrorPayloadField)
		return
	}

	if err := userDAO.Update(user); err != nil {
		http.JSON(c, http.StatusInternalServerError, nil, err)
		return
	}
	http.JSON(c, http.StatusSuccess, user, nil)
}

func DeleteUser(c *gin.Context) {
	user := &struct {
		UUID string `json:"uuid" binding:"required"`
	}{}
	err := c.ShouldBindJSON(&user)
	if err != nil || user.UUID == "" {
		http.JSON(c, http.StatusBadRequest, nil, http.ErrorPayloadField)
		return
	}

	if err := userDAO.RemoveByUUID(user.UUID); err != nil {
		http.JSON(c, http.StatusInternalServerError, nil, err)
		return
	}
	http.JSON(c, http.StatusSuccess, nil, nil)
}
