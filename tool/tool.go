package tool

import (
	"encoding/json"
	"log"
)

func Dump(data interface{}) {
	bites, _ := json.MarshalIndent(data, "", "\t")
	log.Printf("%s", string(bites))
}
