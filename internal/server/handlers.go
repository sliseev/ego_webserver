package server

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vl80s/ego_webserver/internal/database"
	"gorm.io/gorm"
)

// @Summary		Create driver
// @Description	Create new driver
// @Tags			Driver
// @Accept			json
// @Produce		json
// @Param			data	body		Driver	true	"Driver data"
// @Success		200		{object}	Id
// @Failure		400		{object}	ErrorResponse
// @Failure		500		{object}	ErrorResponse
// @Router			/driver [post]
func createDriver(c *gin.Context) {
	var driver Driver
	if err := c.BindJSON(&driver); err != nil {
		c.IndentedJSON(http.StatusBadRequest, ErrorResponse{err.Error()})
		return
	}

	if err := driver.Validate(); err != nil {
		c.IndentedJSON(http.StatusBadRequest, ErrorResponse{err.Error()})
		return
	}

	rec := database.Driver{
		FullName:  driver.Name,
		LicenseId: driver.License,
	}

	db := c.MustGet("db").(*gorm.DB)
	result := db.Create(&rec)
	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, ErrorResponse{result.Error.Error()})
		return
	}

	id := strconv.FormatInt(int64(rec.ID), 10)
	c.IndentedJSON(http.StatusCreated, Id{
		Id: id,
	})
}

// @Summary		Get drivers
// @Description	Return all drivers
// @Tags			Driver
// @Produce		json
// @Success		200		{array}	Driver
// @Failure		500		{object}	ErrorResponse
// @Router			/driver [get]
func getDrivers(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var drivers []database.Driver
	result := db.Find(&drivers)

	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, ErrorResponse{result.Error.Error()})
		return
	}

	response := make([]Driver, len(drivers))
	for idx, driver := range drivers {
		response[idx].Name = driver.FullName
		response[idx].License = driver.LicenseId
	}

	c.IndentedJSON(http.StatusOK, response)
}

// @Summary		Get driver
// @Description	Return driver data
// @Tags			Driver
// @Produce		json
// @Param			id	path		int	true	"Driver ID"
// @Success		200	{object}	Driver
// @Failure		400	{object}	ErrorResponse
// @Failure		404	{object}	ErrorResponse
// @Failure		500	{object}	ErrorResponse
// @Router			/driver/{id} [get]
func getDriver(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, ErrorResponse{"wrong id value passed"})
		return
	}

	db := c.MustGet("db").(*gorm.DB)
	rec := database.Driver{}
	result := db.Take(&rec, id)

	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, ErrorResponse{result.Error.Error()})
		return
	}
	if result.RowsAffected < 1 {
		c.IndentedJSON(http.StatusNotFound, ErrorResponse{"driver not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, Driver{
		Name:    rec.FullName,
		License: rec.LicenseId,
	})
}

// @Summary		Update driver
// @Description	Update existing driver
// @Tags			Driver
// @Accept			json
// @Produce		json
// @Param			data	body		Driver	true	"Driver data"
// @Param			id	path		int	true	"Driver ID"
// @Success		200
// @Failure		400		{object}	ErrorResponse
// @Failure		404		{object}	ErrorResponse
// @Failure		500		{object}	ErrorResponse
// @Router			/driver/{id} [put]
func updateDriver(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, ErrorResponse{"wrong id value passed"})
		return
	}

	var driver Driver
	if err := c.BindJSON(&driver); err != nil {
		c.IndentedJSON(http.StatusBadRequest, ErrorResponse{err.Error()})
		return
	}

	db := c.MustGet("db").(*gorm.DB)
	rec_id := database.Driver{ID: id}
	rec_data := database.Driver{FullName: driver.Name, LicenseId: driver.License}
	result := db.Model(&rec_id).Updates(rec_data)

	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, ErrorResponse{result.Error.Error()})
		return
	}
	if result.RowsAffected < 1 {
		c.IndentedJSON(http.StatusNotFound, ErrorResponse{"driver not found"})
		return
	}

	c.Status(http.StatusOK)
}

// @Summary		Delete driver
// @Description	Delete existing driver
// @Tags			Driver
// @Produce		json
// @Param			id	path		int	true	"Driver ID"
// @Success		200
// @Failure		400		{object}	ErrorResponse
// @Failure		404		{object}	ErrorResponse
// @Failure		500		{object}	ErrorResponse
// @Router			/driver/{id} [delete]
func deleteDriver(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, ErrorResponse{"wrong id value passed"})
		return
	}

	db := c.MustGet("db").(*gorm.DB)
	result := db.Delete(&database.Driver{}, id)

	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, ErrorResponse{result.Error.Error()})
		return
	}
	if result.RowsAffected < 1 {
		c.IndentedJSON(http.StatusNotFound, ErrorResponse{"driver not found"})
		return
	}

	c.Status(http.StatusOK)
}

// @Summary		Get drivers count
// @Description	Return drivers count
// @Tags			Driver
// @Produce		json
// @Success		200		{object}	CountResponse
// @Failure		400		{object}	ErrorResponse
// @Failure		404		{object}	ErrorResponse
// @Failure		500		{object}	ErrorResponse
// @Router			/driver/count [get]
func getDriversCount(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var count int64
	db.Model(&database.Driver{}).Count(&count)

	c.IndentedJSON(http.StatusOK, CountResponse{
		Count: count,
	})
}
