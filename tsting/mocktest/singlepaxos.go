package main

import (
	"fmt"
	"sync"
)

type Proposal struct {
	ID    int
	Value string
}

type Message struct {
	ProposalID int
	Value      string
	ResponseCh chan<- Response
}

type Response struct {
	Promised bool
	Accepted bool
	Proposal *Proposal
}

type Acceptor struct {
	id               int
	promisedID       int
	acceptedProposal *Proposal
	msgCh            chan Message
}

func NewAcceptor(id int) *Acceptor {
	return &Acceptor{
		id:    id,
		msgCh: make(chan Message),
	}
}

func (a *Acceptor) Run() {
	for msg := range a.msgCh {
		if msg.ProposalID > a.promisedID {
			a.promisedID = msg.ProposalID
			if msg.Value != "" {
				a.acceptedProposal = &Proposal{ID: msg.ProposalID, Value: msg.Value}
			}
			msg.ResponseCh <- Response{
				Promised: true,
				Accepted: msg.Value != "",
				Proposal: a.acceptedProposal,
			}
		} else {
			msg.ResponseCh <- Response{
				Promised: false,
				Accepted: false,
			}
		}
	}
}

type Proposer struct {
	id        int
	value     string
	acceptors []*Acceptor
}

func (p *Proposer) Propose() {
	var wg sync.WaitGroup
	// Prepare Phase
	p.id += 1
	prepareCount := 0
	acceptedProposal := p.value
	responseCh := make(chan Response, len(p.acceptors))

	for _, a := range p.acceptors {
		wg.Add(1)
		go func(a *Acceptor) {
			defer wg.Done()
			a.msgCh <- Message{
				ProposalID: p.id,
				ResponseCh: responseCh,
			}
		}(a)
	}
	wg.Wait()

	for range p.acceptors {
		resp := <-responseCh
		if resp.Promised {
			prepareCount++
			if resp.Proposal != nil && resp.Proposal.ID > p.id {
				acceptedProposal = resp.Proposal.Value
			}
		}
	}

	// majority, accept
	if prepareCount >= len(p.acceptors)/2+1 {
		acceptCount := 0

		for _, a := range p.acceptors {
			wg.Add(1)
			go func(a *Acceptor) {
				defer wg.Done()
				a.msgCh <- Message{
					ProposalID: p.id,
					Value:      acceptedProposal,
					ResponseCh: responseCh,
				}
			}(a)
		}
		wg.Wait()
		for range p.acceptors {
			resp := <-responseCh
			if resp.Accepted {
				acceptCount++
			}
		}

		if acceptCount >= len(p.acceptors)/2+1 {
			fmt.Println("Proposal accepted with value:", acceptedProposal)
		} else {
			fmt.Println("Proposal rejected")
		}
	}
}
func main() {
	// acceptor 3 to goroutine 3
	acceptors := []*Acceptor{
		NewAcceptor(1),
		NewAcceptor(2),
		NewAcceptor(3),
	}

	for _, a := range acceptors {
		go a.Run()
	}

	// made proposer and try consensus
	proposer := &Proposer{id: 1, value: "myValue", acceptors: acceptors}
	proposer.Propose()
}
