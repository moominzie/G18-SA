package controllers

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/moominzie/user-record/ent"
	"github.com/moominzie/user-record/ent/billingstatus"
)

type BillingstatusController struct {
	client *ent.Client
	router gin.IRouter
}

type Billingstatus struct {
	Billingstatusname string
}

// CreateBillingstatus handles POST requests for adding billingstatus entities
// @Summary Create billingstatus
// @Description Create billingstatus
// @ID create-billingstatus
// @Accept   json
// @Produce  json
// @Param billingstatus body ent.Billingstatus true "Billingstatus entity"
// @Success 200 {object} ent.Billingstatus
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /billingstatuss [post]
func (ctl *BillingstatusController) CreateBillingstatus(c *gin.Context) {
	obj := Billingstatus{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "billingstatus binding failed",
		})
		return
	}

	bs, err := ctl.client.Billingstatus.
		Create().
		SetBillingstatusname(obj.Billingstatusname).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, bs)
}

// GetBillingstatus handles GET requests to retrieve a billingstatus entity
// @Summary Get a billingstatus entity by ID
// @Description get billingstatus by ID
// @ID get-billingstatus
// @Produce  json
// @Param id path int true "Billingstatus ID"
// @Success 200 {object} ent.Billingstatus
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /billingstatuss/{id} [get]
func (ctl *BillingstatusController) GetBillingstatus(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	bs, err := ctl.client.Billingstatus.
		Query().
		Where(billingstatus.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, bs)
}

// ListBillingstatus handles request to get a list of billingstatus entities
// @Summary List billingstatus entities
// @Description list billingstatus entities
// @ID list-billingstatus
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Billingstatus
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /billingstatuss [get]
func (ctl *BillingstatusController) ListBillingstatus(c *gin.Context) {
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

	billingstatuss, err := ctl.client.Billingstatus.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, billingstatuss)
}

// NewBillingstatusController creates and registers handles for the billingstatus controller
func NewBillingstatusController(router gin.IRouter, client *ent.Client) *BillingstatusController {
	bsc := &BillingstatusController{
		client: client,
		router: router,
	}

	bsc.register()

	return bsc

}

func (ctl *BillingstatusController) register() {
	billingstatuss := ctl.router.Group("/billingstatuss")

	billingstatuss.POST("", ctl.CreateBillingstatus)
	billingstatuss.GET(":id", ctl.GetBillingstatus)
	billingstatuss.GET("", ctl.ListBillingstatus)

}
