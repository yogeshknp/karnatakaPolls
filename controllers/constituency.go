package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	orm "github.com/go-pg/pg/v10/orm"
	models "github.com/karnatakaPolls/models"
)

func CreateConstituencyTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createError := db.Model(&models.Constituency{}).CreateTable(opts)
	if createError != nil {
		log.Printf("Error while creating Constituency table, Reason: %v\n", createError)
		return createError
	}
	log.Printf("Constituency table created")
	return nil
}

// GetAllConstituency		godoc
// @Summary			Get All Constituency
// @Description		Get All Constituency from Db.
// @Produce			application/json
// @Tags			Constituency
// @Success			200 {object} []models.Constituency
// @Router			/constituencies [get]
func GetAllConstituency(c *gin.Context) {
	var constituency []models.Constituency
	err := dbConnect.Model(&constituency).Select()

	if err != nil {
		log.Printf("Error while getting all Constituency, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "All Constituency",
		"data":    constituency,
	})
}

// CreateConstituency		godoc
// @Summary			Create Constituency
// @Description		Create Constituency or register a new Constituency
// @Produce			application/json
// @Param 			constituency body models.Constituency true "Constituency details"
// @Tags			Constituency
// @Success			200
// @Router			/constituency [post]
func CreateConstituency(c *gin.Context) {
	var constituency models.Constituency
	c.BindJSON(&constituency)
	constituency.CreatedAt = time.Now()
	constituency.UpdatedAt = time.Now()

	_, insertError := dbConnect.Model(&constituency).Insert()

	if insertError != nil {
		log.Printf("Error while inserting new Constituency into db, Reason: %v\n", insertError)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Constituency created Successfully",
	})
}

// GetSingleConstituency	godoc
// @Summary					Get Single Constituency
// @Description				Get Single Constituency
// @Produce					application/json
// @Param 					AadhaarID path string true "Aadhaar ID"
// @Tags					Constituency
// @Success					200
// @Router					/constituency/{name} [get]
func GetSingleConstituency(c *gin.Context) {
	name := c.Param("name")
	constituency := &models.Constituency{ConstituencyName: name}
	err := dbConnect.Model(constituency).WherePK().Select()

	if err != nil {
		log.Printf("Error while getting a single constituency, Reason: %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Constituency not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Single Voter",
		"data":    constituency,
	})
}

// EditConstituency	godoc
// @Summary			Edit Constituency
// @Description		Edit Constituency
// @Produce			application/json
// @Param 			name path string true "Constituency Name"
// @Param 			voter body models.Constituency true "Constituency details"
// @Tags			Constituency
// @Success			200
// @Router			/constituency/{name} [put]
func EditConstituency(c *gin.Context) {
	var constituency models.Constituency
	c.BindJSON(&constituency)
	constituency.ConstituencyName = c.Param("name")

	_, err := dbConnect.Model(&constituency).WherePK().Update()
	if err != nil {
		log.Printf("Error, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Constituency Edited Successfully",
	})
}

// DeleteConstituency		godoc
// @Summary			Delete Constituency
// @Description		Delete Constituency
// @Produce			application/json
// @Param 			name path string true "Constituency Name"
// @Tags			Constituency
// @Success			200
// @Router			/constituency/{name} [delete]
func DeleteConstituency(c *gin.Context) {
	name := c.Param("name")
	constituency := &models.Constituency{ConstituencyName: name}

	_, err := dbConnect.Model(constituency).WherePK().Delete()
	if err != nil {
		log.Printf("Error while deleting a single constituency, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Constituency deleted successfully",
	})
}
