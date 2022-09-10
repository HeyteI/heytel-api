package controllers

import (
	"Heytel/database"
	"Heytel/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/secured/invoice

// CreateInvoice godoc
// @Summary creating invoice
// @Schemes
// @Description creates invoice
// @Tags invoice
// @Accept json
// @Produce json
// @Success 200 {json} {invoice_data}
// @Router /api/secured/invoice/ [post]
func CreateInvoice(ctx *gin.Context) {
	var invoice models.Invoice
	if err := ctx.ShouldBindJSON(&invoice); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	db, err := database.GetDatabaseConnection()
	if err != nil {
		ctx.Abort()
		panic(err)
	}

	record := db.Create(&invoice)
	if record.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"data":  invoice,
	})
}

// GetInvoice godoc
// @Summary get invoice by id
// @Schemes
// @Description gets invoice by id
// @Tags invoice
// @Accept json
// @Produce json
// @Success 200 {json} {invoice_data}
// @Router /api/secured/invoice/ [get]
func GetInvoice(ctx *gin.Context) {

	db, err := database.GetDatabaseConnection()
	if err != nil {
		ctx.Abort()
		panic(err)
	}

	var invoice models.Invoice

	log.Print(ctx.Param("id"))
	if result := db.Where("id = ?", ctx.Param("id")).First(&invoice); result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": result.Error.Error(),
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": invoice,
	})
}

// GetInvoiceByRoom godoc
// @Summary get invoice by room_id
// @Schemes
// @Description gets invoice by room_id
// @Tags invoice
// @Accept json
// @Produce json
// @Success 200 {json} {invoice_data}
// @Router /api/secured/invoice/ [get]
func GetInvoiceByRoom(ctx *gin.Context) {

	db, err := database.GetDatabaseConnection()
	if err != nil {
		ctx.Abort()
		panic(err)
	}

	var invoice models.Invoice

	if result := db.Where("room_id = ?", ctx.Param("id")).First(&invoice); result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": result.Error.Error(),
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": invoice,
	})
}

// AllInvoices godoc
// @Summary get all invoices
// @Schemes
// @Description get all saved invoices
// @Tags invoice
// @Accept json
// @Produce json
// @Success 200 {json} {list of invoice_data}
// @Router /api/secured/invoice/all [get]
func AllInvoices(ctx *gin.Context) {

	db, err := database.GetDatabaseConnection()
	if err != nil {
		ctx.Abort()
		panic(err)
	}

	var invoices []models.Invoice

	if result := db.Find(&invoices); result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": result.Error.Error(),
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": invoices,
	})
}

type invoiceUpdate struct {
	Date        string `json:"date_range"`
	Cancelled   string `json:"cancelled"`
	Paid        string `json:"paid"`
	PeopleCount string `json:"people_count"`
}

// UpdateInvoice godoc
// @Summary update invoice
// @Schemes
// @Description update invoice data by it's id
// @Tags invoice
// @Accept json
// @Produce json
// @Success 200 {json} {invoice_data}
// @Router /api/secured/invoice/ [patch]
func UpdateInvoice(ctx *gin.Context) {
	db, err := database.GetDatabaseConnection()
	if err != nil {
		ctx.Abort()
		panic(err)
	}

	var invoice models.Invoice

	if err := db.Where("id = ?", ctx.Param("id")).First(&invoice).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input invoiceUpdate
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Model(&invoice).Updates(input)

	ctx.JSON(http.StatusOK, gin.H{"data": invoice})
}

func DeleteInvoice(ctx *gin.Context) {
	db, err := database.GetDatabaseConnection()
	if err != nil {
		ctx.Abort()
		panic(err)
	}

	var invoice models.Invoice

	log.Print(ctx.Param("id"))
	if result := db.Where("id = ?", ctx.Param("id")).First(&invoice); result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": result.Error.Error(),
		})
		ctx.Abort()
		return
	}

	db.Delete(&invoice)

	ctx.JSON(http.StatusOK, gin.H{
		"data": invoice,
	})
}