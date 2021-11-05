package main

import (
	"encoding/json"
	"fmt"
)

type Astruct struct {
	Name string `json:"name"`
}

func main() {
	// s := "{\"name\":\"name\",\"type\":\"type\"}"
	amap := make(map[string]string)
	// var astruct Astruct
	// _ = json.Unmarshal([]byte(s), &amap)
	amap["name"] = "name"
	amap["type"] = "type"
	fmt.Printf("unmarshal: %v\n", amap)
	data, _ := json.Marshal(amap)
	fmt.Printf("marshal: %s\n", string(data))
}
