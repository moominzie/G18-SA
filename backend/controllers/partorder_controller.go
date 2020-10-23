package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/moominzie/user-record/ent"
	"github.com/moominzie/user-record/ent/employee"
	"github.com/moominzie/user-record/ent/part"
	"github.com/moominzie/user-record/ent/repairinvoice"
)

type PartorderController struct {
	client *ent.Client
	router gin.IRouter
}

type Partorder struct {
	Repairinvoice int
	Employee      int
	Part          int
}

// CreatePartorder handles POST requests for adding partorder entities
// @Summary Create partorder
// @Description Create partorder
// @ID create-partorder
// @Accept   json
// @Produce  json
// @Param partorder body Partorder true "Partorder entity"
// @Success 200 {object} Partorder
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /partorders [post]
func (ctl *PartorderController) CreatePartorder(c *gin.Context) {
	obj := Partorder{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "part order binding failed",
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

	p, err := ctl.client.Part.
		Query().
		Where(part.IDEQ(int(obj.Part))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "part not found",
		})
		return
	}

	po, err := ctl.client.Partorder.
		Create().
		SetRepairinvoice(r).
		SetEmployee(em).
		SetPart(p).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, po)
}

// ListPartorder handles request to get a list of partorder entities
// @Summary List partorder entities
// @Description list partorder entities
// @ID list-partorder
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Partorder
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /partorders [get]
func (ctl *PartorderController) ListPartorder(c *gin.Context) {
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

	partorders, err := ctl.client.Partorder.
		Query().
		WithRepairinvoice().
		WithEmployee().
		WithPart().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, partorders)
}

// DeletePartorder handles DELETE requests to delete a partorder entity
// @Summary Delete a partorder entity by ID
// @Description get partorder by ID
// @ID delete-partorder
// @Produce  json
// @Param id path int true "Partorder ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /partorders/{id} [delete]
func (ctl *PartorderController) DeletePartorder(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Partorder.
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

// NewPartorderController creates and registers handles for the partorder controller
func NewPartorderController(router gin.IRouter, client *ent.Client) *PartorderController {
	pvc := &PartorderController{
		client: client,
		router: router,
	}

	pvc.register()

	return pvc

}

func (ctl *PartorderController) register() {
	partorders := ctl.router.Group("/partorders")

	partorders.POST("", ctl.CreatePartorder)
	partorders.GET("", ctl.ListPartorder)
	partorders.DELETE(":id", ctl.DeletePartorder)

}
