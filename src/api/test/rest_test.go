package test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"bytes"
	"io/ioutil"
)

func TestWhenPassACorrectJsonGetOkResponseCode(t *testing.T) {
	asserter := assert.New(t)

	var jsonStr = []byte(`{"Name":"pelota"}`)
	req, err := http.NewRequest("POST", ServerPath+ "/products", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Error(err.Error())
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Error(err.Error())
	}

	asserter.Equal(http.StatusCreated, resp.StatusCode)
	buff ,_:= ioutil.ReadAll(resp.Body)
	asserter.Equal(string(buff), `{"Id":1,"Name":"pelota"}`)
	
}
