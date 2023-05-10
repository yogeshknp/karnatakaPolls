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

// GetAllCandidate		godoc
// @Summary			GetAllCandidate
// @Description		GetAllCandidate from Db.
// @Produce			application/json
// @Tags			Candidate
// @Success			200 {object} []models.Candidate
// @Router			/candidates [get]
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
}

// CreateCandidate		godoc
// @Summary			CreateCandidate
// @Description		CreateCandidate or register a new Candidate
// @Produce			application/json
// @Param 			Candidate body models.Candidate true "Candidate details"
// @Tags			Candidate
// @Success			200
// @Router			/candidate [post]
func CreateCandidate(c *gin.Context) {
	var candidate models.Candidate
	c.BindJSON(&candidate)
	candidate.Votes = 0
	candidate.CreatedAt = time.Now()
	candidate.UpdatedAt = time.Now()
	_, insertError := dbConnect.Model(&candidate).Insert()

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
}

// GetSingleCandidate	godoc
// @Summary			Get Single Candidate
// @Description		Get Single Candidate
// @Produce			application/json
// @Param 			AadhaarID path string true "Aadhaar ID"
// @Tags			Candidate
// @Success			200
// @Router			/candidate/{AadhaarID} [get]
func GetSingleCandidate(c *gin.Context) {
	aadhaarID := c.Param("AadhaarID")
	candidate := &models.Candidate{AadhaarID: aadhaarID}
	err := dbConnect.Model(candidate).WherePK().Select()

	if err != nil {
		log.Printf("Error while getting a single candidate, Reason: %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Candidate not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Single Candidate",
		"data":    candidate,
	})
}

// EditCandidate	godoc
// @Summary			Edit Candidate
// @Description		Edit Candidate
// @Produce			application/json
// @Param 			AadhaarID path string true "Aadhaar ID"
// @Param 			Candidate body models.Candidate true "Candidate details"
// @Tags			Candidate
// @Success			200
// @Router			/candidate/{AadhaarID} [put]
func EditCandidate(c *gin.Context) {
	var candidate models.Candidate
	c.BindJSON(&candidate)
	candidate.AadhaarID = c.Param("aadhaarID")
	candidate.UpdatedAt = time.Now()

	_, err := dbConnect.Model(&candidate).WherePK().Update()
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
		"message": "Candidate Edited Successfully",
	})
}

// DeleteCandidate	godoc
// @Summary			Delete Candidate
// @Description		Delete Candidate
// @Produce			application/json
// @Param 			AadhaarID path string true "Aadhaar ID"
// @Tags			Candidate
// @Success			200
// @Router			/candidate/{AadhaarID} [delete]
func DeleteCandidate(c *gin.Context) {
	aadhaarId := c.Param("aadhaarId")
	candidate := &models.Candidate{AadhaarID: aadhaarId}

	_, err := dbConnect.Model(candidate).WherePK().Delete()
	if err != nil {
		log.Printf("Error while deleting a single candidate, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Candidate deleted successfully",
	})
}
