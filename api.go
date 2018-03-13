package aliceapi

import (
	"io/ioutil"
	"encoding/json"
	"net/http"
	"fmt"
	"log"
)

func Handle(pattern string, handler func (Request) (Response)) {
	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		bytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}

		var req req
		err = json.Unmarshal(bytes, &req)
		if err != nil {
			log.Fatal(err)
		}

		response := res{
			Version: req.Version,
			Session: req.Session,
			Response: handler(req.Request),
		}

		bytes, err = json.Marshal(response)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Fprint(w, string(bytes))
	})
}