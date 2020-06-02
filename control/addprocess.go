package control

import (
	"io/ioutil"
	"log"
	"net/http"
)

//AddProcess adiciona os processos recebidos pelo client
func AddProcess(w http.ResponseWriter, r *http.Request) {

	body := r.Body
	bytes, err := ioutil.ReadAll(body)

	if err != nil {
		log.Println("[ERROR] Can't read response body", err)
	}

	log.Println(string(bytes))

}
