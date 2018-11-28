package handler

import (
	http "inwinstack/cgmh/apiserver/pkg/httpwrapper"
	"inwinstack/cgmh/apiserver/pkg/models"
	"inwinstack/cgmh/apiserver/pkg/services"

	"github.com/gin-gonic/gin"
)

type FormHandler struct {
	svc *service.DataAccess
}

func (h *FormHandler) Get(c *gin.Context) {
	form, err := h.svc.Form.FindByID(c.Param("id"))
	if err != nil {
		http.InternalServerError(c, err)
		return
	}
	http.Success(c, form)
}

func (h *FormHandler) List(c *gin.Context) {
	query := &model.Query{}
	err := c.ShouldBindQuery(query)
	if err != nil {
		http.BadRequest(c, http.ErrorQueryParams)
		return
	}

	forms, err := h.svc.Form.FindAll(query)
	if err != nil {
		http.InternalServerError(c, err)
		return
	}
	http.Success(c, forms)
}

func (h *FormHandler) Create(c *gin.Context) {
	form := &model.Form{}
	if err := c.ShouldBindJSON(form); err != nil {
		http.BadRequest(c, http.ErrorPayloadField)
		return
	}

	if err := h.svc.Form.Insert(form); err != nil {
		http.InternalServerError(c, err)
		return
	}
	http.Success(c, nil)
}

func (h *FormHandler) Update(c *gin.Context) {
	form := &model.Form{}
	if err := c.ShouldBindJSON(form); err != nil || form.ID == "" {
		http.BadRequest(c, http.ErrorPayloadField)
		return
	}

	if err := h.svc.Form.Update(form); err != nil {
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

	if err := h.svc.Form.RemoveByID(form.ID); err != nil {
		http.InternalServerError(c, err)
		return
	}
	http.Success(c, nil)
}
