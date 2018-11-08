package handler

import (
	"time"

	http "inwinstack/cgmh/apiserver/pkg/httpwrapper"
	"inwinstack/cgmh/apiserver/pkg/models"
	"inwinstack/cgmh/apiserver/pkg/util"

	"github.com/gin-gonic/gin"
)

var userDAO = &models.User{}

func Login(c *gin.Context) {
	login := struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}{}
	err := c.ShouldBindJSON(&login)
	if err != nil || login.Email == "" || login.Password == "" {
		http.JSON(c, http.StatusBadRequest, nil, http.ErrorPayloadField)
		return
	}

	if !userDAO.VerifyAccount(login.Email, util.Base64Encode(login.Password)) {
		http.JSON(c, http.StatusBadRequest, nil, http.ErrorUserLogin)
		return
	}

	user, err := userDAO.FindByEmail(login.Email)
	if err != nil {
		http.JSON(c, http.StatusInternalServerError, nil, err)
		return
	}

	if !user.Active {
		http.JSON(c, http.StatusBadRequest, nil, http.ErrorUserNotActive)
		return
	}

	token, err := util.GenerateToken(user.Email, user.UUID, 1*time.Hour)
	if err != nil {
		http.JSON(c, http.StatusInternalServerError, nil, err)
		return
	}
	http.JSON(c, http.StatusSuccess, gin.H{"token": token, "user": user}, nil)
}

func Register(c *gin.Context) {
	login := struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}{}
	err := c.ShouldBindJSON(&login)
	if err != nil || login.Email == "" || login.Password == "" {
		http.JSON(c, http.StatusBadRequest, nil, http.ErrorPayloadField)
		return
	}

	if userDAO.IsExistByEmail(login.Email) {
		http.JSON(c, http.StatusBadRequest, nil, http.ErrorUserRegister)
		return
	}

	user := &models.User{Email: login.Email}
	passwd := &models.Password{Secret: util.Base64Encode(login.Password)}
	if err := userDAO.Insert(user, passwd); err != nil {
		http.JSON(c, http.StatusInternalServerError, nil, err)
		return
	}
	http.JSON(c, http.StatusSuccess, nil, nil)
}
