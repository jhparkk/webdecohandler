package main

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexPage(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	rsps, err := http.Get(ts.URL)

	assert.NoError(err)
	assert.Equal(http.StatusOK, rsps.StatusCode)
	data, _ := ioutil.ReadAll(rsps.Body)

	assert.Equal("hello world", string(data))
}

func TestDecorator(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	buf := &bytes.Buffer{}
	log.SetOutput(buf)

	rsps, err := http.Get(ts.URL)

	assert.NoError(err)
	assert.Equal(http.StatusOK, rsps.StatusCode)

	r := bufio.NewReader(buf)
	line, _, err := r.ReadLine()
	assert.NoError(err)
	assert.Contains(string(line), "[LOGGER1] Start")
}

func TestDecoratorSTDOUT(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	rsps, err := http.Get(ts.URL)

	assert.NoError(err)
	assert.Equal(http.StatusOK, rsps.StatusCode)
	//log.Print()
}
