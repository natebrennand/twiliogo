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

func TestPostValidate(t *testing.T) {
	r, err := ioutil.ReadAll(testUpdateFixture.GetReader())
	assert.Nil(t, err)
	form := string(r)

	assert.Contains(t, form, "MaxSize=300")
	assert.Contains(t, form, "FriendlyName=newname")
}

func TestDequeueValidate(t *testing.T) {
	r, err := ioutil.ReadAll(testActionFixture.GetReader())
	assert.Nil(t, err)
	form := string(r)

	assert.Contains(t, form, "Url=http%3A%2F%2Fwww.example.com")
	assert.Contains(t, form, "Method=POST")
}

func TestGetQueue(t *testing.T) {
	act := Account{}
	_, err := act.Get("garbage")
	assert.Error(t, err)

	// TODO: test HTTP get
}

func TestGetFront(t *testing.T) {
	act := Account{}
	_, err := act.Front("garbage")
	assert.Error(t, err)
}

func TestGetMember(t *testing.T) {
	act := Account{}
	_, err := act.GetMember("garbage", "garbage")
	assert.Error(t, err)
	_, err2 := act.GetMember("garbage", "CA386025c9bf5d6052a1d1ea42b4d16663")
	assert.Error(t, err2)
	_, err3 := act.GetMember("QU386025c9bf5d6052a1d1ea42b4d16663", "garbage")
	assert.Error(t, err3)
}

func TestNextFail(t *testing.T) {
	var ql QueueList
	ql.Page = 0
	ql.NumPages = 1
	assert.Error(t, ql.Next())

	// TODO: test HTTP call
}
