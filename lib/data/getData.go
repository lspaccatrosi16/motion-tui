package data

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/lspaccatrosi16/motion-tui/lib/request"
	"github.com/lspaccatrosi16/motion-tui/lib/types"
)

func GetList(days int) error {
	appData, err := types.GetAppData()
	if err != nil {
		return err
	}

	cursor := ""

	for {
		data, err := makeListRequest(cursor)
		if err != nil {
			return err
		}

		for _, rawTask := range data.Tasks {
			start := new(types.DT).FromISO(rawTask.Start)
			end := new(types.DT).FromISO(rawTask.End)
			due := new(types.DT).FromISO(rawTask.Due)

			var priority types.TaskPriority

			switch rawTask.Priority {
			case "ASAP":
				priority = types.ASAP
			case "HIGH":
				priority = types.HIGH
			case "MEDIUM":
				priority = types.MEDIUM
			case "LOW":
				priority = types.LOW
			}

			var project types.Project

			if rawTask.Project.Name != "" {
				project.Name = rawTask.Project.Name
				project.Description = rawTask.Project.Description
				project.Id = rawTask.Project.Id
			} else {
				project.Name = "No Project"
			}

			task := types.Task{
				Start:       *start,
				End:         *end,
				Due:         *due,
				Name:        rawTask.Name,
				Description: rawTask.Description,
				Priority:    priority,
				Id:          rawTask.Id,
				Project:     project,
			}

			if task.Start.InDays(days) && !start.HasError && !end.HasError {
				appData.Tasks = append(appData.Tasks, &task)
			}
		}

		if data.MetaNextCursor == "" {
			break
		}
		cursor = data.MetaNextCursor
	}

	sort.Sort(appData.Tasks)

	return nil
}

func makeListRequest(cursor string) (*reqTaskList, error) {
	url := "https://api.usemotion.com/v1/tasks"

	if cursor != "" {
		url += "?cursor=" + cursor
	}

	unparsed, err := request.MakeGetRequest(url)

	if err != nil {
		return nil, fmt.Errorf("request error: %s", err.Error())
	}

	var taskListRequest reqTaskList

	err = json.Unmarshal(unparsed, &taskListRequest)

	if err != nil {
		return nil, fmt.Errorf("unmarshal error: %s", err.Error())
	}

	return &taskListRequest, nil
}
