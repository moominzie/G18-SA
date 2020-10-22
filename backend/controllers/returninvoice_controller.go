package controllers

import (
	"context"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/moominzie/user-record/ent"
	"github.com/moominzie/user-record/ent/employee"
	"github.com/moominzie/user-record/ent/repairinvoice"
	"github.com/moominzie/user-record/ent/statust"
)

// ReturninvoiceController defines the struct for the returninvoice controller
type ReturninvoiceController struct {
	client *ent.Client
	router gin.IRouter
}

type Returninvoice struct {
	Addedtime     string
	Employee      int
	Repairinvoice int
	Statust       int
}

// CreateReturninvoice handles POST requests for adding returninvoice entities
// @Summary Create returninvoice
// @Description Create returninvoice
// @ID create-returninvoice
// @Accept   json
// @Produce  json
// @Param returninvoice body Returninvoice true "Returninvoice entity"
// @Success 200 {object} Returninvoice
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /returninvoices [post]
func (ctl *ReturninvoiceController) CreateReturninvoice(c *gin.Context) {
	obj := Returninvoice{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "returninvoice binding failed",
		})
		return
	}

	em, err := ctl.client.Employee.
		Query().
		Where(employee.IDEQ(int(obj.Employee))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "employee not found",
		})
		return
	}

	r, err := ctl.client.RepairInvoice.
		Query().
		Where(repairinvoice.IDEQ(int(obj.Repairinvoice))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "repairinvoice not found",
		})
		return
	}

	s, err := ctl.client.Statust.
		Query().
		Where(statust.IDEQ(int(obj.Statust))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "statust not found",
		})
		return
	}

	time, err := time.Parse(time.RFC3339, obj.Addedtime)
	ret, err := ctl.client.Returninvoice.
		Create().
		SetAddedtime(time).
		SetEmployee(em).
		SetRepairinvoice(r).
		SetStatust(s).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, ret)
}

// ListReturninvoice handles request to get a list of returninvoice entities
// @Summary List returninvoice entities
// @Description list returninvoice entities
// @ID list-returninvoice
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Returninvoice
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /returninvoices [get]
func (ctl *ReturninvoiceController) ListReturninvoice(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {
			limit = int(limit64)
		}
	}

	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {
			offset = int(offset64)
		}
	}

	returninvoices, err := ctl.client.Returninvoice.
		Query().
		WithEmployee().
		WithRepairinvoice().
		WithStatust().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, returninvoices)
}

// NewReturninvoiceController creates and registers handles for the returninvoice controller
func NewReturninvoiceController(router gin.IRouter, client *ent.Client) *ReturninvoiceController {
	retc := &ReturninvoiceController{
		client: client,
		router: router,
	}

	retc.register()

	return retc

}

func (ctl *ReturninvoiceController) register() {
	returninvoices := ctl.router.Group("/returninvoices")

	returninvoices.POST("", ctl.CreateReturninvoice)
	returninvoices.GET("", ctl.ListReturninvoice)

}
