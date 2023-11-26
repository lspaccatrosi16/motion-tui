package data

type reqTask struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Due         string `json:"dueDate"`
	Start       string `json:"scheduledStart"`
	End         string `json:"scheduledEnd"`
	Priority    string `json:"priority"`
}

type reqTaskList struct {
	Tasks          []reqTask `json:"tasks"`
	MetaNextCursor string    `json:"meta.nextCursor"`
	MetaPageSize   int       `json:"meta.pageSize"`
}
