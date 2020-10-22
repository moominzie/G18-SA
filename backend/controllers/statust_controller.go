package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/moominzie/user-record/ent"
	"github.com/moominzie/user-record/ent/statust"
)

type StatustController struct {
	client *ent.Client
	router gin.IRouter
}

type Statust struct {
	Statustname string
}

// CreateStatust handles POST requests for adding statust entities
// @Summary Create statust
// @Description Create statust
// @ID create-statust
// @Accept   json
// @Produce  json
// @Param statust body ent.Statust true "Statust entity"
// @Success 200 {object} ent.Statust
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /statusts [post]
func (ctl *StatustController) CreateStatust(c *gin.Context) {
	obj := Statust{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "statust binding failed",
		})
		return
	}

	s, err := ctl.client.Statust.
		Create().
		SetStatustname(obj.Statustname).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, s)
}

// GetStatust handles GET requests to retrieve a statust entity
// @Summary Get a statust entity by ID
// @Description get statust by ID
// @ID get-statust
// @Produce  json
// @Param id path int true "Statust ID"
// @Success 200 {object} ent.Statust
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /statusts/{id} [get]
func (ctl *StatustController) GetStatust(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	s, err := ctl.client.Statust.
		Query().
		Where(statust.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, s)
}

// ListStatust handles request to get a list of statust entities
// @Summary List statust entities
// @Description list statust entities
// @ID list-statust
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Statust
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /statusts [get]
func (ctl *StatustController) ListStatust(c *gin.Context) {
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

	statusts, err := ctl.client.Statust.
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

	c.JSON(200, statusts)
}

// DeleteStatust handles DELETE requests to delete a statust entity
// @Summary Delete a statust entity by ID
// @Description get statust by ID
// @ID delete-statust
// @Produce  json
// @Param id path int true "Statust ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /statust/{id} [delete]
func (ctl *StatustController) DeleteStatust(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Statust.
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

// NewStatustController creates and registers handles for the statust controller
func NewStatustController(router gin.IRouter, client *ent.Client) *StatustController {
	sc := &StatustController{
		client: client,
		router: router,
	}

	sc.register()

	return sc

}

func (ctl *StatustController) register() {
	statusts := ctl.router.Group("/statusts")

	statusts.POST("", ctl.CreateStatust)
	statusts.GET(":id", ctl.GetStatust)
	statusts.GET("", ctl.ListStatust)
	statusts.DELETE("", ctl.DeleteStatust)

}
