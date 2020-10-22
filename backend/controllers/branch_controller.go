package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/moominzie/user-record/ent"
	"github.com/moominzie/user-record/ent/branch"
	"github.com/moominzie/user-record/ent/faculty"
)

type BranchController struct {
	client *ent.Client
	router gin.IRouter
}

type Branch struct {
	Brname string
	id     int
}

// CreateBranch handles POST requests for adding branch entities
// @Summary Create branch
// @Description Create branch
// @ID create-branch
// @Accept   json
// @Produce  json
// @Param branch body ent.Branch true "Branch entity"
// @Success 200 {object} ent.Branch
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /branchs [post]
func (ctl *BranchController) CreateBranch(c *gin.Context) {
	obj := Branch{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "branch binding failed",
		})
		return
	}

	f, err := ctl.client.Faculty.
		Query().
		Where(faculty.IDEQ(int(obj.id))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "faculty not found",
		})
		return
	}

	br, err := ctl.client.Branch.
		Create().
		SetBrname(obj.Brname).
		SetFaculty(f).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, br)
}

// GetBranch handles GET requests to retrieve a branch entity
// @Summary Get a branch entity by ID
// @Description get branch by ID
// @ID get-branch
// @Produce  json
// @Param id path int true "Branch ID"
// @Success 200 {object} ent.Branch
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /branchs/{id} [get]
func (ctl *BranchController) GetBranch(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	br, err := ctl.client.Branch.
		Query().
		WithFaculty().
		Where(branch.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, br)
}

// ListBranch handles request to get a list of branch entities
// @Summary List branch entities
// @Description list branch entities
// @ID list-branch
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Branch
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /branchs [get]
func (ctl *BranchController) ListBranch(c *gin.Context) {
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

	branchs, err := ctl.client.Branch.
		Query().
		WithFaculty().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, branchs)
}

// DeleteBranch handles DELETE requests to delete a branch entity
// @Summary Delete a branch entity by ID
// @Description get branch by ID
// @ID delete-branch
// @Produce  json
// @Param id path int true "Branch ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /branch/{id} [delete]
func (ctl *BranchController) DeleteBranch(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Branch.
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

// NewBranchController creates and registers handles for the branch controller
func NewBranchController(router gin.IRouter, client *ent.Client) *BranchController {
	brc := &BranchController{
		client: client,
		router: router,
	}

	brc.register()

	return brc

}

func (ctl *BranchController) register() {
	branchs := ctl.router.Group("/branchs")

	branchs.POST("", ctl.CreateBranch)
	branchs.GET(":id", ctl.GetBranch)
	branchs.GET("", ctl.ListBranch)
	branchs.DELETE("", ctl.DeleteBranch)
}
