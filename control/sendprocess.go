package control

import (
	"encoding/json"
	"log"
	"net/http"
	"processlist/database"
)

//SendProcess through connection
func SendProcess(w http.ResponseWriter, r *http.Request) {

	processes, err := database.SearchAllProcess()

	if err != nil {
		w.Write([]byte("[ERROR] It's not possible to pick processes"))
		log.Println("[ERROR] It's not possible to pick processes")
	}

	bpro, err := json.Marshal(processes)

	w.Write(bpro)

}
