package models

import (
	"fmt"
	"hash/fnv"
	"strconv"
)

// Vote represents a single vote
type Vote struct {
	Hash      string
	Candidate int
	StudentID int
}

// HashVote calcluate if it does not exist some hash of the vote, must be repeatable
func (v *Vote) HashVote(voter *Voter) {
	h := fnv.New32a()
	h.Write([]byte(strconv.Itoa(voter.StudentID)))
	v.Hash = fmt.Sprint(h.Sum32())
}
