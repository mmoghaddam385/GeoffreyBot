package types

// This struct encapsulates the information in a group me message post request
type GroupMeMessagePost struct {
	Id string
	GroupId string
	Sender string
	SenderType string
	MessageText string
}

type ConsoleCommand interface {
	// Name returns the name of this command
	Name() string

	// Usage returns a short help string to be displayed when the user asks
	// for help
	Usage() string

	// Execute should run the command and return a result code.
	// 0 for success, anything else for error
	Execute(args []string) int
}