package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func ParseResponse[T any](res *http.Response) (*T, error) {
	raw, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode >= 400 {
		var out ErrorResponse
		if err := json.Unmarshal(raw, &out); err != nil {
			return nil, fmt.Errorf("failed to unmarshal error response: %w", err)
		}

		return nil, fmt.Errorf("request failed: %d - %s", res.StatusCode, string(raw))
	}

	var out T
	if err := json.Unmarshal(raw, &out); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &out, nil
}
