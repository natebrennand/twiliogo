package voice

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestListBuild(t *testing.T) {
	var c CallList
	resp := http.Response{Body: ioutil.NopCloser(strings.NewReader(testListFixtureString))}
	err := c.Build(&resp)
	if err != nil {
		t.Errorf("Building call list from json failed with error => %s", err.Error())
	}
	if len(*c.Calls) != len(*testListFixture.Calls) {
		t.Errorf(
			"Building call list from json string failed to properly allocate the list of calls, expected: %d, found %d",
			len(*testListFixture.Calls),
			len(*c.Calls),
		)
	}
}
