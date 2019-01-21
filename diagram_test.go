package cacoo

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"

	"github.com/bookerzzz/grok"
)

func TestListDiagrams(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/users/1MUJPfNEEeVUox15.json", func(w http.ResponseWriter, r *http.Request) {
		method(t, r, "GET")
		fmt.Fprint(w, `{
		    "result": [
		        {
		            "url": "https://cacoo.com/diagrams/00e77f4dc9973517",
		            "imageUrl": "https://cacoo.com/diagrams/00e77f4dc9973517.png",
		            "imageUrlForApi": "https://cacoo.com/api/v1/diagrams/00e77f4dc9973517.png",
		            "diagramId": "00e77f4dc9973517",
		            "title": "Wireframe",
		            "description": "Current project wireframe",
		            "security": "url",
		            "type": "normal",
		            "ownerName": "1MUJPfNEEeVUox15",
		            "ownerNickname": "Yoko",
		            "owner": {
		                "name": "1MUJPfNEEeVUox15",
		                "nickname": "Yoko",
		                "type": "cacoo",
		                "imageUrl": "https://cacoo.com/account/1MUJPfNEEeVUox15/image/32x32"
		            },
		            "editing": true,
		            "own": true,
		            "shared": false,
		            "folderId": 10001,
		            "folderName": "Design",
		            "sheetCount": 3,
		            "created": "Mon, 10 Aug 2009 17:00:00 +0900",
		            "updated": "Mon, 10 Aug 2009 17:00:00 +0900"
		        }
		    ],
		    "count": 1
		}`)
	})

	diagrams, _, err := client.Diagram.ListDiagrams(context.Background())
	if err != nil {
		t.Errorf("diagrams.listdiagrams returned error: %v", err)
	}
	dr := DiagramResult{
		URL:            String("https://cacoo.com/diagrams/00e77f4dc9973517"),
		ImageURL:       String("https://cacoo.com/diagrams/00e77f4dc9973517.png"),
		ImageURLForAPI: String("https://cacoo.com/api/v1/diagrams/00e77f4dc9973517.png"),
		DiagramID:      String("00e77f4dc9973517"),
		Title:          String("Wireframe"),
		Description:    String("Current project wireframe"),
		Security:       String("url"),
		Type:           String("normal"),
		OwnerName:      String("1MUJPfNEEeVUox15"),
		OwnerNickname:  String("Yoko"),
		Owner: &Account{
			Name:     String("1MUJPfNEEeVUox15"),
			Nickname: String("Yoko"),
			Type:     String("cacoo"),
			ImageURL: String("https://cacoo.com/account/1MUJPfNEEeVUox15/image/32x32"),
		},
		Editing:    Bool(true),
		Own:        Bool(true),
		Shared:     Bool(false),
		FolderID:   Int(10001),
		FolderName: String("Design"),
		SheetCount: Int(3),
		Created:    &CacooTime{time.Date(dateFormat)},
		Updated:    String("Mon, 10 Aug 2009 17:00:00 +0900"),
	}
	want := &Diagrams{
		Result: []DiagramResult{dr},
		Count:  Int(1),
	}

	if !reflect.DeepEqual(user, want) {
		t.Errorf("want: %v, but got: %v", want, user)
	}
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
