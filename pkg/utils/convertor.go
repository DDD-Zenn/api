package utils

import (
	"encoding/json"

	"github.com/google/uuid"
)

func MarshalAndInsert(data any, box any) {
	marshaledData, _ := json.Marshal(data)
	json.Unmarshal(marshaledData, box)
}

func GenId() string {
	uuidWithHyphen, _ := uuid.NewRandom()
	return uuidWithHyphen.String()
}
