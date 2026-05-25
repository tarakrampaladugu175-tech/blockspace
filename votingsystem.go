package main

import "fmt"

type Vote struct {
	voteAddress map[string]bool
	candidate   map[string]int
}

// Constructor
func NewVote() *Vote {
	return &Vote{
		voteAddress: make(map[string]bool),
		candidate:   make(map[string]int),
	}
}

// Add candidate
func (v *Vote) CandidateCheck(candidates string) {

	if _, exists := v.candidate[candidates]; exists {
		fmt.Println("candidate already exists")
	} else {
		v.candidate[candidates] = 0
	}
}

// Vote checking
func (v *Vote) VoteCheck(voteaddress string, candidates string) {

	// Check if already voted
	if v.voteAddress[voteaddress] {
		fmt.Println("you have already voted")
		return
	}

	// Check if candidate exists
	if _, exists := v.candidate[candidates]; !exists {
		fmt.Println("candidate does not exist")
		return
	}

	// Add vote
	v.candidate[candidates]++
	v.voteAddress[voteaddress] = true
}

// Show results
func (v *Vote) Result() {

	maxVote := 0
	winner := ""

	fmt.Println("Voting Results:")

	for candidate, votes := range v.candidate {

		fmt.Printf("%s : %d votes\n", candidate, votes)

		if votes > maxVote {
			maxVote = votes
			winner = candidate
		}
	}

	fmt.Println("Winner:", winner)
}

func main() {

	voting := NewVote()

	voting.CandidateCheck("tarakram")
	voting.CandidateCheck("mahesh")

	voting.VoteCheck("voter1", "tarakram")
	voting.VoteCheck("voter2", "mahesh")
	voting.VoteCheck("voter1", "mahesh")
	voting.VoteCheck("voter3", "tarakram")

	voting.Result()
}
