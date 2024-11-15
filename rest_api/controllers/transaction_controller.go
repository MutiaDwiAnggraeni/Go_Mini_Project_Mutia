// controllers/transaction_controller.go
package controllers

import (
	"net/http"
	"rest/config"
	"rest/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

// CreateTransactiontHandler - Menambahkan data Transaction
func CreateTransactiontHandler(c echo.Context) error {
	transaction := models.Transaction{}
	if err := c.Bind(&transaction); err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{false, "Input tidak valid", nil})
	}

	result := config.DB.Create(&transaction)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{false, "Gagal menambahkan transaction", nil})
	}

	// Memuat data relasi User dan Product secara otomatis
	config.DB.Preload("User").Preload("Product").Preload("Product.Category").First(&transaction)

	return c.JSON(http.StatusOK, models.BaseResponse{true, "Transaction berhasil ditambahkan", transaction})
}

// GetTransactiontHandler - Mendapatkan detail Transaction berdasarkan ID
func GetTransactiontHandler(c echo.Context) error {
	id := c.Param("id")
	idNumber, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{false, "ID tidak valid", nil})
	}

	transaction := models.Transaction{}
	result := config.DB.Preload("User").Preload("Product").Preload("Product.Category").First(&transaction, idNumber)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, models.BaseResponse{false, "Transaction tidak ditemukan", nil})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{true, "Transaction ditemukan", transaction})
}

// GetAllTransactionHandler - Mendapatkan semua data Transaction
func GetAllTransactionHandler(c echo.Context) error {
	var transactions []models.Transaction
	result := config.DB.Preload("User").Preload("Product").Preload("Product.Category").Find(&transactions)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{false, "Gagal mengambil data Transaction", nil})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{true, "Data Transaction berhasil diambil", transactions})
}

// UpdateTransactionHandler - Mengupdate data Transaction berdasarkan ID
func UpdateTransactionHandler(c echo.Context) error {
	id := c.Param("id")
	idNumber, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{false, "ID tidak valid", nil})
	}

	transaction := models.Transaction{}
	result := config.DB.Preload("User").Preload("Product").Preload("Product.Category").First(&transaction, idNumber)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, models.BaseResponse{false, "Transaction tidak ditemukan", nil})
	}

	if err := c.Bind(&transaction); err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{false, "Input tidak valid", nil})
	}

	result = config.DB.Save(&transaction)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{false, "Gagal mengupdate Transaction", nil})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{true, "Transaction berhasil diperbarui", transaction})
}

// DeleteTransactionHandler - Menghapus data Transaction berdasarkan ID
func DeleteTransactionHandler(c echo.Context) error {
	id := c.Param("id")
	idNumber, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{false, "ID tidak valid", nil})
	}

	transaction := models.Transaction{}
	result := config.DB.First(&transaction, idNumber)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, models.BaseResponse{false, "Transaction tidak ditemukan", nil})
	}

	result = config.DB.Delete(&transaction)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{false, "Gagal menghapus Transaction", nil})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{true, "Transaction berhasil dihapus", nil})
}
