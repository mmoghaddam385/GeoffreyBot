package types

// This struct encapsulates the information in a group me message post request
type GroupMeMessagePost struct {
	Id string
	GroupId string
	Sender string
	SenderType string
	MessageText string
}