package voice

import (
	"encoding/json"
	"testing"
)

func TestListBuild(t *testing.T) {
	var cl CallList
	err := json.Unmarshal([]byte(testListFixtureString), &cl)
	if err != nil {
		t.Errorf("Building call list from json failed with error => %s", err.Error())
	}
	if len(*cl.Calls) != len(*testListFixture.Calls) {
		t.Errorf(
			"Building call list from json string failed to properly allocate the list of calls, expected: %d, found %d",
			len(*testListFixture.Calls),
			len(*cl.Calls),
		)
	}
}
