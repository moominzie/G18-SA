package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/moominzie/user-record/ent"
	"github.com/moominzie/user-record/ent/statusr"
)

type StatusRController struct {
	client *ent.Client
	router gin.IRouter
}

type StatusR struct {
	Sname string
}

// CreateStatusR handles POST requests for adding statusr entities
// @Summary Create statusr
// @Description Create statusr
// @ID create-statusr
// @Accept   json
// @Produce  json
// @Param statusr body ent.StatusR true "StatusR entity"
// @Success 200 {object} ent.StatusR
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /statusrs [post]
func (ctl *StatusRController) CreateStatusR(c *gin.Context) {
	obj := StatusR{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "statusr binding failed",
		})
		return
	}

	s, err := ctl.client.StatusR.
		Create().
		SetSname(obj.Sname).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving faileds",
		})
		return
	}

	c.JSON(200, s)
}

// GetStatusR handles GET requests to retrieve a statusr entity
// @Summary Get a statusr entity by ID
// @Description get statusr by ID
// @ID get-statusr
// @Produce  json
// @Param id path int true "StatusR ID"
// @Success 200 {object} ent.StatusR
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /statusrs/{id} [get]
func (ctl *StatusRController) GetStatusR(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	s, err := ctl.client.StatusR.
		Query().
		Where(statusr.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, s)
}

// ListStatusR handles request to get a list of statusr entities
// @Summary List statusr entities
// @Description list statusr entities
// @ID list-statusr
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.StatusR
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /statusrs [get]
func (ctl *StatusRController) ListStatusR(c *gin.Context) {
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

	statusrs, err := ctl.client.StatusR.
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

	c.JSON(200, statusrs)
}

// DeleteStatusR handles DELETE requests to delete a statusr entity
// @Summary Delete a statusr entity by ID
// @Description get statusr by ID
// @ID delete-statusr
// @Produce  json
// @Param id path int true "StatusR ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /statusr/{id} [delete]
func (ctl *StatusRController) DeleteStatusR(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.StatusR.
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

// NewStatusRController creates and registers handles for the statusr controller
func NewStatusRController(router gin.IRouter, client *ent.Client) *StatusRController {
	st := &StatusRController{
		client: client,
		router: router,
	}

	st.register()

	return st

}

func (ctl *StatusRController) register() {
	statusrs := ctl.router.Group("/statusrs")

	statusrs.POST("", ctl.CreateStatusR)
	statusrs.GET(":id", ctl.GetStatusR)
	statusrs.GET("", ctl.ListStatusR)
	statusrs.DELETE("", ctl.DeleteStatusR)

}
