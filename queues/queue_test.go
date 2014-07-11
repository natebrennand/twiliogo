package queues

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	"testing"

	"github.com/natebrennand/twiliogo/common"
	"github.com/stretchr/testify/assert"
)

const (
	goodQsid      = "QUasdfghjklqwertyuiopzxcvbnm123456"
	badQsid       = "QUnope"
	mockQueueJson = `{"sid": "QUasdfghjklqwertyuiopzxcvbnm123456", ` +
		`"friendly_name": "persistent_queue1", ` +
		`"current_size": 1, ` +
		`"average_wait_time": 2, ` +
		`"max_size": 10, ` +
		`"date_created": "Mon, ` +
		`26 Mar 2012 22:00:14 +0000", ` +
		`"date_updated": "Mon, ` +
		`26 Mar 2012 22:00:14 +0000", ` +
		`"uri": "/2010-04-01/Accounts/AC5ef87.../Queues/QUasdfghjklqwertyuiopzxcvbnm123456.json" }`
)

func TestValidateQueueSid(t *testing.T) {
	good := validateQueueSid(goodQsid)
	assert.True(t, good)
	bad := validateQueueSid(badQsid)
	assert.False(t, bad)
}

func TestUnmarshalQueue(t *testing.T) {
	var q Queue
	err := json.Unmarshal([]byte(mockQueueJson), &q)
	assert.NoError(t, err)
	assert.Exactly(t, q.Sid, goodQsid)
	assert.Exactly(t, q.CurrentSize, 1)
	assert.Exactly(t, q.AverageWaitTime, 2)
	assert.Exactly(t, q.MaxSize, 10)
	assert.Exactly(t, q.URI, "/2010-04-01/Accounts/AC5ef87.../Queues/QUasdfghjklqwertyuiopzxcvbnm123456.json")
}

func TestGetQueue(t *testing.T) {
	common.BaseURL = ""
	acct := Account{}
	acct.Client = http.Client{}
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, mockQueueJson)
	}))
	defer serv.Close()
	q := new(Queue)
	err := acct.getQueue(serv.URL, q)
	assert.NoError(t, err)
	assert.Exactly(t, q.Sid, goodQsid)
	assert.Exactly(t, q.CurrentSize, 1)
	assert.Exactly(t, q.AverageWaitTime, 2)
	assert.Exactly(t, q.MaxSize, 10)
	assert.Exactly(t, q.URI, "/2010-04-01/Accounts/AC5ef87.../Queues/QUasdfghjklqwertyuiopzxcvbnm123456.json")
}
