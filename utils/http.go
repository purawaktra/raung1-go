package utils

import (
	"context"
	"github.com/google/uuid"
	"github.com/purawaktra/raung1-go/functions"
	"io"
	"net/http"
	"time"
)

func CreateHTPPRequest(url string, method string, username string, password string, body io.Reader) (*http.Response, error, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	request, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		Error(err, "CreateHTPPRequest", "")
		return nil, err, "RC"
	}

	// create header and set basic auth
	requestId := uuid.New().String()
	request.Header.Set("host", AppCode)
	request.Header.Set("request-id", requestId)
	request.SetBasicAuth(username, password)

	// perform the HTTP request and check err
	client := http.DefaultClient
	response, err := client.Do(request)
	if err != nil {
		Error(err, "CreateHTPPRequest", "")
		return nil, err, "RR"
	}

	// check if basic auth ok
	if response.StatusCode == http.StatusUnauthorized {
		Error(err, "CreateHTPPRequest", response)
		return nil, err, "AU"
	}

	// check if response is bad request
	if response.StatusCode == http.StatusBadRequest {
		Error(err, "CreateHTPPRequest", response)
		return nil, err, "FR"
	}

	// check request id, response id on response
	if response.Header.Get("request-id") != requestId {
		Error(err, "CreateHTPPRequest", response)
		return nil, err, "FC"
	}
	if !functions.IsValidUUID(response.Header.Get("response-id")) {
		Error(err, "CreateHTPPRequest", response)
		return nil, err, "FC"
	}

	// create return
	return response, nil, "00"
}
