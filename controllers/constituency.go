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

func CreateConstituency(c *gin.Context) {
	var constituency models.Constituency
	c.BindJSON(&constituency)
	constituencyName := constituency.ConstituencyName
	totalVoters := constituency.TotalVoters
	totalCandidates := constituency.TotalCandidates

	_, insertError := dbConnect.Model(&models.Constituency{
		ConstituencyName: constituencyName,
		TotalVoters:      totalVoters,
		TotalCandidates:  totalCandidates,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}).Insert()

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

func EditConstituency(c *gin.Context) {
	constituencyId := c.Param("constituencyId")
	var constituency models.Constituency
	c.BindJSON(&constituency)
	constituencyName := c.Param("constituencyName")

	_, err := dbConnect.Model(&models.Voter{}).Set("constituencyName = ?", constituencyName).Where("constituencyID = ?", constituencyId).Update()
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

func DeleteConstituency(c *gin.Context) {
	var constituency models.Constituency
	c.BindJSON(&constituency)

	_, err := dbConnect.Model(constituency).WherePK().Delete()
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
