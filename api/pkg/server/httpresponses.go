package server

import "encoding/json"

type errorResponse struct {
	Error string `json:"error"`
}

type successResponse struct {
	Success bool `json:"success"`
}

func marshalResponse(v any) string {
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return string(b)
}

var (
	respUnauthorized      = marshalResponse(errorResponse{Error: "unauthorized"})
	respNotFound          = marshalResponse(errorResponse{Error: "not found"})
	respRateLimitExceeded = marshalResponse(errorResponse{Error: "rate limit exceeded"})
	respInternalServer    = marshalResponse(errorResponse{Error: "internal server error"})
	respSuccess           = marshalResponse(successResponse{Success: true})
)
