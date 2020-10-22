package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/moominzie/user-record/ent"
	"github.com/moominzie/user-record/ent/branch"
	"github.com/moominzie/user-record/ent/building"
	"github.com/moominzie/user-record/ent/faculty"
	"github.com/moominzie/user-record/ent/room"
)

type UserController struct {
	client *ent.Client
	router gin.IRouter
}

type User struct {
	PersonalID   string
	PersonalName string
	Faculty      int
	Branch       int
	Building     int
	Room         int
}

// CreateUser handles POST requests for adding user entities
// @Summary Create user
// @Description Create user
// @ID create-user
// @Accept   json
// @Produce  json
// @Param user body User true "User entity"
// @Success 200 {object} User
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /users [post]
func (ctl *UserController) CreateUser(c *gin.Context) {
	obj := User{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "user binding failed",
		})
		return
	}

	f, err := ctl.client.Faculty.
		Query().
		Where(faculty.IDEQ(int(obj.Faculty))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "faculty not found",
		})
		return
	}

	br, err := ctl.client.Branch.
		Query().
		Where(branch.IDEQ(int(obj.Branch))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "branch not found",
		})
		return
	}

	bu, err := ctl.client.Building.
		Query().
		Where(building.IDEQ(int(obj.Building))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "building not found",
		})
		return
	}

	r, err := ctl.client.Room.
		Query().
		Where(room.IDEQ(int(obj.Room))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "room not found",
		})
		return
	}

	u, err := ctl.client.User.
		Create().
		SetPersonalID(obj.PersonalID).
		SetPersonalName(obj.PersonalName).
		SetFaculty(f).
		SetBranch(br).
		SetBuilding(bu).
		SetRoom(r).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, u)
}

// DeleteUser handles DELETE requests to delete a user entity
// @Summary Delete a user entity by ID
// @Description get user by ID
// @ID delete-user
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /users/{id} [delete]
func (ctl *UserController) DeleteUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.User.
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

// ListUser handles request to get a list of user entities
// @Summary List user entities
// @Description list user entities
// @ID list-user
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.User
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /users [get]
func (ctl *UserController) ListUser(c *gin.Context) {
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

	users, err := ctl.client.User.
		Query().
		WithFaculty().
		WithBuilding().
		WithBranch().
		WithRoom().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, users)
}

// NewUserController creates and registers handles for the user controller
func NewUserController(router gin.IRouter, client *ent.Client) *UserController {
	pvc := &UserController{
		client: client,
		router: router,
	}

	pvc.register()

	return pvc

}

func (ctl *UserController) register() {
	users := ctl.router.Group("/users")

	users.POST("", ctl.CreateUser)
	users.GET("", ctl.ListUser)
	users.DELETE(":id", ctl.DeleteUser)

}
