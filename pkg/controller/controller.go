package controller

import (
	"fmt"
	"triproute/pkg/repository"
)

// RequestInsertRoute user request for new route
type RequestInsertRoute struct {
	Start string `form:"start" json:"start"`
	End   string `form:"end" json:"end"`
	Cost  int64  `form:"cost" json:"cost"`
}

// ErrorResponse json response in case of error
type ErrorResponse struct {
	Msg string `json:"message,omitempty"`
}

// RequestBestRoute request parameters for BestRoute
type RequestBestRoute struct {
	Start, End string
}

// ResponseBestRoute struct to be formated to json response
type ResponseBestRoute struct {
	Cost  int64  `json:"cost,omitempty"`
	Route string `json:"route,omitempty"`
}

// GetBestRoute return the best rout for user
func GetBestRoute(get RequestBestRoute, r *repository.Repository) (ResponseBestRoute, error) {
	route, cost, err := r.FindRoute(get.Start, get.End)

	if err != nil {
		return ResponseBestRoute{}, err
	}

	response := FormatResponse(route, cost)

	return ResponseBestRoute{Route: response, Cost: cost}, nil
}

// InsertNewRoute handle new route to be saved
func InsertNewRoute(post RequestInsertRoute, r *repository.Repository) error {

	if post.Start == "" || post.End == "" || post.Cost == 0 {
		return fmt.Errorf("Invalid route formate: start=\"%s\", end=\"%s\", cost=\"%d\"", post.Start, post.End, post.Cost)
	}

	err := r.InsertRoute(post.Start, post.End, post.Cost)
	return err
}
