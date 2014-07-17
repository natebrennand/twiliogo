package recording

import (
	"encoding/json"
	"testing"
)

func TestRecordingBuild(t *testing.T) {
	var r Recording
	err := json.Unmarshal([]byte(testResponseFixtureString), &r)
	if err != nil {
		t.Errorf("Building Recording from json string failed with error => %s", err.Error())
	}
	if r.Sid != testResponseFixture.Sid {
		t.Errorf(
			"Building Transcription from json string failed to properly set Sid of recording, expected %s, found %s",
			testResponseFixture.Sid,
			r.Sid,
		)
	}
}

func TestRecordingListBuild(t *testing.T) {
	var r RecordingList
	err := json.Unmarshal([]byte(testListFixtureString), &r)
	if err != nil {
		t.Errorf("Building Transcription from json string failed with error => %s", err.Error())
	}

	if len(*r.Recordings) != len(*testListFixture.Recordings) {
		t.Errorf(
			"Building TranscriptionList from json string failed to properly allocate the list of media, expected: %d, found %d",
			len(*testListFixture.Recordings),
			len(*r.Recordings),
		)
	}
}
