package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

// Define Microservice's JSON Request API
type voteRequest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Define Microservice's JSON Response
type voteResponse struct {
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}

// VoteService Exposes the CanVote function to transport layer
type VoteService interface {
	CanVote(string, int) (string, error)
}

// The actual service where we implement the interface
type voteService struct{}

func main() {
	svc := voteService{}

	votingHandler := httptransport.NewServer(
		makeVotingEndpoint(svc),
		decodeVoteRequest,
		encodeResponse,
	)

	http.Handle("/api/vote", votingHandler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func (voteService) CanVote(name string, age int) (string, error) {
	if name == "" || age <= 0 {
		return "", errors.New("Invalid properties")
	} else if age >= 18 {
		return "Hello " + name + ", you can vote!", nil
	}
	return "Hello " + name + ", you can't vote!", nil
}

// Endpoints are a primary abstraction in go-kit. An endpoint represents a single RPC (method in our service interface)
func makeVotingEndpoint(svc VoteService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(voteRequest)
		resp, err := svc.CanVote(req.Name, req.Age)
		if err != nil {
			return voteResponse{resp, err.Error()}, nil
		}
		return voteResponse{resp, ""}, nil
	}
}

// Decode JSON
func decodeVoteRequest(_ context.Context, httpRequest *http.Request) (interface{}, error) {
	var request voteRequest
	if err := json.NewDecoder(httpRequest.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

// Encode JSON
func encodeResponse(_ context.Context, rw http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(rw).Encode(response)
}
