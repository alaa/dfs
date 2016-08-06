package main

import (
	"fmt"

	"github.com/hashicorp/memberlist"
)

func main() {
	list, err := memberlist.Create(memberlist.DefaultLocalConfig())
	if err != nil {
		panic("Failed to create memberlist: " + err.Error())
	}

	_, err = list.Join([]string{"127.0.0.1"})
	if err != nil {
		panic("Failed to join cluster: " + err.Error())
	}

	for _, member := range list.Members() {
		fmt.Printf("Member: %s %s\n", member.Name, member.Addr)
	}
}
