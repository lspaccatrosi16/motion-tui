package display

import (
	"bytes"
	"fmt"
	"io"
	"sort"
	"strings"
	"text/tabwriter"
	"time"

	panels "github.com/lspaccatrosi16/go-data-panels"
	"github.com/lspaccatrosi16/motion-tui/lib/types"
)

type taskTree struct {
	Tree *panels.DataTree
	Task *types.Task
}

type dayMap map[int]*[]taskTree

type dayTasks struct {
	Date  types.DT
	Tasks []taskTree
}

type dayList []dayTasks

func (d dayList) Len() int {
	return len(d)
}

func (d dayList) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

func (d dayList) Less(i, j int) bool {
	return d[i].Date.UnixSecs < d[j].Date.UnixSecs
}

func DisplayData() error {
	appData, err := types.GetAppData()
	if err != nil {
		return err
	}

	items := []*panels.MenuItem{
		{Name: "Tasks", Details: "View Tasks", Shortcut: 't'},
		{Name: "Days", Details: "View Day Summaries", Shortcut: 'd'},
	}

	baseTrees := makeIndividualTaskTrees(appData)

	tasksTree := makeTasksTree(baseTrees)
	daysTree := makeDaysTree(baseTrees)

	gui := panels.MakeGui(panels.GuiData{
		MenuItems: items,
		DataViews: []*panels.DataTree{tasksTree, daysTree},
	})

	gui.Run()

	return nil
}

func makeIndividualTaskTrees(appData *types.AppData) []taskTree {
	arr := []taskTree{}

	for i := 0; i < len(appData.Tasks); i++ {
		task := appData.Tasks[i]
		tree := panels.NewDataTree(task.Name)

		buf := bytes.NewBuffer(nil)
		tr := tabwriter.NewWriter(buf, 10, 0, 0, ' ', 0)

		addKv := addKvToWriter(tr)

		addKv("Description", task.Description)
		addKv("Due", task.Due.ShortString())
		addKv("Priority", task.Priority.String())
		addKv("Start", task.Start.LongString())
		addKv("End", task.End.LongString())

		tr.Flush()
		bLines := strings.Split(buf.String(), "\n")

		for _, l := range bLines {
			if l == "" {
				continue
			}
			tree.AddChild(l)
		}

		arr = append(arr, taskTree{
			Tree: tree,
			Task: task,
		})
	}

	return arr
}

func makeTasksTree(baseTrees []taskTree) *panels.DataTree {
	tree := panels.NewDataTree("Tasks")
	for _, bt := range baseTrees {
		tree.InheritTree(bt.Tree)
	}
	return tree
}

func makeDaysTree(baseTrees []taskTree) *panels.DataTree {
	tree := panels.NewDataTree("Days")

	daySortedMap := dayMap{}

	for _, bt := range baseTrees {
		key := time.Unix(int64(bt.Task.Start.UnixSecs), 0).YearDay()
		arrAddr := daySortedMap[key]
		arr := []taskTree{}
		if arrAddr != nil {
			arr = *arrAddr
		}

		arr = append(arr, bt)
		daySortedMap[key] = &arr
	}

	list := dayList{}

	for _, v := range daySortedMap {
		day := dayTasks{
			Date:  (*v)[0].Task.Start,
			Tasks: *v,
		}

		list = append(list, day)
	}

	sort.Sort(list)

	for _, day := range list {
		dayTree := tree.AddChild(day.Date.ShortString())
		for _, task := range day.Tasks {
			dayTree.InheritTree(task.Tree)
		}
	}

	return tree
}

func addKvToWriter(w io.Writer) func(key string, val string) {
	return func(key string, val string) {
		fmt.Fprintf(w, "%s\t : %s\n", key, val)
	}
}
