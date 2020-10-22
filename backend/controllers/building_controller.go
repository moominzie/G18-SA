package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/moominzie/user-record/ent"
	"github.com/moominzie/user-record/ent/building"
)

type BuildingController struct {
	client *ent.Client
	router gin.IRouter
}

type Building struct {
	Buname string
}

// CreateBuilding handles POST requests for adding building entities
// @Summary Create building
// @Description Create building
// @ID create-building
// @Accept   json
// @Produce  json
// @Param building body ent.Building true "Building entity"
// @Success 200 {object} ent.Building
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /buildings [post]
func (ctl *BuildingController) CreateBuilding(c *gin.Context) {
	obj := Building{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "building binding failed",
		})
		return
	}

	bu, err := ctl.client.Building.
		Create().
		SetBuname(obj.Buname).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, bu)
}

// GetBuilding handles GET requests to retrieve a building entity
// @Summary Get a building entity by ID
// @Description get building by ID
// @ID get-building
// @Produce  json
// @Param id path int true "Building ID"
// @Success 200 {object} ent.Building
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /buildings/{id} [get]
func (ctl *BuildingController) GetBuilding(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	bu, err := ctl.client.Building.
		Query().
		Where(building.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, bu)
}

// ListBuilding handles request to get a list of building entities
// @Summary List building entities
// @Description list building entities
// @ID list-building
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Building
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /buildings [get]
func (ctl *BuildingController) ListBuilding(c *gin.Context) {
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

	buildings, err := ctl.client.Building.
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

	c.JSON(200, buildings)
}

// DeleteBuilding handles DELETE requests to delete a building entity
// @Summary Delete a building entity by ID
// @Description get building by ID
// @ID delete-building
// @Produce  json
// @Param id path int true "Building ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /building/{id} [delete]
func (ctl *BuildingController) DeleteBuilding(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Building.
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

// NewBuildingController creates and registers handles for the building controller
func NewBuildingController(router gin.IRouter, client *ent.Client) *BuildingController {
	buc := &BuildingController{
		client: client,
		router: router,
	}

	buc.register()

	return buc

}

func (ctl *BuildingController) register() {
	buildings := ctl.router.Group("/buildings")

	buildings.POST("", ctl.CreateBuilding)
	buildings.GET(":id", ctl.GetBuilding)
	buildings.GET("", ctl.ListBuilding)
	buildings.DELETE("", ctl.DeleteBuilding)

}
