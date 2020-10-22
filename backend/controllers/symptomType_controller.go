package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/moominzie/user-record/ent"
	"github.com/moominzie/user-record/ent/symptom"
)

type SymptomController struct {
	client *ent.Client
	router gin.IRouter
}

type Symptom struct {
	Syname string
}

// CreateSymptom handles POST requests for adding symptom entities
// @Summary Create symptom
// @Description Create symptom
// @ID create-symptom
// @Accept   json
// @Produce  json
// @Param symptom body ent.Symptom true "Symptom entity"
// @Success 200 {object} ent.Symptom
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /symptoms [post]
func (ctl *SymptomController) CreateSymptom(c *gin.Context) {
	obj := Symptom{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "symptom binding failed",
		})
		return
	}

	sy, err := ctl.client.Symptom.
		Create().
		SetSyname(obj.Syname).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failedsy",
		})
		return
	}

	c.JSON(200, sy)
}

// GetSymptom handles GET requests to retrieve a symptom entity
// @Summary Get a symptom entity by ID
// @Description get symptom by ID
// @ID get-symptom
// @Produce  json
// @Param id path int true "Symptom ID"
// @Success 200 {object} ent.Symptom
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /symptoms/{id} [get]
func (ctl *SymptomController) GetSymptom(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	sy, err := ctl.client.Symptom.
		Query().
		Where(symptom.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, sy)
}

// ListSymptom handles request to get a list of symptom entities
// @Summary List symptom entities
// @Description list symptom entities
// @ID list-symptom
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Symptom
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /symptoms [get]
func (ctl *SymptomController) ListSymptom(c *gin.Context) {
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

	symptoms, err := ctl.client.Symptom.
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

	c.JSON(200, symptoms)
}

// DeleteSymptom handles DELETE requests to delete a symptom entity
// @Summary Delete a symptom entity by ID
// @Description get symptom by ID
// @ID delete-symptom
// @Produce  json
// @Param id path int true "Symptom ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /symptom/{id} [delete]
func (ctl *SymptomController) DeleteSymptom(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Symptom.
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

// NewSymptomController creates and registers handles for the symptom controller
func NewSymptomController(router gin.IRouter, client *ent.Client) *SymptomController {
	sym := &SymptomController{
		client: client,
		router: router,
	}

	sym.register()

	return sym

}

func (ctl *SymptomController) register() {
	symptoms := ctl.router.Group("/symptoms")

	symptoms.POST("", ctl.CreateSymptom)
	symptoms.GET(":id", ctl.GetSymptom)
	symptoms.GET("", ctl.ListSymptom)
	symptoms.DELETE("", ctl.DeleteSymptom)

}
