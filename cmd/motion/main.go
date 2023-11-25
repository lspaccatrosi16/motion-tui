package main

import (
	"github.com/lspaccatrosi16/go-cli-tools/command"
)

func main() {
	manager := command.NewManager(command.ManagerConfig{Searchable: true})
	manager.Tui()
}
