package module2

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

func TestGreet(t *testing.T) {
	customHeaders := map[string]string{"a": "1", "b": "2", "c": "3"}
	req, err := http.NewRequest(http.MethodGet, "http://127.0.0.1/greet", nil)
	assert.Nil(t, err)
	for k, v := range customHeaders {
		req.Header.Add(k, v)
	}
	resp, err := http.DefaultClient.Do(req)
	assert.Nil(t, err)
	for k, v := range customHeaders {
		assert.Equal(t, resp.Header.Get(k), v)
	}
	assert.Equal(t, resp.StatusCode, http.StatusOK)
	assert.Equal(t, resp.Header.Get(VERSION), os.Getenv(VERSION))
}

func TestHealthz(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "http://127.0.0.1/healthz", nil)
	assert.Nil(t, err)
	resp, err := http.DefaultClient.Do(req)
	assert.Nil(t, err)
	assert.Equal(t, resp.StatusCode, http.StatusOK)
}

func TestMain(m *testing.M) {
	go RunServer()
	m.Run()
}
