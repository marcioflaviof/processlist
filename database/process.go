package database

import (
	"context"
	"log"
	"processlist/configs"
	"processlist/model"

	"go.mongodb.org/mongo-driver/bson"
)

func InsertProcess(p model.Process) (err error) {

	c := Db.Collection(configs.PROCESSES_COLLECTION)

	process, err := c.InsertOne(context.TODO(), p)

	if err != nil {
		log.Printf("[ERROR] Error in insertexam: %v, %v", err, p)
	}

	log.Printf("[INFO] exam inserted: %v", process)

	return

}

func SearchAllProcess() (p []model.Process, err error) {

	c := Db.Collection(configs.PROCESSES_COLLECTION)

	cur, err := c.Find(context.TODO(), bson.D{{}})

	if err != nil {
		log.Printf("[ERROR] Fail to find all results: %v %v", err, cur)
		return
	}

	if err = cur.All(context.TODO(), &p); err != nil {
		log.Printf("[ERROR] Fail to get all results %v", err, p)
	}

	log.Println("[INFO] All results searched")

	return

}
