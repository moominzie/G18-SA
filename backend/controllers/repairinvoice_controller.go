package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/moominzie/user-record/ent"
	"github.com/moominzie/user-record/ent/device"
	"github.com/moominzie/user-record/ent/statusr"
	"github.com/moominzie/user-record/ent/symptom"
	"github.com/moominzie/user-record/ent/user"
)

type RepairInvoiceController struct {
	client *ent.Client
	router gin.IRouter
}

type RepairInvoice struct {
	Rename  string
	Device  int
	StatusR int
	Symptom int
	User    int
}

// CreateRepairInvoice handles POST requests for adding repairInvoice entities
// @Summary Create repairInvoice
// @Description Create repairInvoice
// @ID create-repairInvoice
// @Accept   json
// @Produce  json
// @Param repairInvoice body RepairInvoice true "RepairInvoice entity"
// @Success 200 {object} RepairInvoice
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /repairinvoices [post]
func (ctl *RepairInvoiceController) CreateRepairInvoice(c *gin.Context) {
	obj := RepairInvoice{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "repairinvoice binding failed",
		})
		return
	}

	s, err := ctl.client.StatusR.
		Query().
		Where(statusr.IDEQ(int(obj.StatusR))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "statusr not found",
		})
		return
	}

	d, err := ctl.client.Device.
		Query().
		Where(device.IDEQ(int(obj.Device))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "device not found",
		})
		return
	}

	if err != nil {
		c.JSON(400, gin.H{
			"error": "user not found",
		})
		return
	}

	sy, err := ctl.client.Symptom.
		Query().
		Where(symptom.IDEQ(int(obj.Symptom))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "symptom not found",
		})
		return
	}

	u, err := ctl.client.User.
		Query().
		Where(user.IDEQ(int(obj.User))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "user not found",
		})
		return
	}

	r, err := ctl.client.RepairInvoice.
		Create().
		SetRename(obj.Rename).
		SetStatus(s).
		SetSymptom(sy).
		SetDevice(d).
		SetUser(u).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failedr",
		})
		return
	}

	c.JSON(200, r)
}

// DeleteRepairInvoice handles DELETE requests to delete a repairInvoice entity
// @Summary Delete a repairInvoice entity by ID
// @Description get repairInvoice by ID
// @ID delete-repairInvoice
// @Produce  json
// @Param id path int true "RepairInvoice ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /repairinvoices/{id} [delete]
func (ctl *RepairInvoiceController) DeleteRepairInvoice(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.RepairInvoice.
		DeleteOneID(int(id)).
		Exec(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"result": fmt.Sprintf("ok deleted %v", id)})
}

// ListRepairInvoice handles request to get a list of repairInvoice entities
// @Summary List repairInvoice entities
// @Description list repairInvoice entities
// @ID list-repairInvoice
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.RepairInvoice
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /repairinvoices [get]
func (ctl *RepairInvoiceController) ListRepairInvoice(c *gin.Context) {
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

	repairinvoices, err := ctl.client.RepairInvoice.
		Query().
		WithDevice().
		WithSymptom().
		WithStatus().
		WithUser().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, repairinvoices)
}

// NewRepairInvoiceController creates and registers handles for the repairInvoice controller
func NewRepairInvoiceController(router gin.IRouter, client *ent.Client) *RepairInvoiceController {
	rp := &RepairInvoiceController{
		client: client,
		router: router,
	}

	rp.register()

	return rp

}

func (ctl *RepairInvoiceController) register() {
	repairinvoices := ctl.router.Group("/repairinvoices")

	repairinvoices.POST("", ctl.CreateRepairInvoice)
	repairinvoices.GET("", ctl.ListRepairInvoice)
	repairinvoices.DELETE(":id", ctl.DeleteRepairInvoice)

}
