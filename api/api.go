package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Api struct {
	Endpoint string
}

type GetKeysResponse struct {
	Success bool     `json:"success"`
	Error   string   `json:"error"`
	Data    []string `json:"data"`
}

func (receiver *Api) GetKeys(ctx context.Context, serverKey string) ([]string, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s/serverKeys?key=%s", receiver.Endpoint, serverKey), nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	buffer, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &GetKeysResponse{}
	err = json.Unmarshal(buffer, response)
	if err != nil {
		return nil, err
	}
	if !response.Success {
		return nil, errors.New(response.Error)
	}

	return response.Data, nil
}
