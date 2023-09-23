package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type Test struct {
	Name string `json:"name_hihi"`
	Age int `json:",omitempty"` // 空值不顯示
	Email string `json:"-"` // 不顯示
	Arr map[string]string
}

func main() {
	
	test := Test{
		Name: "Mary",
		Age: 0,
		Arr: map[string]string{
			"1": "a",
			"2": "b",
		},
	}
	data, err := json.Marshal(test)
	fmt.Println(data, string(data), err)

	buf :=new(bytes.Buffer)
	json.Indent(buf, data, "", "\t")
	fmt.Println(buf.String())

	// 反向
	var test2 Test
	err = json.Unmarshal(data, &test2) // 將json string轉成struct
	fmt.Println("Unmarshal",test2, err) 
	
	// 檢測
	fmt.Println("Valid",json.Valid(data))

}
