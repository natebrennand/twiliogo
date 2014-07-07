package voice

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/natebrennand/twiliogo/common"
	"strconv"
	"time"
)

type VoiceResponseJson struct {
	common.ResponseCore
}