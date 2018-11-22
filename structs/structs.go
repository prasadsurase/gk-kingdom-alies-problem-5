package structs

// State holds the state's details
type State struct {
	Name             string
	Emblem           string
	Alies            []*State
	IncomingMessages []map[string]string
	OutgoingMessages []map[string]string
}

// Message holds the message's details
type Message struct {
	Sender   *State
	Receiver *State
	Message  string
}
