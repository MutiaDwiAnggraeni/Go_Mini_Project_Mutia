// controllers/category_controller.go
package controllers

import (
	"net/http"
	"rest/config"
	"rest/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateCategoryHandler(c echo.Context) error {
	category := models.Category{}
	if err := c.Bind(&category); err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{false, "Input tidak valid", nil})
	}

	result := config.DB.Create(&category)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{false, "Gagal menambahkan kategori", nil})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{true, "Kategori berhasil ditambahkan", category})
}

func GetCategoryHandler(c echo.Context) error {
	id := c.Param("id")
	idNumber, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{false, "ID tidak valid", nil})
	}

	category := models.Category{}
	result := config.DB.First(&category, idNumber)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, models.BaseResponse{false, "Kategori tidak ditemukan", nil})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{true, "Kategori ditemukan", category})
}

func GetAllCategoriesHandler(c echo.Context) error {
	var categories []models.Category
	result := config.DB.Find(&categories)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{false, "Gagal mengambil data kategori", nil})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{true, "Data kategori berhasil diambil", categories})
}

func UpdateCategoryHandler(c echo.Context) error {
	id := c.Param("id")
	idNumber, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{false, "ID tidak valid", nil})
	}

	category := models.Category{}
	result := config.DB.First(&category, idNumber)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, models.BaseResponse{false, "Kategori tidak ditemukan", nil})
	}

	if err := c.Bind(&category); err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{false, "Input tidak valid", nil})
	}

	result = config.DB.Save(&category)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{false, "Gagal mengupdate kategori", nil})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{true, "Kategori berhasil diperbarui", category})
}

func DeleteCategoryHandler(c echo.Context) error {
	id := c.Param("id")
	idNumber, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{false, "ID tidak valid", nil})
	}

	category := models.Category{}
	result := config.DB.First(&category, idNumber)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, models.BaseResponse{false, "Kategori tidak ditemukan", nil})
	}

	result = config.DB.Delete(&category)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{false, "Gagal menghapus kategori", nil})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{true, "Kategori berhasil dihapus", nil})
}
