package main

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
)

func TestValidCEP(t *testing.T) {
	client := http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:8080/getweather/22070011", nil)
	assert.Nil(t, err)
	res, err := client.Do(req)
	assert.Nil(t, err)
	assert.Equal(t, res.StatusCode, http.StatusOK)
	body, err := io.ReadAll(res.Body)
	assert.Nil(t, err)
	var actualResponse ResponseHTTP
	err = json.Unmarshal(body, &actualResponse)
	assert.Nil(t, err)
	assert.Equal(t, "Rio De Janeiro", actualResponse.City)
}

func TestInvalidCEP(t *testing.T) {
	client := http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:8080/getweather/1234", nil)
	assert.Nil(t, err)
	res, err := client.Do(req)
	assert.Nil(t, err)
	assert.Equal(t, res.StatusCode, http.StatusUnprocessableEntity)
	body, err := io.ReadAll(res.Body)
	assert.Nil(t, err)
	assert.Equal(t, "invalid zipcode", string(body))
}

func TestNotFoundCEP(t *testing.T) {
	client := http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:8080/getweather/12345678", nil)
	assert.Nil(t, err)
	res, err := client.Do(req)
	assert.Nil(t, err)
	assert.Equal(t, res.StatusCode, http.StatusNotFound)
	body, err := io.ReadAll(res.Body)
	assert.Nil(t, err)
	assert.Equal(t, "can not find zipcode", string(body))
}
