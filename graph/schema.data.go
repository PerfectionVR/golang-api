package graph

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/PerfectionVR/golang-api/graph/model"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

func getJson(url string, target interface{}) error {
	fmt.Println("Requesting", url)
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(target)
}

func getApplications() []*model.Application {
	data := make([]*model.Application, 0)
	getJson("https://api.utomik.com/v1/applications", &data)
	return data
}
