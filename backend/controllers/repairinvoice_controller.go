package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/moominzie/user-record/ent"
	"github.com/moominzie/user-record/ent/repairinvoice"
)

type RepairinvoiceController struct {
	client *ent.Client
	router gin.IRouter
}

type Repairinvoice struct {
	Symptomid      int
	Deviceid       int
	Userid         int
	Statusrepairid int
}

// CreateRepairinvoice handles POST requests for adding repairinvoice entities
// @Summary Create repairinvoice
// @Description Create repairinvoice
// @ID create-repairinvoice
// @Accept   json
// @Produce  json
// @Param repairinvoice body ent.Repairinvoice true "Repairinvoice entity"
// @Success 200 {object} ent.Repairinvoice
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /repairinvoices [post]
func (ctl *RepairinvoiceController) CreateRepairinvoice(c *gin.Context) {
	obj := Repairinvoice{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "repairinvoice binding failed",
		})
		return
	}

	rep, err := ctl.client.Repairinvoice.
		Create().
		SetSymptomid(obj.Symptomid).
		SetDeviceid(obj.Deviceid).
		SetUserid(obj.Userid).
		SetStatusrepairid(obj.Statusrepairid).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, rep)
}

// GetRepairinvoice handles GET requests to retrieve a repairinvoice entity
// @Summary Get a repairinvoice entity by ID
// @Description get repairinvoice by ID
// @ID get-repairinvoice
// @Produce  json
// @Param id path int true "Repairinvoice ID"
// @Success 200 {object} ent.Repairinvoice
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /repairinvoices/{id} [get]
func (ctl *RepairinvoiceController) GetRepairinvoice(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	rep, err := ctl.client.Repairinvoice.
		Query().
		Where(repairinvoice.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, rep)
}

// ListRepairinvoice handles request to get a list of repairinvoice entities
// @Summary List repairinvoice entities
// @Description list repairinvoice entities
// @ID list-repairinvoice
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Repairinvoice
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /repairinvoices [get]
func (ctl *RepairinvoiceController) ListRepairinvoice(c *gin.Context) {
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

	repairinvoices, err := ctl.client.Repairinvoice.
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

	c.JSON(200, repairinvoices)
}

// DeleteRepairinvoice handles DELETE requests to delete a repairinvoice entity
// @Summary Delete a repairinvoice entity by ID
// @Description get repairinvoice by ID
// @ID delete-repairinvoice
// @Produce  json
// @Param id path int true "Repairinvoice ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /repairinvoice/{id} [delete]
func (ctl *RepairinvoiceController) DeleteRepairinvoice(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Repairinvoice.
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

// NewRepairinvoiceController creates and registers handles for the repairinvoice controller
func NewRepairinvoiceController(router gin.IRouter, client *ent.Client) *RepairinvoiceController {
	repc := &RepairinvoiceController{
		client: client,
		router: router,
	}

	repc.register()

	return repc

}

func (ctl *RepairinvoiceController) register() {
	repairinvoices := ctl.router.Group("/repairinvoices")

	repairinvoices.POST("", ctl.CreateRepairinvoice)
	repairinvoices.GET(":id", ctl.GetRepairinvoice)
	repairinvoices.GET("", ctl.ListRepairinvoice)
	repairinvoices.DELETE("", ctl.DeleteRepairinvoice)

}
