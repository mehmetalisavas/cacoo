package cacoo

import (
	"context"
	"fmt"
	"testing"

	"github.com/bookerzzz/grok"
)

func TestLicense(t *testing.T) {
	c := NewClient(TestToken)
	// c := NewClient("asd")

	licence, _, err := c.License.GetLicense(context.Background())
	if err != nil {
		fmt.Println("err is:", err)
	}

	grok.V(licence)
}
