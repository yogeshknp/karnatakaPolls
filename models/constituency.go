package models

import (
	"time"
)

type Constituency struct {
	ConstituencyName string    `pg:"constituencyName,pk"`
	TotalVoters      int64     `pg:"totalVoters"`
	TotalCandidates  int       `pg:"totalCandidates"`
	WinnerId         int       `pg:"WinnerId"`
	CreatedAt        time.Time `pg:"created_at"`
	UpdatedAt        time.Time `pg:"updated_at"`
}

// 1. Voters: Name, Adharcard number, constituency, votedTo,
// 2. Constituency: Name, Number of voters, number of candidates, winner
// 3. Candidates: Name, Adharcard number, constituency, number of votes(data redundancy can happen with table voters),
//
