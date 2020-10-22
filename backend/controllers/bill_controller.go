package controllers

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/moominzie/user-record/ent"
	"github.com/moominzie/user-record/ent/billingstatus"
	"github.com/moominzie/user-record/ent/employee"
	"github.com/moominzie/user-record/ent/repairinvoice"
)

type BillController struct {
	client *ent.Client
	router gin.IRouter
}

type Bill struct {
	Price         int
	Time          int
	Employee      int
	Billingstatus int
	Repairinvoice int
}

// CreateBill handles POST requests for adding bill entities
// @Summary Create bill
// @Description Create bill
// @ID create-bill
// @Accept   json
// @Produce  json
// @Param bill body Bill true "Bill entity"
// @Success 200 {object} Bill
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /bills [post]
func (ctl *BillController) CreateBill(c *gin.Context) {
	obj := Bill{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "bill binding failed",
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

	bs, err := ctl.client.Billingstatus.
		Query().
		Where(billingstatus.IDEQ(int(obj.Billingstatus))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "billingstatus not found",
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

	b, err := ctl.client.Bill.
		Create().
		SetPrice(obj.Price).
		SetTime(obj.Time).
		SetEmployee(em).
		SetRepairinvoice(r).
		SetBillingstatus(bs).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, b)
}

// ListBill handles request to get a list of bill entities
// @Summary List bill entities
// @Description list bill entities
// @ID list-bill
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Bill
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /bills [get]
func (ctl *BillController) ListBill(c *gin.Context) {
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

	bills, err := ctl.client.Bill.
		Query().
		WithEmployee().
		WithRepairinvoice().
		WithBillingstatus().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, bills)
}

// NewBillController creates and registers handles for the bill controller
func NewBillController(router gin.IRouter, client *ent.Client) *BillController {
	bic := &BillController{
		client: client,
		router: router,
	}

	bic.register()

	return bic

}

func (ctl *BillController) register() {
	bills := ctl.router.Group("/bills")

	bills.POST("", ctl.CreateBill)
	bills.GET("", ctl.ListBill)

}
