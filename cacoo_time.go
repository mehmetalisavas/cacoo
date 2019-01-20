package cacoo

import (
	"strings"
	"time"
)

// dateFormat represents the RFC2822 format for custom unmarshal functions.
const dateFormat = "Mon, 02 Jan 2006 15:04:05 -0700"

// CacooTime is used to parse default cacoo time (RFC2822). Unmarshal or Decode
// functions of json don't work because default time format is RFC3339 in Go.
type CacooTime struct {
	time.Time
}

// UnmarshalJSON implements the json Marshaller interface to parse the incoming
// struct
func (ct *CacooTime) UnmarshalJSON(input []byte) error {
	strInput := strings.Trim(string(input), `"`)

	cTime, err := time.Parse(dateFormat, strInput)
	if err != nil {
		return err
	}

	ct.Time = cTime
	return nil
}

// String implements  the string version of the timestamp
func (ct CacooTime) String() string {
	return ct.Time.String()
}
