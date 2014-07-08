package common

import (
	"fmt"
)

type Error struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
	MoreInfo string `json:"more_info"`
	Status   string `json:"status"`
}

func (e Error) Error() string {
	return fmt.Sprintf("Twilio Error %d => %s, more info @ %s", e.Code, e.Message, e.MoreInfo)
}
