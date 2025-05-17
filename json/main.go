package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First string `json:"first"`
	Last  string `json:"last"`
	Age   int    `json:"age"`
}

func main() {
	p1 := `[{"First":"tharika","Age":4},{"First":"henry","Age":24},{"First":"James","Age":54}]`
	bs := []byte(p1)
	var people []Person
	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}

	for i, v := range people {
		fmt.Println(i, v)
	}
}

// func Marshall(v interface{})([]byte,error)
// func UnMarshall(data []byte,v interface{}) error
