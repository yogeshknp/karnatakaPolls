package models

import (
	"time"
)

type Candidate struct {
	AadhaarID    string        `pg:"aadhaarId,pk"`
	Name         string        `pg:"name"`
	Constituency *Constituency `pg:"constituency,rel:has-one,fk:constituency"`
	Party        string        `pg:"party"`
	Votes        int64         `pg:"votes"`
	CreatedAt    time.Time     `pg:"created_at"`
	UpdatedAt    time.Time     `pg:"updated_at"`
}

// 1. Voters: Name, Adharcard number, constituency, votedTo,
// 2. Constituency: Name, Number of voters, number of candidates, winner
// 3. Candidates: Name, Adharcard number, constituency, number of votes(data redundancy can happen with table voters),
// 4.
