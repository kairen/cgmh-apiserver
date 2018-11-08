package v1

import (
	http "inwinstack/cgmh/apiserver/pkg/httpwrapper"
	"inwinstack/cgmh/apiserver/pkg/models"

	"github.com/gin-gonic/gin"
)

var formDAO = &models.Form{}

func ListForm(c *gin.Context) {
	query := &models.Query{}
	err := c.ShouldBindQuery(query)
	if err != nil {
		http.JSON(c, http.StatusBadRequest, nil, http.ErrorQueryParams)
		return
	}

	forms, err := formDAO.FindAll(query)
	if err != nil {
		http.JSON(c, http.StatusInternalServerError, nil, err)
		return
	}
	http.JSON(c, http.StatusSuccess, forms, nil)
}

func GetForm(c *gin.Context) {
	form, err := formDAO.FindByID(c.Param("id"))
	if err != nil {
		http.JSON(c, http.StatusInternalServerError, nil, err)
		return
	}
	http.JSON(c, http.StatusSuccess, form, nil)
}

func CreateForm(c *gin.Context) {
	form := &models.Form{}
	if err := c.ShouldBindJSON(form); err != nil {
		http.JSON(c, http.StatusBadRequest, nil, http.ErrorPayloadField)
		return
	}

	if err := formDAO.Insert(form); err != nil {
		http.JSON(c, http.StatusInternalServerError, nil, err)
		return
	}
	http.JSON(c, http.StatusSuccess, nil, nil)
}

func UpdateForm(c *gin.Context) {
	form := &models.Form{}
	if err := c.ShouldBindJSON(form); err != nil || form.ID == "" {
		http.JSON(c, http.StatusBadRequest, nil, http.ErrorPayloadField)
		return
	}

	if err := formDAO.Update(form); err != nil {
		http.JSON(c, http.StatusInternalServerError, nil, err)
		return
	}
	http.JSON(c, http.StatusSuccess, form, nil)
}

func DeleteForm(c *gin.Context) {
	form := &struct {
		ID string `json:"id" binding:"required"`
	}{}
	err := c.ShouldBindJSON(&form)
	if err != nil {
		http.JSON(c, http.StatusBadRequest, nil, http.ErrorPayloadField)
		return
	}

	if err := formDAO.RemoveByID(form.ID); err != nil {
		http.JSON(c, http.StatusInternalServerError, nil, err)
		return
	}
	http.JSON(c, http.StatusSuccess, nil, nil)
}
