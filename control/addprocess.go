package control

import (
	"io/ioutil"
	"log"
	"net/http"
	"processlist/database"
	"processlist/model"
	"strings"

	"github.com/rs/xid"
)

//AddProcess adiciona os processos recebidos pelo client
func AddProcess(w http.ResponseWriter, r *http.Request) {

	var process model.Process

	body := r.Body
	bytes, err := ioutil.ReadAll(body)

	if err != nil {
		log.Println("[ERROR] Can't read response body", err)
	}

	strbytes := strings.Split(string(bytes), ",")

	for _, i := range strbytes {

		process.ID = xid.New().String()

		process.Process = i

		err := database.InsertProcess(process)

		if err != nil {
			log.Println("Database is offline, please try again later")
		}

	}

	w.Write([]byte("All processes have been added"))

}
