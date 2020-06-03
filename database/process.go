package database

import (
	"context"
	"log"
	"processlist/configs"
	"processlist/model"
)

func InsertProcess(p model.Processes) (err error) {

	c := Db.Collection(configs.PROCESSES_COLLECTION)

	process, err := c.InsertOne(context.TODO(), p.Process)

	if err != nil {
		log.Printf("[ERROR] Error in insertexam: %v, %v", err, p)
	}

	log.Printf("[INFO] exam inserted: %v", process)

	return

}
