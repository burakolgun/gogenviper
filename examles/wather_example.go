package main

import (
	"fmt"
	"github.com/burakolgun/gogenviper"
	"github.com/burakolgun/gogenviper/resources"
	"time"
)

func main() {

	var m resources.ExampleCfgModel
	w, err := gogenviper.Init("./resources", "cfg", "json", &m)
	if err != nil {
		panic(err)
	}

	w.Watch()

	for true {
		time.Sleep(time.Second * 1)
		fmt.Println(m)
	}
}
