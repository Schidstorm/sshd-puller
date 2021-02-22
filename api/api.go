package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Api struct {
	Endpoint string
}

func (receiver *Api) GetKeys(serverKey string) ([]string, error)  {
	res, err := http.Get(fmt.Sprintf("%s/serverKeys?key=%s", receiver.Endpoint, serverKey))
	if err != nil {
		return nil, err
	}

	buffer, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, err
	}

	var keys []string
	err = json.Unmarshal(buffer, &keys)
	if err != nil {
		return nil, err
	}

	return keys, nil
}