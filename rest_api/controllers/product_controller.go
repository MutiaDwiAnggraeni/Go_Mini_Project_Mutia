// controllers/category_controller.go
package controllers

import (
	"net/http"
	"rest/config"
	"rest/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateProductHandler(c echo.Context) error {
	product := models.Product{}
	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{false, "Input tidak valid", nil})
	}

	result := config.DB.Create(&product)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{false, "Gagal menambahkan product", nil})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{true, "product berhasil ditambahkan", product})
}

func GetProductHandler(c echo.Context) error {
	id := c.Param("id")
	idNumber, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{false, "ID tidak valid", nil})
	}

	category := models.Category{}
	result := config.DB.First(&category, idNumber)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, models.BaseResponse{false, "product tidak ditemukan", nil})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{true, "product ditemukan", category})
}

func GetAllProductsHandler(c echo.Context) error {
	var categories []models.Product
	result := config.DB.Find(&categories)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{false, "Gagal mengambil data Product", nil})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{true, "Data Product berhasil diambil", categories})
}

func UpdateProductsHandler(c echo.Context) error {
	id := c.Param("id")
	idNumber, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{false, "ID tidak valid", nil})
	}

	category := models.Category{}
	result := config.DB.First(&category, idNumber)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, models.BaseResponse{false, "Product tidak ditemukan", nil})
	}

	if err := c.Bind(&category); err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{false, "Input tidak valid", nil})
	}

	result = config.DB.Save(&category)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{false, "Gagal mengupdate Product", nil})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{true, "Product berhasil diperbarui", category})
}

func DeleteProductHandler(c echo.Context) error {
	id := c.Param("id")
	idNumber, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{false, "ID tidak valid", nil})
	}

	category := models.Category{}
	result := config.DB.First(&category, idNumber)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, models.BaseResponse{false, "Product tidak ditemukan", nil})
	}

	result = config.DB.Delete(&category)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{false, "Gagal menghapus Product", nil})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{true, "Product berhasil dihapus", nil})
}
