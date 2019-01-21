package cacoo

import (
	"context"
	"fmt"
	"testing"
)

func TestDiagramImage(t *testing.T) {
	// c := NewClient(TestToken)
	c := NewClient("asd")
	diagramID := "rkzgn5jGe2asBRCQ"
	_, err := c.Image.DownloadDiagram(context.Background(), diagramID)
	if err != nil {
		fmt.Println("err is:", err)
	}

}
