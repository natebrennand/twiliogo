package conference

import (
	"encoding/json"
	"github.com/natebrennand/twiliogo/common"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestConferenceBuild(t *testing.T) {
	var c Conference
	err := json.Unmarshal([]byte(testFixtureString), &c)
	if err != nil {
		t.Errorf("Building Recording from json string failed with error => %s", err.Error())
	}
	if c.Sid != testFixture.Sid {
		t.Errorf(
			"Building Transcription from json string failed to properly set Sid of recording, expected %s, found %s",
			testFixture.Sid,
			c.Sid,
		)
	}
}

func TestConferenceListBuild(t *testing.T) {
	var cl List
	err := json.Unmarshal([]byte(testFixtureListString), &cl)
	if err != nil {
		t.Errorf("Building Transcription from json string failed with error => %s", err.Error())
	}

	if len(*cl.Conferences) != len(*testFixtureList.Conferences) {
		t.Errorf(
			"Building TranscriptionList from json string failed to properly allocate the list of media, expected: %d, found %d",
			len(*testFixtureList.Conferences),
			len(*cl.Conferences),
		)
	}
}

func TestParticipantBuild(t *testing.T) {
	var p Participant
	err := json.Unmarshal([]byte(testParticipantFixtureString), &p)
	if err != nil {
		t.Errorf("Building Recording from json string failed with error => %s", err.Error())
	}
	if p.Muted != testParticipantFixture.Muted {
		t.Errorf(
			"Building Transcription from json string failed to properly set Sid of recording, expected %t, found %t",
			testParticipantFixture.Muted,
			p.Muted,
		)
	}
}

func TestQueryString(t *testing.T) {
	f := ListFilter{
		Status:       "completed",
		FriendlyName: "board meeting",
		DateCreated:  &common.JSONTime{Time: time.Date(2010, time.August, 16, 3, 45, 01, 0, &time.Location{})},
		DateUpdated:  &common.JSONTime{Time: time.Date(2010, time.August, 16, 3, 45, 03, 0, &time.Location{})},
	}
	qs := f.getQueryString()

	assert.Contains(t, qs, "Status=completed")
	assert.Contains(t, qs, "FriendlyName=board+meeting")
}

func TestConferenceGetBadReq(t *testing.T) {
	var c Conference
	respC, err := testAccount.Get("sldfkj")
	assert.Error(t, err)
	assert.Equal(t, c, respC)
}

func TestConferenceList(t *testing.T) {
	// TODO: once we concise HTTP testing
}
