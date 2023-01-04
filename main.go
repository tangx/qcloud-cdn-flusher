package main

import "github.com/tangx/qcloud-cdn-flusher/cmd"

func main() {
	err := cmd.Execute()
	if err != nil {
		panic(err)
	}
}
