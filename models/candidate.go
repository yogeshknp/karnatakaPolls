package models

import (
	"time"
)

type Candidate struct {
	AadhaarID    string    `json:"aadhaarId"`
	Name         string    `json:"name"`
	Constituency string    `json:"constituency"`
	Party        string    `json:"party"`
	Votes        int64     `json:"votes"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// 1. Voters: Name, Adharcard number, constituency, votedTo,
// 2. Constituency: Name, Number of voters, number of candidates, winner
// 3. Candidates: Name, Adharcard number, constituency, number of votes(data redundancy can happen with table voters),
// 4.
