package handler

import (
	"fmt"
	http "inwinstack/cgmh/apiserver/pkg/httpwrapper"
	"inwinstack/cgmh/apiserver/pkg/models"
	"inwinstack/cgmh/apiserver/pkg/services"
	"inwinstack/cgmh/apiserver/pkg/util"

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

	if !checkUserUUID(c, h.svc, form.UserUUID) {
		return
	}
	http.Success(c, form)
}

func (h *FormHandler) List(c *gin.Context) {
	query := &model.Query{}
	if err := c.ShouldBindQuery(query); err != nil {
		http.BadRequest(c, http.ErrorQueryParams)
		return
	}

	if !checkUserUUID(c, h.svc, query.UserUUID) {
		return
	}

	forms, err := h.svc.Form.FindAll(query)
	if err != nil {
		http.InternalServerError(c, err)
		return
	}
	http.Success(c, forms)
}

func (h *FormHandler) calculateCharge(form *model.Form, level *model.Level) int {
	ed := util.ElapsedDay(form.ExpectTime.Start, form.ExpectTime.End)
	return (form.NumberOfGPU * level.GPUPrice) + (ed * level.DayPrice)
}

func (h *FormHandler) Create(c *gin.Context) {
	form := &model.Form{}
	if err := c.ShouldBindJSON(form); err != nil {
		fmt.Println(err)
		http.BadRequest(c, http.ErrorPayloadField)
		return
	}

	user, err := h.svc.User.FindByUUID(form.UserUUID)
	if err != nil {
		http.InternalServerError(c, err)
		return
	}

	level, err := h.svc.Level.FindByName(user.Level)
	if err != nil {
		http.InternalServerError(c, err)
		return
	}

	expect := h.calculateCharge(form, level)
	form.Charge.Expect = expect
	form.Charge.Actual = expect
	form.CreationTime = util.NowTime()
	form.LastUpdateTime = util.NowTime()
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

	f, err := h.svc.Form.FindByID(form.ID.Hex())
	if err != nil {
		http.InternalServerError(c, err)
		return
	}

	if !checkUserUUID(c, h.svc, f.UserUUID) {
		return
	}

	user, err := h.svc.User.FindByUUID(form.UserUUID)
	if err != nil {
		http.InternalServerError(c, err)
		return
	}

	level, err := h.svc.Level.FindByName(user.Level)
	if err != nil {
		http.InternalServerError(c, err)
		return
	}

	form.Charge.Expect = h.calculateCharge(form, level)
	form.LastUpdateTime = util.NowTime()
	if err := h.svc.Form.Update(form); err != nil {
		http.InternalServerError(c, err)
		return
	}
	http.Success(c, form)
}

func (h *FormHandler) UpdateStatus(c *gin.Context) {
	if !checkAdmin(c, h.svc) {
		return
	}

	status := &model.FormStatus{}
	if err := c.ShouldBindJSON(&status); err != nil || !status.Validate() {
		http.BadRequest(c, http.ErrorPayloadField)
		return
	}

	if err := h.svc.Form.UpdateStatus(status); err != nil {
		http.InternalServerError(c, err)
		return
	}
	http.Success(c, status)
}

func (h *FormHandler) Delete(c *gin.Context) {
	if !checkAdmin(c, h.svc) {
		return
	}

	obj := &struct {
		ID string `json:"id" binding:"required"`
	}{}
	if err := c.ShouldBindJSON(&obj); err != nil || obj.ID == "" {
		http.BadRequest(c, http.ErrorPayloadField)
		return
	}

	if err := h.svc.Form.Remove(obj.ID); err != nil {
		http.InternalServerError(c, err)
		return
	}
	http.Success(c, nil)
}
