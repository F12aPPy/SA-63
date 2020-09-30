package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/F12aPPy/app/ent"
	"github.com/F12aPPy/app/ent/pregnant"
	"github.com/gin-gonic/gin"
)

// PregnantController defines the struct for the pregnant controller
type PregnantController struct {
	client *ent.Client
	router gin.IRouter
}

// CreatePregnant handles POST requests for adding pregnant entities
// @Summary Create pregnant
// @Description Create pregnant
// @ID create-pregnant
// @Accept   json
// @Produce  json
// @Param pregnant body ent.Pregnant true "Pregnant entity"
// @Success 200 {object} ent.Pregnant
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /pregnants [post]
func (ctl *PregnantController) CreatePregnant(c *gin.Context) {
	obj := ent.Pregnant{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "pregnant binding failed",
		})
		return
	}

	u, err := ctl.client.Pregnant.
		Create().
		SetPREGNANTAGE(obj.PREGNANTAGE).
		SetPREGNANTNAME(obj.PREGNANTNAME).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, u)
}

// GetPregnant handles GET requests to retrieve a pregnant entity
// @Summary Get a pregnant entity by ID
// @Description get pregnant by ID
// @ID get-pregnant
// @Produce  json
// @Param id path int true "Pregnant ID"
// @Success 200 {object} ent.Pregnant
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /pregnants/{id} [get]
func (ctl *PregnantController) GetPregnant(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	u, err := ctl.client.Pregnant.
		Query().
		Where(pregnant.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, u)
}

// ListPregnant handles request to get a list of pregnant entities
// @Summary List pregnant entities
// @Description list pregnant entities
// @ID list-pregnant
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Pregnant
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /pregnants [get]
func (ctl *PregnantController) ListPregnant(c *gin.Context) {
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

	pregnants, err := ctl.client.Pregnant.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, pregnants)
}

// DeletePregnant handles DELETE requests to delete a pregnant entity
// @Summary Delete a pregnant entity by ID
// @Description get pregnant by ID
// @ID delete-pregnant
// @Produce  json
// @Param id path int true "Pregnant ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /pregnants/{id} [delete]
func (ctl *PregnantController) DeletePregnant(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Pregnant.
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

// UpdatePregnant handles PUT requests to update a pregnant entity
// @Summary Update a pregnant entity by ID
// @Description update pregnant by ID
// @ID update-pregnant
// @Accept   json
// @Produce  json
// @Param id path int true "Pregnant ID"
// @Param pregnant body ent.Pregnant true "Pregnant entity"
// @Success 200 {object} ent.Pregnant
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /pregnants/{id} [put]
func (ctl *PregnantController) UpdatePregnant(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Pregnant{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "pregnant binding failed",
		})
		return
	}
	obj.ID = int(id)
	u, err := ctl.client.Pregnant.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, u)
}

// NewPregnantController creates and registers handles for the pregnant controller
func NewPregnantController(router gin.IRouter, client *ent.Client) *PregnantController {
	uc := &PregnantController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// InitPregnantController registers routes to the main engine
func (ctl *PregnantController) register() {
	pregnants := ctl.router.Group("/pregnants")

	pregnants.GET("", ctl.ListPregnant)

	// CRUD
	pregnants.POST("", ctl.CreatePregnant)
	pregnants.GET(":id", ctl.GetPregnant)
	pregnants.PUT(":id", ctl.UpdatePregnant)
	pregnants.DELETE(":id", ctl.DeletePregnant)
}
