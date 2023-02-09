package websockets

// NOTE: BroadcastError
type _broadcastMessageError struct {
	s string
}

func (e *_broadcastMessageError) Error() string {
	if e.s != "" {
		return "BroadcastError: " + e.s
	} else {
		return "BroadcastError occurred!!!"
	}
}

func BroadcastMessageError(message ...string) error {
	// Checks if there is some message
	var errMessage string
	if len(message) != 0 {
		errMessage = message[0]
	} else {
		errMessage = ""
	}

	// Returns pointer to Error struct
	newError := _broadcastMessageError{errMessage}
	return &newError
}
