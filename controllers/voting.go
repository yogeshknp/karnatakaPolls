package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	models "github.com/karnatakaPolls/models"
)

// Vote		godoc
// @Summary			Edit Voter
// @Description		Edit Voter
// @Produce			application/json
// @Param 			AadhaarID path string true "Aadhaar ID of Voter"
// @Param 			AadhaarID path string true "Aadhaar ID of Candidate"
// @Tags			Voting
// @Success			200
// @Router			/vote/{VoterAadhaarID}/{CanAadhaarID} [put]
func Vote(c *gin.Context) {

	voterAadhaarID := c.Param("VoterAadhaarID")
	voter := &models.Voter{AadhaarID: voterAadhaarID}
	err := dbConnect.Model(voter).WherePK().Select()
	if err != nil {
		log.Printf("Error while finding voter, Reason: %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Voter not found",
		})
		return
	}

	canAadhaarID := c.Param("CanAadhaarID")
	candidate := &models.Candidate{AadhaarID: canAadhaarID}
	err = dbConnect.Model(candidate).WherePK().Select()
	if err != nil {
		log.Printf("Error while finding candidate, Reason: %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Candidate not found",
		})
		return
	}

	if candidate.Constituency != voter.Constituency {
		log.Printf("Error: constituencies of voter and candidate doesn't match")
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "constituencies of voter and candidate doesn't match",
		})
		return
	}

	candidate.Votes = candidate.Votes + 1
	candidate.UpdatedAt = time.Now()
	_, err1 := dbConnect.Model(&candidate).WherePK().Update()
	if err1 != nil {
		log.Printf("Error, Reason: %v\n", err1)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Voted Successfully",
	})
}
