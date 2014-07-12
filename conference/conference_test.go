package conference

import (
	"encoding/json"
	"testing"
)

func TestConferenceListBuild(t *testing.T) {
	var cl ConferenceList
	err := json.Unmarshal([]byte(testFixtureListString), &cl)
	if err != nil {
		t.Error("Building Transcription from json string failed with error => %s", err.Error())
	}

	if len(*cl.Conferences) != len(*testFixtureList.Conferences) {
		t.Errorf(
			"Building TranscriptionList from json string failed to properly allocate the list of media, expected: %d, found %d",
			len(*testFixtureList.Conferences),
			len(*cl.Conferences),
		)
	}
}
