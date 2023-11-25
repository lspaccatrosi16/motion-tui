package main

import (
	"fmt"

	"github.com/lspaccatrosi16/motion-tui/lib/types"
)

func main() {
	rt := new(types.RT)
	err := rt.Load()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Num Requests: %d\n", rt.NumRequests())
	fmt.Printf("Can request: %t\n", rt.CanRequest())
	if rt.CanRequest() {
		rt.LogRequest()
	}

	fmt.Printf("Num Requests: %d\n", rt.NumRequests())

	err = rt.Save()
	if err != nil {
		panic(err)
	}
	// manager := command.NewManager(command.ManagerConfig{Searchable: true})
	// manager.Tui()
}
