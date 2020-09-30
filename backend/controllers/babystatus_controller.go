package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/F12aPPy/app/ent"
	"github.com/F12aPPy/app/ent/babystatus"
	"github.com/gin-gonic/gin"
)

// BabystatusController defines the struct for the babystatus controller
type BabystatusController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateBabystatus handles POST requests for adding babystatus entities
// @Summary Create babystatus
// @Description Create babystatus
// @ID create-babystatus
// @Accept   json
// @Produce  json
// @Param babystatus body ent.Babystatus true "Babystatus entity"
// @Success 200 {object} ent.Babystatus
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /babystatuss [post]
func (ctl *BabystatusController) CreateBabystatus(c *gin.Context) {
	obj := ent.Babystatus{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "babystatus binding failed",
		})
		return
	}

	u, err := ctl.client.Babystatus.
		Create().
		SetSTATUSBABYNAME(obj.STATUSBABYNAME).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, u)
}

// GetBabystatus handles GET requests to retrieve a babystatus entity
// @Summary Get a babystatus entity by ID
// @Description get babystatus by ID
// @ID get-babystatus
// @Produce  json
// @Param id path int true "Babystatus ID"
// @Success 200 {object} ent.Babystatus
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /babystatuss/{id} [get]
func (ctl *BabystatusController) GetBabystatus(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	u, err := ctl.client.Babystatus.
		Query().
		Where(babystatus.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, u)
}

// ListBabystatus handles request to get a list of babystatus entities
// @Summary List babystatus entities
// @Description list babystatus entities
// @ID list-babystatus
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Babystatus
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /babystatuss [get]
func (ctl *BabystatusController) ListBabystatus(c *gin.Context) {
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

	babystatuss, err := ctl.client.Babystatus.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, babystatuss)
}

// DeleteBabystatus handles DELETE requests to delete a babystatus entity
// @Summary Delete a babystatus entity by ID
// @Description get babystatus by ID
// @ID delete-babystatus
// @Produce  json
// @Param id path int true "Babystatus ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /babystatuss/{id} [delete]
func (ctl *BabystatusController) DeleteBabystatus(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Babystatus.
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

// UpdateBabystatus handles PUT requests to update a babystatus entity
// @Summary Update a babystatus entity by ID
// @Description update babystatus by ID
// @ID update-babystatus
// @Accept   json
// @Produce  json
// @Param id path int true "Babystatus ID"
// @Param babystatus body ent.Babystatus true "Babystatus entity"
// @Success 200 {object} ent.Babystatus
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /babystatuss/{id} [put]
func (ctl *BabystatusController) UpdateBabystatus(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Babystatus{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "babystatus binding failed",
		})
		return
	}
	obj.ID = int(id)
	u, err := ctl.client.Babystatus.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, u)
}

// NewBabystatusController creates and registers handles for the babystatus controller
func NewBabystatusController(router gin.IRouter, client *ent.Client) *BabystatusController {
	uc := &BabystatusController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// InitBabystatusController registers routes to the main engine
func (ctl *BabystatusController) register() {
	babystatuss := ctl.router.Group("/babystatuss")

	babystatuss.GET("", ctl.ListBabystatus)

	// CRUD
	babystatuss.POST("", ctl.CreateBabystatus)
	babystatuss.GET(":id", ctl.GetBabystatus)
	babystatuss.PUT(":id", ctl.UpdateBabystatus)
	babystatuss.DELETE(":id", ctl.DeleteBabystatus)
}
