package main

import (
	"encoding/json"
	"fmt"
)

type Astruct struct {
	Name string `json:"name"`
}

func main() {
	ss := []string{"a", "b", "c", "d", "c", "b", "a"}
	j, _ := json.Marshal(ss)
	result := make([]string, 0)
	_ = json.Unmarshal([]byte(j), &result)
	fmt.Printf("result:%v", result)
	// amap := make(map[string]string)
	// // var astruct Astruct
	// // _ = json.Unmarshal([]byte(s), &amap)
	// amap["name"] = "name"
	// amap["type"] = "type"
	// fmt.Printf("unmarshal: %v\n", amap)
	// data, _ := json.Marshal(amap)
	// fmt.Printf("marshal: %s\n", string(data))
}
