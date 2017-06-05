package main

import (
	"fmt"

	"github.com/YingBurden/study-golang/pkg/json/cfg2json"
)

func main() {
	file := "/home/user/Desktop/slb/test/topology.cfg"
	gjson := cfg2json.NewJson(file)
	json1 := gjson.GetJson()
	fmt.Println(json1)
}
