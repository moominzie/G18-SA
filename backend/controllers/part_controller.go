package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/moominzie/user-record/ent"
	"github.com/moominzie/user-record/ent/part"
)

type PartController struct {
	client *ent.Client
	router gin.IRouter
}

type Part struct {
	Pname string
}

// CreatePart handles POST requests for adding part entities
// @Summary Create part
// @Description Create part
// @ID create-part
// @Accept   json
// @Produce  json
// @Param part body ent.Part true "Part entity"
// @Success 200 {object} ent.Part
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /part [post]
func (ctl *PartController) CreatePart(c *gin.Context) {
	obj := ent.Part{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "part binding failed",
		})
		return
	}

	p, err := ctl.client.Part.
		Create().
		SetPname(obj.Pname).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, p)
}

// GetPart handles GET requests to retrieve a part entity
// @Summary Get a part entity by ID
// @Description get part by ID
// @ID get-part
// @Produce  json
// @Param id path int true "Part ID"
// @Success 200 {object} ent.Part
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /parts/{id} [get]
func (ctl *EmployeeController) GetPart(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	p, err := ctl.client.Part.
		Query().
		Where(part.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, p)
}

// ListPart handles request to get a list of part entities
// @Summary List part entities
// @Description list part entities
// @ID list-part
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Part
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /parts [get]
func (ctl *PartController) ListPart(c *gin.Context) {
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

	parts, err := ctl.client.Part.
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

	c.JSON(200, parts)
}

// DeletePart handles DELETE requests to delete a part entity
// @Summary Delete a part entity by ID
// @Description get part by ID
// @ID delete-part
// @Produce  json
// @Param id path int true "Part ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /parts/{id} [delete]
func (ctl *PartController) DeletePart(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Part.
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

// NewPartController creates and registers handles for the part controller
func NewPartController(router gin.IRouter, client *ent.Client) *PartController {
	rc := &PartController{
		client: client,
		router: router,
	}

	rc.register()

	return rc

}

func (ctl *PartController) register() {
	parts := ctl.router.Group("/parts")

	parts.POST("", ctl.CreatePart)
	parts.DELETE(":id", ctl.DeletePart)
	parts.GET("", ctl.ListPart)

}
