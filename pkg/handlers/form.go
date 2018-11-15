package handler

import (
	"inwinstack/cgmh/apiserver/pkg/dao"
	http "inwinstack/cgmh/apiserver/pkg/httpwrapper"
	"inwinstack/cgmh/apiserver/pkg/models"

	"github.com/gin-gonic/gin"
)

type FormHandler struct {
	dao *dao.DataAccess
}

func (h *FormHandler) Get(c *gin.Context) {
	form, err := h.dao.Form.FindByID(c.Param("id"))
	if err != nil {
		http.InternalServerError(c, err)
		return
	}
	http.Success(c, form)
}

func (h *FormHandler) List(c *gin.Context) {
	query := &models.Query{}
	err := c.ShouldBindQuery(query)
	if err != nil {
		http.BadRequest(c, http.ErrorQueryParams)
		return
	}

	forms, err := h.dao.Form.FindAll(query)
	if err != nil {
		http.InternalServerError(c, err)
		return
	}
	http.Success(c, forms)
}

func (h *FormHandler) Create(c *gin.Context) {
	form := &models.Form{}
	if err := c.ShouldBindJSON(form); err != nil {
		http.BadRequest(c, http.ErrorPayloadField)
		return
	}

	if err := h.dao.Form.Insert(form); err != nil {
		http.InternalServerError(c, err)
		return
	}
	http.Success(c, nil)
}

func (h *FormHandler) Update(c *gin.Context) {
	form := &models.Form{}
	if err := c.ShouldBindJSON(form); err != nil || form.ID == "" {
		http.BadRequest(c, http.ErrorPayloadField)
		return
	}

	if err := h.dao.Form.Update(form); err != nil {
		http.InternalServerError(c, err)
		return
	}
	http.Success(c, form)
}

func (h *FormHandler) Delete(c *gin.Context) {
	form := &struct {
		ID string `json:"id" binding:"required"`
	}{}
	err := c.ShouldBindJSON(&form)
	if err != nil {
		http.BadRequest(c, http.ErrorPayloadField)
		return
	}

	if err := h.dao.Form.RemoveByID(form.ID); err != nil {
		http.InternalServerError(c, err)
		return
	}
	http.Success(c, nil)
}
