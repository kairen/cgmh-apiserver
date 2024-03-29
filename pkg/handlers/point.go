package handler

import (
	http "inwinstack/cgmh/apiserver/pkg/httpwrapper"
	model "inwinstack/cgmh/apiserver/pkg/models"
	"inwinstack/cgmh/apiserver/pkg/services"

	"github.com/gin-gonic/gin"
)

type PointHandler struct {
	svc *service.DataAccess
}

func (h *PointHandler) List(c *gin.Context) {
	query := &model.Query{}
	if err := c.ShouldBindQuery(query); err != nil {
		http.BadRequest(c, http.ErrorQueryParams)
		return
	}

	if !checkUserUUID(c, h.svc, query.UserUUID) {
		return
	}

	points, err := h.svc.Point.FindAll(query)
	if err != nil {
		http.InternalServerError(c, err)
		return
	}
	http.Success(c, points)
}
