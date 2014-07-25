package conference

import (
	"github.com/natebrennand/twiliogo/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	testAccount = Account{common.Account{}}
	updateFalse = ParticipantUpdate{Muted: false}
	updateTrue  = ParticipantUpdate{Muted: true}
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

func TestGetParticipantBadReq(t *testing.T) {
	var p Participant
	rp, err := testAccount.GetParticipant("sldkfj", "slkdfj")
	assert.Error(t, err)
	assert.Equal(t, p, rp)

	rp, err = testAccount.GetParticipant("CF00000000000000000000000000000000", "slkdfj")
	assert.Error(t, err)
	assert.Equal(t, p, rp)

	// TODO: concisely mock out actual HTTP call
}

func TestParticipantUpdateBadReq(t *testing.T) {
	var p Participant
	rp, err := testAccount.SetMute("sldkfjl", "sldkfj", updateTrue)
	assert.Error(t, err)
	assert.Equal(t, p, rp)

	rp, err = testAccount.SetMute("CF00000000000000000000000000000000", "slkdfj", updateTrue)
	assert.Error(t, err)
	assert.Equal(t, p, rp)

	// TODO: concisely mock out actual HTTP call
}
