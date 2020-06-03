package control

import (
	"io/ioutil"
	"log"
	"net/http"
	"processlist/model"
	"strings"
)

//AddProcess adiciona os processos recebidos pelo client
func AddProcess(w http.ResponseWriter, r *http.Request) {

	var processes model.Processes

	body := r.Body
	bytes, err := ioutil.ReadAll(body)

	if err != nil {
		log.Println("[ERROR] Can't read response body", err)
	}

	log.Println(string(bytes))

	strbytes := strings.Split(string(bytes), ",")

	for _, i := range strbytes {

		processes.AddProcess(i)

	}

}
