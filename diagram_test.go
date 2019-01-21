package cacoo

import (
	"context"
	"fmt"
	"testing"

	"github.com/bookerzzz/grok"
)

func TestDiagrams(t *testing.T) {
	c := NewClient(TestToken)
	// c := NewClient("asd")
	diagrams, _, err := c.Diagram.ListDiagrams(context.Background())
	if err != nil {
		fmt.Println("err is:", err)
	}

	grok.V(diagrams)
	fmt.Println(diagrams)
	_ = diagrams
}

func TestDiagramWithName(t *testing.T) {
	c := NewClient(TestToken)
	// c := NewClient("asd")
	name := "rkzgn5jGe2asBRCQ"
	diagram, _, err := c.Diagram.GetDiagram(context.Background(), name)
	if err != nil {
		fmt.Println("err is:", err)
	}

	grok.V(diagram)
	fmt.Println(diagram)
	_ = diagram
}
func TestDiagramGet(t *testing.T) {
	c := NewClient(TestToken)
	// c := NewClient("asd")
	name := "rkzgn5jGe2asBRCQ"
	diagram, _, err := c.Diagram.GetDiagram(context.Background(), name)
	if err != nil {
		fmt.Println("err is:", err)
	}

	grok.V(diagram)
	fmt.Println(diagram)
	_ = diagram
}
func TestDiagramCreate(t *testing.T) {
	c := NewClient(TestToken)
	// c := NewClient("asd")
	opt := &DiagramOption{
		FolderID:    351198,
		Title:       "diagram-title test",
		Description: "diagram-description test2",
		Security:    "url",
	}

	diagram, _, err := c.Diagram.Create(context.Background(), opt)
	if err != nil {
		fmt.Println("err is:", err)
	}

	grok.V(diagram)
	fmt.Println(diagram)
	_ = diagram
}
func TestDiagramDelete(t *testing.T) {
	c := NewClient(TestToken)
	// c := NewClient("asd")

	diagramID := "voD5W1IRnsizEpcm"

	_, err := c.Diagram.Delete(context.Background(), diagramID)
	if err != nil {
		fmt.Println("err is:", err)
	}

}

func TestDiagramToken(t *testing.T) {
	c := NewClient(TestToken)
	// c := NewClient("asd")

	diagramID := "rkzgn5jGe2asBRCQ"

	token, _, err := c.Diagram.GetToken(context.Background(), diagramID)
	if err != nil {
		fmt.Println("err is:", err)
	}

	_ = token

}
