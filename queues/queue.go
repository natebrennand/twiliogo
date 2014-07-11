package queues

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/natebrennand/twiliogo/common"
)

const (
	queuesURL = "/2010-04-01/Accounts/%s/Queues/%s.json" // takes an AccountSid & QueueSid
)

// Account wraps the common Account struct to embed the AccountSid & Token.
type Account struct {
	common.Account
}

// Queue represents a queue json object returned by the Twilio api.
type Queue struct {
	Sid             string          `json:"sid"`
	FriendlyName    string          `json:"friendly_name"`
	CurrentSize     int             `json:"current_size"`
	AverageWaitTime int             `json:"average_wait_time"`
	MaxSize         int             `json:"max_size"`
	DateCreated     common.JSONTime `json:"date_created"`
	DateUpdated     common.JSONTime `json:"date_updated"`
	URI             string          `json:"uri"`
}

func validateQueueSid(sid string) bool {
	match, _ := regexp.MatchString(`^(QU)[0-9a-z]{32}$`, string(sid))
	return match
}

// Get a queue with the given SID.
func (act Account) Get(qsid string) (q Queue, err error) {
	if !validateQueueSid(qsid) {
		return q, errors.New("Invalid queue sid")
	}

	err = act.getQueue(fmt.Sprintf(queuesURL, act.AccountSid, qsid), &q)
	return
}

func (act Account) getQueue(url string, q *Queue) error {
	return common.SendGetRequest(url, act, &q)
}
