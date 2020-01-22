package Tool

import (
	"encoding/json"
	"log"
	"time"
)

func Dump(data interface{}) {
	bites, _ := json.MarshalIndent(data, "", "\t")
	log.Printf("%s", string(bites))
}

func Log(message string, data interface{}) {
	now := time.Now().Format("Y-m-d H:i:s")
	bites, _ := json.MarshalIndent(data, "", "\t")
	dataString := string(bites)
	log.Printf("[%s] %s : %s", now, message, dataString)
}
