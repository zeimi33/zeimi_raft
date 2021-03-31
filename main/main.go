package main

import (
	"fmt"
	"github.com/zeimi33/zeimi_raft"
)

func main() {
	a := zeimi_raft.NewRaftLogger()
	a.Infof("fuck")
	fmt.Println(a)
}
