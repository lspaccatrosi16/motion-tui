package types

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
