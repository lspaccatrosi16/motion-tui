package types

import "time"

type DT struct {
	unixSecs int
	hasError bool
}

func (d *DT) ToISO() string {
	return time.Unix(int64(d.unixSecs), 0).Format(time.RFC3339)
}

func (d *DT) FromISO(str string) *DT {
	t, err := time.Parse(time.RFC3339, str)
	if err != nil {
		d.hasError = true
	} else {
		d.unixSecs = int(t.Unix())
	}
	return d
}

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
	Name        string
	Description string
	Priority    TaskPriority
}
