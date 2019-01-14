package main

import (
	"fmt"
	"github.com/tejasmanohar/timerange-go"
	"time"
)

var find  = `while`
var replase = `golang`

func timer() {
	start := time.Date(2019, 1, 14, 16, 24, 0, 0, time.UTC)
	end := time.Date(2020, 8, 28, 11, 0, 0, 0, time.UTC)
	iter := timerange.New(start, end, time.Second)
	for iter.Next() {
		t := iter.Current()
		fmt.Println(t)
	}
}

func main() {
	//var path = `C:/Users/Administrator/go/src/website/content/en/docs/tasks/administer-cluster/kubeadm/`
	//examples.GetFilelist(path)
	//cmd := exec.Command("D:/bin/go-gitlab-webhook.exe", "D:/bin/config.json")
	//if err := cmd.Run(); err != nil {
	//	fmt.Println("Error: ", err)
	//}

	//alarmClock()
	//cmd := exec.Command("D:/bin/go-gitlab-webhook.exe", "D:/bin/config.json")
	//if err := cmd.Run(); err != nil {
	//	fmt.Println("Error: ", err)
	//}
}