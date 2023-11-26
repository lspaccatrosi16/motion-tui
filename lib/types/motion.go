package types

import (
	"bytes"
	"fmt"
	"strings"
	"text/tabwriter"
)

type TaskPriority int

const (
	ASAP TaskPriority = iota
	HIGH
	MEDIUM
	LOW
)

func (t TaskPriority) String() string {
	switch t {
	case ASAP:
		return "ASAP"
	case HIGH:
		return "HIGH"
	case MEDIUM:
		return "MEDIUM"
	case LOW:
		return "LOW"
	default:
		return "INVALID PRIORITY"
	}
}

type Task struct {
	Start       DT
	End         DT
	Due         DT
	Name        string
	Description string
	Priority    TaskPriority
	Id          string
}

func (t *Task) String() string {
	buf := bytes.NewBuffer(nil)

	tr := tabwriter.NewWriter(buf, 10, 0, 1, ' ', 0)

	fmt.Fprintln(buf, strings.Repeat("=", 50))
	fmt.Fprintf(tr, "Name\t%s\n", t.Name)
	fmt.Fprintf(tr, "Description\t%s\n", t.Description)
	fmt.Fprintf(tr, "Due\t%s\n", t.Due.ShortString())
	fmt.Fprintf(tr, "Priority\t%s\n", t.Priority.String())
	fmt.Fprintf(tr, "Start\t%s\n", t.Start.LongString())
	fmt.Fprintf(tr, "End\t%s\n", t.End.LongString())

	tr.Flush()

	fmt.Fprintln(buf, strings.Repeat("=", 50))

	return buf.String()
}
