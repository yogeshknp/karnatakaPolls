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
// @Tags			voters
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

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "All Voters",
		"data":    voters,
	})
}

func CreateVoter(c *gin.Context) {
	var voter models.Voter
	c.BindJSON(&voter)
	name := voter.Name
	aadhaarId := voter.AadhaarID
	_, insertError := dbConnect.Model(&models.Voter{
		AadhaarID: aadhaarId,
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}).Insert()
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

func GetSingleVoter(c *gin.Context) {
	voterId := c.Param("voterId")
	voter := &models.Voter{AadhaarID: voterId}
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

func EditVoter(c *gin.Context) {
	voterId := c.Param("voterId")
	var voter models.Voter
	c.BindJSON(&voter)
	name := c.Param("name")

	_, err := dbConnect.Model(&models.Voter{}).Set("name = ?", name).Where("id = ?", voterId).Update()
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
