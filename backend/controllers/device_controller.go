package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/moominzie/user-record/ent"
	"github.com/moominzie/user-record/ent/device"
)

type DeviceController struct {
	client *ent.Client
	router gin.IRouter
}

type Device struct {
	Dname string
	id    int
}

// CreateDevice handles POST requests for adding device entities
// @Summary Create device
// @Description Create device
// @ID create-device
// @Accept   json
// @Produce  json
// @Param device body ent.Device true "Device entity"
// @Success 200 {object} ent.Device
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /devices [post]
func (ctl *DeviceController) CreateDevice(c *gin.Context) {
	obj := Device{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "device binding failed",
		})
		return
	}

	d, err := ctl.client.Device.
		Create().
		SetDname(obj.Dname).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failedd",
		})
		return
	}

	c.JSON(200, d)
}

// GetDevice handles GET requests to retrieve a device entity
// @Summary Get a device entity by ID
// @Description get device by ID
// @ID get-device
// @Produce  json
// @Param id path int true "Device ID"
// @Success 200 {object} ent.Device
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /devices/{id} [get]
func (ctl *DeviceController) GetDevice(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	d, err := ctl.client.Device.
		Query().
		Where(device.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, d)
}

// ListDevice handles request to get a list of device entities
// @Summary List device entities
// @Description list device entities
// @ID list-device
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Device
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /devices [get]
func (ctl *DeviceController) ListDevice(c *gin.Context) {
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

	devices, err := ctl.client.Device.
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

	c.JSON(200, devices)
}

// DeleteDevice handles DELETE requests to delete a device entity
// @Summary Delete a device entity by ID
// @Description get device by ID
// @ID delete-device
// @Produce  json
// @Param id path int true "Device ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /device/{id} [delete]
func (ctl *DeviceController) DeleteDevice(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Device.
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

// NewDeviceController creates and registers handles for the device controller
func NewDeviceController(router gin.IRouter, client *ent.Client) *DeviceController {
	dv := &DeviceController{
		client: client,
		router: router,
	}

	dv.register()

	return dv

}

func (ctl *DeviceController) register() {
	devices := ctl.router.Group("/devices")

	devices.POST("", ctl.CreateDevice)
	devices.GET(":id", ctl.GetDevice)
	devices.GET("", ctl.ListDevice)
	devices.DELETE("", ctl.DeleteDevice)
}
