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

var (
	validateCallSid = regexp.MustCompile(`^CA[0-9a-z]{32}$`).MatchString
)

var member = struct {
	Post, Front, Get, List string
}{
	Post:  "/2010-04-01/Accounts/%s/Queues/%s/Members/%s.json",    // takes an AccountSid & QueueSid & CallSid
	Front: "/2010-04-01/Accounts/%s/Queues/%s/Members/Front.json", // Takes an AccountSid & QueueSid
	Get:   "/2010-04-01/Accounts/%s/Queues/%s/Members/%s.json",    // takes an AccountSid & QueueSid & CallSid
	List:  "/2010-04-01/Accounts/%s/Queues/%s/Members.json",       // takes an AccountSid & QueueSid
}

// Member represents a queue member
// http://www.twilio.com/docs/api/rest/member#instance
type Member struct {
	CallSID      string          `json:"call_sid"`
	DateEnqueued common.JSONTime `json:"date_enqueued"`
	WaitTime     int             `json:"wait_time"`
	Position     int             `json:"position"`
}

// Front returns the front member of a queue
// http://www.twilio.com/docs/api/rest/member#instance-get
func (act Account) Front(queueSid string) (Member, error) {
	var m Member
	if !validateQueueSid(queueSid) {
		return m, errors.New("Invalid queue sid")
	}
	err := common.SendGetRequest(fmt.Sprintf(member.Front, act.AccountSid, queueSid), act, &m)
	return m, err
}

// GetMember returns a member with a given callsid from a queue
// http://www.twilio.com/docs/api/rest/member#instance-get
func (act Account) GetMember(queueSid string, callSid string) (Member, error) {
	var m Member
	if !validateQueueSid(queueSid) {
		return m, errors.New("Invalid queue sid")
	}
	if !validateCallSid(callSid) {
		return m, errors.New("Invalid call sid")
	}
	err := common.SendGetRequest(fmt.Sprintf(member.Get, act.AccountSid, queueSid, callSid), act, &m)
	return m, err
}

// MemberList represents a list of members in a queue
// http://www.twilio.com/docs/api/rest/member#list
type MemberList struct {
	common.ListResponseCore
	QueueMembers *[]Member `json:"queue_members"`
	act          *Account
}

// ListMembers returns a list of members in a queue
// http://www.twilio.com/docs/api/rest/member#list-get
func (act Account) ListMembers() (MemberList, error) {
	var memberList MemberList
	err := common.SendGetRequest(fmt.Sprintf(member.List, act.AccountSid), act, &memberList)
	memberList.act = &act
	return memberList, err
}

// Next sets the MemberList to the next page of the list resource, returns an error in the
// case that there are no more pages left.
func (ml *MemberList) Next() error {
	if ml.Page == ml.NumPages-1 {
		return errors.New("No more new pages")
	}
	return common.SendGetRequest(ml.NextPageURI, *ml.act, ml)
}

// Action is an action to be performed on a queue member
type Action struct {
	URL    string
	Method string
}

// GetReader implements the common.twilioPost interface
func (a Action) GetReader() io.Reader {
	vals := url.Values{}
	if a.URL != "" {
		vals.Set("URL", a.URL)
	}
	if a.Method != "" {
		vals.Set("Method", a.Method)
	}
	return strings.NewReader(vals.Encode())
}

// Validate implements the common.twilioPost interface
func (a Action) Validate() error {
	return nil
}

// DequeueFront dequeues a member from the front of the queue and perform an action on it
func (act Account) DequeueFront(a Action, queueSid string) (Member, error) {
	var m Member
	if !validateQueueSid(queueSid) {
		return m, errors.New("Invalid queue sid")
	}
	err := common.SendPostRequest(fmt.Sprintf(member.Front, act.AccountSid, queueSid), a, act, &m)
	return m, err
}

// DequeueCall dequeues a member with a given callsid from the queue and perform an action on it
func (act Account) DequeueCall(a Action, queueSid string, callSid string) (Member, error) {
	var m Member
	if !validateQueueSid(queueSid) {
		return m, errors.New("Invalid queue sid")
	}
	if !validateCallSid(callSid) {
		return m, errors.New("Invalid call sid")
	}
	err := common.SendPostRequest(fmt.Sprintf(member.Get, act.AccountSid, queueSid, callSid), a, act, &m)
	return m, err
}
