package conference

import (
	"github.com/natebrennand/twiliogo/common"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

var (
	testAccount  = Account{common.Account{}}
	updateFalse  = ParticipantUpdate{Muted: false}
	updateTrue   = ParticipantUpdate{Muted: true}
	validConfSid = "CF00000000000000000000000000000000"
	validCallSid = "CA00000000000000000000000000000000"
)

func TestParticipantUpdateQS(t *testing.T) {
	qs := updateFalse.getParticipantQueryString()
	assert.Equal(t, "?Muted=false", qs)

	qs = updateTrue.getParticipantQueryString()
	assert.Equal(t, "?Muted=true", qs)
}

func TestParticipantUpdateUselessValidate(t *testing.T) {
	assert.Nil(t, updateTrue.Validate())
	assert.Nil(t, updateFalse.Validate())
}

func TestParticipantUpdateReader(t *testing.T) {
	r := updateFalse.GetReader()
	b, err := ioutil.ReadAll(r)
	assert.Nil(t, err)
	assert.Equal(t, "Muted=false", string(b))
}

func TestGetParticipant(t *testing.T) {
	var p Participant
	rp, err := testAccount.GetParticipant("sldkfj", "slkdfj")
	assert.Error(t, err)
	assert.Equal(t, p, rp)

	rp, err = testAccount.GetParticipant(validConfSid, "slkdfj")
	assert.Error(t, err)
	assert.Equal(t, p, rp)

	// TODO: concisely mock out actual HTTP call
}

func TestKickParticipant(t *testing.T) {
	err := testAccount.Kick("sldkfj", "slkdfj")
	assert.Error(t, err)

	err = testAccount.Kick(validConfSid, "slkdfj")
	assert.Error(t, err)

	// TODO: concisely mock out actual HTTP call
}

func TestParticipantUpdate(t *testing.T) {
	var p Participant
	rp, err := testAccount.SetMute("sldkfjl", "sldkfj", updateTrue)
	assert.Error(t, err)
	assert.Equal(t, p, rp)

	rp, err = testAccount.SetMute(validConfSid, "slkdfj", updateTrue)
	assert.Error(t, err)
	assert.Equal(t, p, rp)

	// TODO: concisely mock out actual HTTP call
}

func TestParticipantList(t *testing.T) {
	_, err := testAccount.ListParticipants(ParticipantStatus{Muted: true}, "sldkfjl")
	assert.Error(t, err)

	// TODO: concisely mock out actual HTTP call
}

func TestParticipantListNext(t *testing.T) {
	var pl ParticipantList
	pl.NumPages = 1
	pl.Page = 0
	assert.Error(t, pl.Next())

	// TODO: concisely mock out actual HTTP call
}
