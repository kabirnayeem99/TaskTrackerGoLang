package task

type Command string

const (
	Add            Command = "add"
	Update         Command = "update"
	List           Command = "list"
	MarkDone       Command = "mark-done"
	MarkInProgress Command = "mark-in-progress"
	Delete         Command = "delete"
)
