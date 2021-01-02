package app

import (
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"net/http"

	jsoniter "github.com/json-iterator/go"
)

//parse body of http Request
func parseBody(r *http.Request, req interface{}) bool {
	if r.Body == nil {
		return false
	}
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("ioutil.ReadAll", err)
		return false
	}
	err = jsoniter.Unmarshal(b, req)
	if err != nil {
		log.Printf("ERROR: %v | req body: %v \n", err, string(b))
		return false
	}
	return true
}

func writeJSONBody(response http.ResponseWriter, dto interface{}) (err error) {
	response.Header().Set("Content-Type", "application/json")

	body, err := jsoniter.Marshal(dto)
	if err != nil {
		return errors.New("error")
	}

	_, err = response.Write(body)
	if err != nil {
		return errors.New("error")
	}

	return nil
}