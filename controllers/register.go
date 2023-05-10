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

// Create Voter Table
func CreateVoterTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createError := db.Model(&models.Voter{}).CreateTable(opts)
	if createError != nil {
		log.Printf("Error while creating voter table, Reason: %v\n", createError)
		return createError
	}
	log.Printf("Voter table created")
	return nil
}

// INITIALIZE DB CONNECTION (TO AVOID TOO MANY CONNECTION)
var dbConnect *pg.DB

func InitiateDB(db *pg.DB) {
	dbConnect = db
}

// GetAllVoters		godoc
// @Summary			Get All Voters
// @Description		Get All Voters from Db.
// @Produce			application/json
// @Tags			voter
// @Success			200 {object} []models.Voter
// @Router			/voters [get]
func GetAllVoters(c *gin.Context) {
	var voters []models.Voter
	err := dbConnect.Model(&voters).Select()

	if err != nil {
		log.Printf("Error while getting all voters, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, voters)
}

// CreateVoter		godoc
// @Summary			Create Voter
// @Description		Create Voter or register a new voter
// @Produce			application/json
// @Param 			voter body models.Voter true "voter details"
// @Tags			voter
// @Success			200
// @Router			/voter [post]
func CreateVoter(c *gin.Context) {
	var voter models.Voter
	c.BindJSON(&voter)
	voter.CreatedAt = time.Now()
	voter.UpdatedAt = time.Now()

	_, insertError := dbConnect.Model(&voter).Insert()

	if insertError != nil {
		log.Printf("Error while inserting new voter into db, Reason: %v\n", insertError)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Voter created Successfully",
	})
}

// GetSingleVoter		godoc
// @Summary			Get Single Voter
// @Description		Get Single Voter
// @Produce			application/json
// @Param 			AadhaarID path string true "Aadhaar ID"
// @Tags			voter
// @Success			200
// @Router			/voter/{AadhaarID} [get]
func GetSingleVoter(c *gin.Context) {
	aadhaarID := c.Param("AadhaarID")
	voter := &models.Voter{AadhaarID: aadhaarID}
	err := dbConnect.Model(voter).WherePK().Select()

	if err != nil {
		log.Printf("Error while getting a single voter, Reason: %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Voter not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Single Voter",
		"data":    voter,
	})
}

// EditVoter		godoc
// @Summary			Edit Voter
// @Description		Edit Voter
// @Produce			application/json
// @Param 			AadhaarID path string true "Aadhaar ID"
// @Param 			voter body models.Voter true "voter details"
// @Tags			voter
// @Success			200
// @Router			/voter/{AadhaarID} [put]
func EditVoter(c *gin.Context) {
	var voter models.Voter
	c.BindJSON(&voter)
	voter.AadhaarID = c.Param("aadhaarID")
	voter.UpdatedAt = time.Now()

	_, err := dbConnect.Model(&voter).WherePK().Update()
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
		"message": "Voter Edited Successfully",
	})
}

// DeleteVoter		godoc
// @Summary			Delete Voter
// @Description		Delete Voter
// @Produce			application/json
// @Param 			AadhaarID path string true "Aadhaar ID"
// @Tags			voter
// @Success			200
// @Router			/voter/{AadhaarID} [delete]
func DeleteVoter(c *gin.Context) {
	aadhaarId := c.Param("aadhaarId")
	voter := &models.Voter{AadhaarID: aadhaarId}

	_, err := dbConnect.Model(voter).WherePK().Delete()
	if err != nil {
		log.Printf("Error while deleting a single voter, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Voter deleted successfully",
	})
}
