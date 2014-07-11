package transcription

import (
	"encoding/json"
	"testing"
)

func TestTranscriptionBuild(t *testing.T) {
	var trans Transcription
	err := json.Unmarshal([]byte(testResponseFixtureString), &trans)
	if err != nil {
		t.Errorf("Building Transcription from json string failed with error => %s", err.Error())
	}
	if trans.TranscriptionText != testResponseFixture.TranscriptionText {
		t.Errorf(
			"Building Transcription from json string failed to properly set TranscriptionText, expected %s, found %s",
			testResponseFixture.TranscriptionText,
			trans.TranscriptionText,
		)
	}
}

func TestTranscriptionListBuild(t *testing.T) {
	var trans TranscriptionList
	err := json.Unmarshal([]byte(testListFixtureString), &trans)
	if err != nil {
		t.Error("Building Transcription from json string failed with error => %s", err.Error())
	}

	if len(*trans.Transcriptions) != len(*testListFixture.Transcriptions) {
		t.Errorf(
			"Building TranscriptionList from json string failed to properly allocate the list of media, expected: %d, found %d",
			len(*testListFixture.Transcriptions),
			len(*trans.Transcriptions),
		)
	}
}
