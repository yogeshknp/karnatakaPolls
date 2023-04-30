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

func CreateCandidateTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createError := db.Model(&models.Candidate{}).CreateTable(opts)
	if createError != nil {
		log.Printf("Error while creating Candidate table, Reason: %v\n", createError)
		return createError
	}
	log.Printf("Candidate table created")
	return nil
}

func GetAllCandidate(c *gin.Context) {
	var candidates []models.Candidate
	err := dbConnect.Model(&candidates).Select()

	if err != nil {
		log.Printf("Error while getting all Candidates, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "All Candidates",
		"data":    candidates,
	})
	return
}

func CreateCandidate(c *gin.Context) {
	var candidate models.Candidate
	c.BindJSON(&candidate)
	aadhaarId := candidate.AadhaarID
	name := candidate.Name
	constituency := candidate.Constituency
	party := candidate.Party

	_, insertError := dbConnect.Model(&models.Candidate{
		AadhaarID:    aadhaarId,
		Name:         name,
		Constituency: constituency,
		Party:        party,
		Votes:        0,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}).Insert()

	if insertError != nil {
		log.Printf("Error while inserting new candidate into db, Reason: %v\n", insertError)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "candidate created Successfully",
	})
	return
}

func EditCandidate(c *gin.Context) {
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
	return
}

func DeleteCandidate(c *gin.Context) {
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
	return
}
