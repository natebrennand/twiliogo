package queues

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

var (
	validQueueSid = "QU5ef8732a3c49700934481addd5ce1659"
)

func TestValidateQueuePostSuccess(t *testing.T) {
	u := Update{MaxSize: 1000, FriendlyName: "friendly_queue"}
	if nil != u.Validate() {
		t.Error("Validation of valid queue update failed.")
	}

	u = Update{MaxSize: 1000}
	if nil != u.Validate() {
		t.Error("Validation of valid queue update failed.")
	}

	u = Update{FriendlyName: "friendly_queue"}
	if nil != u.Validate() {
		t.Error("Validation of valid queue update failed.")
	}

}

func TestValidateQueuePostFailure(t *testing.T) {
	u := Update{}
	if nil == u.Validate() {
		t.Error("Validation of queue update missing both fields failed.")
	}
}

func TestPostValidate(t *testing.T) {
	r, err := ioutil.ReadAll(testUpdateFixture.GetReader())
	assert.Nil(t, err)
	form := string(r)

	assert.Contains(t, form, "MaxSize=300")
	assert.Contains(t, form, "FriendlyName=newname")
}

func TestGetQueue(t *testing.T) {
	act := Account{}
	_, err := act.Get("garbage")
	assert.Error(t, err)

	// TODO: test HTTP get
}

func TestNextFail(t *testing.T) {
	var ql QueueList
	ql.Page = 0
	ql.NumPages = 1
	assert.Error(t, ql.Next())

	// TODO: test HTTP call
}
