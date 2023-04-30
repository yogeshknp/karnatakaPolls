package models

import (
	"time"
)

type Constituency struct {
	ConstituencyID   int       `json:"constituencyID"`
	ConstituencyName string    `json:"ConstituencyName"`
	TotalVoters      int64     `json:"totalVoters"`
	TotalCandidates  int       `json:"TotalCandidates"`
	WinnerId         int       `json:"WinnerId"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

// 1. Voters: Name, Adharcard number, constituency, votedTo,
// 2. Constituency: Name, Number of voters, number of candidates, winner
// 3. Candidates: Name, Adharcard number, constituency, number of votes(data redundancy can happen with table voters),
// 4.
