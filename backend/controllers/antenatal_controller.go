package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/F12aPPy/app/ent"
	"github.com/F12aPPy/app/ent/antenatal"
	"github.com/gin-gonic/gin"
)

// AntenatalController defines the struct for the antenatal controller
type AntenatalController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateAntenatal handles POST requests for adding antenatal entities
// @Summary Create antenatal
// @Description Create antenatal
// @ID create-antenatal
// @Accept   json
// @Produce  json
// @Param antenatal body ent.Antenatal true "Antenatal entity"
// @Success 200 {object} ent.Antenatal
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /antenatals [post]
func (ctl *AntenatalController) CreateAntenatal(c *gin.Context) {
	obj := ent.Antenatal{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "antenatal binding failed",
		})
		return
	}

	u, err := ctl.client.Antenatal.
		Create().
		SetADDEDTIME(obj.ADDEDTIME).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, u)
}

// GetAntenatal handles GET requests to retrieve a antenatal entity
// @Summary Get a antenatal entity by ID
// @Description get antenatal by ID
// @ID get-antenatal
// @Produce  json
// @Param id path int true "Antenatal ID"
// @Success 200 {object} ent.Antenatal
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /antenatals/{id} [get]
func (ctl *AntenatalController) GetAntenatal(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	u, err := ctl.client.Antenatal.
		Query().
		Where(antenatal.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, u)
}

// ListAntenatal handles request to get a list of antenatal entities
// @Summary List antenatal entities
// @Description list antenatal entities
// @ID list-antenatal
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Antenatal
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /antenatals [get]
func (ctl *AntenatalController) ListAntenatal(c *gin.Context) {
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

	antenatals, err := ctl.client.Antenatal.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, antenatals)
}

// DeleteAntenatal handles DELETE requests to delete a antenatal entity
// @Summary Delete a antenatal entity by ID
// @Description get antenatal by ID
// @ID delete-antenatal
// @Produce  json
// @Param id path int true "Antenatal ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /antenatals/{id} [delete]
func (ctl *AntenatalController) DeleteAntenatal(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Antenatal.
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

// UpdateAntenatal handles PUT requests to update a antenatal entity
// @Summary Update a antenatal entity by ID
// @Description update antenatal by ID
// @ID update-antenatal
// @Accept   json
// @Produce  json
// @Param id path int true "Antenatal ID"
// @Param antenatal body ent.Antenatal true "Antenatal entity"
// @Success 200 {object} ent.Antenatal
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /antenatals/{id} [put]
func (ctl *AntenatalController) UpdateAntenatal(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Antenatal{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "antenatal binding failed",
		})
		return
	}
	obj.ID = int(id)
	u, err := ctl.client.Antenatal.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, u)
}

// NewAntenatalController creates and registers handles for the antenatal controller
func NewAntenatalController(router gin.IRouter, client *ent.Client) *AntenatalController {
	uc := &AntenatalController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// InitAntenatalController registers routes to the main engine
func (ctl *AntenatalController) register() {
	antenatals := ctl.router.Group("/antenatals")

	antenatals.GET("", ctl.ListAntenatal)

	// CRUD
	antenatals.POST("", ctl.CreateAntenatal)
	antenatals.GET(":id", ctl.GetAntenatal)
	antenatals.PUT(":id", ctl.UpdateAntenatal)
	antenatals.DELETE(":id", ctl.DeleteAntenatal)
}
