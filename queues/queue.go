package queues

import (
	"errors"
	"fmt"
	"github.com/natebrennand/twiliogo/common"
	"io"
	"net/url"
	"regexp"
	"strings"
)

var queue = struct {
	Post, Get, List string
}{
	Post: "/2010-04-01/Accounts/%s/Queues/%s.json", // takes an AccountSid & QueueSid
	Get:  "/2010-04-01/Accounts/%s/Queues/%s.json", // takes an AccountSid & QueueSid
	List: "/2010-04-01/Accounts/%s/Queues.json",    // takes an AccountSid
}

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

var (
	validateQueueSid = regexp.MustCompile(`^QU[0-9a-z]{32}$`).MatchString
)

// Get a queue with the given SID.
func (act Account) Get(queueSid string) (Queue, error) {
	var q Queue
	if !validateQueueSid(queueSid) {
		return q, errors.New("Invalid queue sid")
	}

	err := common.SendGetRequest(fmt.Sprintf(queue.Get, act.AccountSid, queueSid), act, &q)
	return q, err
}

// Update is used to modify a queue
type Update struct {
	MaxSize      string
	FriendlyName string
}

// GetReader implements the common.twilioPost interface
func (u Update) GetReader() io.Reader {
	vals := url.Values{}
	if u.MaxSize != "" {
		vals.Set("MaxSize", u.MaxSize)
	}
	if u.FriendlyName != "" {
		vals.Set("FriendlyName", u.FriendlyName)
	}
	return strings.NewReader(vals.Encode())
}

// Validate implements the common.twilioPost interface
func (u Update) Validate() error {
	return nil
}

// Post updates a queue
func (act Account) Post(u Update, queueSid string) (Queue, error) {
	var q Queue
	if !validateQueueSid(queueSid) {
		return q, errors.New("Invalid queue sid")
	}
	err := common.SendPostRequest(fmt.Sprintf(queue.Post, act.AccountSid, queueSid), u, act, &q)
	return q, err
}

// Delete a queue from your account
// http://www.twilio.com/docs/api/rest/queue#instance-delete
func (act Account) Delete(queueSid string) error {
	if !validateQueueSid(queueSid) {
		return errors.New("Invalid queue sid")
	}
	return common.SendDeleteRequest(fmt.Sprintf(queue.Get, act.AccountSid, queueSid), act)
}

// QueueList represents a list of queues
type QueueList struct {
	common.ListResponseCore
	Queues *[]Queue `json:"queues"`
	act    *Account
}

// List returns a list of Twilio queues
// http://www.twilio.com/docs/api/rest/queue#list-get
func (act Account) List() (QueueList, error) {
	var queueList QueueList
	err := common.SendGetRequest(fmt.Sprintf(queue.List, act.AccountSid), act, &queueList)
	queueList.act = &act
	return queueList, err
}

// Next sets the QueueList to the next page of the list resource, returns an error in the
// case that there are no more pages left.
func (ql *QueueList) Next() error {
	if ql.Page == ql.NumPages-1 {
		return errors.New("No more new pages")
	}
	return common.SendGetRequest(ql.NextPageURI, *ql.act, ql)
}

// CreateQueue creates a new queue on this account, optionally with a specified FriendlyName
// and MaxSize
// http://www.twilio.com/docs/api/rest/queue#list-post
func (act Account) CreateQueue(u Update) (Queue, error) {
	var q Queue
	err := common.SendPostRequest(fmt.Sprintf(queue.List, act.AccountSid), u, act, &q)
	return q, err
}
