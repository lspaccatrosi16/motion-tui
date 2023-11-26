package main

import (
	"strconv"

	"github.com/lspaccatrosi16/go-cli-tools/input"
	"github.com/lspaccatrosi16/motion-tui/lib/data"
	"github.com/lspaccatrosi16/motion-tui/lib/display"
)

func main() {
	dayStr := input.GetInput("Days to get schedule for")

	days, err := strconv.Atoi(dayStr)
	if err != nil {
		panic(err)
	}

	err = data.GetList(days)

	if err != nil {
		panic(err)
	}

	err = display.DisplayData()
	if err != nil {
		panic(err)
	}

	// for _, task := range appData.Tasks {
	// 	fmt.Println(task.String())
	// }

}
