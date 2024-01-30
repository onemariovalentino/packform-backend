package handlers

import (
	"net/http"
	"packform-backend/src/app/orders/schemas"
	"packform-backend/src/app/orders/usecases"
	"packform-backend/src/pkg/helper"
	"time"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

type (
	orderHTTPHandler struct {
		usecase usecases.OrderUsecaseInterface
	}
)

func NewOrderHTTPHandler(usecase usecases.OrderUsecaseInterface) *orderHTTPHandler {
	return &orderHTTPHandler{usecase: usecase}
}

func (h *orderHTTPHandler) Mount(g *gin.Engine) {
	g.GET("/orders", h.GetOrders)
}

func (h *orderHTTPHandler) GetOrders(c *gin.Context) {
	reqID := requestid.Get(c)

	tz, _ := time.LoadLocation("Australia/Melbourne")

	req := schemas.OrderRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		resp := helper.NewResponse(http.StatusBadRequest, ``, err.Error(), nil)
		resp.RequestID = reqID
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	perPage := cast.ToInt(c.Query("per_page"))
	if perPage == 0 || perPage > 50 {
		perPage = 5
	}
	page := cast.ToInt(c.Query("page"))
	if page == 0 || page < 1 {
		page = 1
	}

	var start time.Time
	if req.StartDate != "" {
		st, _ := time.Parse("2006-01-02", req.StartDate)
		start = st.In(tz)
	}
	var end time.Time
	if req.EndDate != "" {
		et, _ := time.Parse("2006-01-02", req.EndDate)
		end = et.In(tz)
	}

	result, err := h.usecase.GetOrderDetails(c, req.Search, start.In(time.UTC), end.In(time.UTC), page, perPage)
	if err != nil {
		resp := helper.NewResponse(http.StatusNotFound, ``, err.Error(), nil)
		resp.RequestID = reqID
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp := helper.NewResponse(http.StatusOK, `success to get order details`, "", result)
	resp.RequestID = reqID
	c.JSON(http.StatusOK, resp)
}
