package cacoo

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestDiagramListDiagrams(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/diagrams.json", func(w http.ResponseWriter, r *http.Request) {
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

	want := testDiagrams(t)

	if !reflect.DeepEqual(diagrams, want) {
		t.Errorf("want: %v, but got: %v", want, diagrams)
	}
}

func TestDiagramGetDiagram(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/diagrams/00e77f4dc9973517.json", func(w http.ResponseWriter, r *http.Request) {
		method(t, r, "GET")
		fmt.Fprint(w, `{
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
		    "updated": "Mon, 10 Aug 2009 17:00:00 +0900",
		    "sheets": [
		        {
		            "url": "https://cacoo.com/diagrams/00e77f4dc9973517#aaaa",
		            "imageUrl": "https://cacoo.com/diagrams/00e77f4dc9973517-aaaa.png",
		            "imageUrlForApi": "https://cacoo.com/api/v1/diagrams/00e77f4dc9973517-aaaa.png",
		            "uid": "aaaa",
		            "name": "sheet1",
		            "width": 200,
		            "height": 100
		        }
		    ],
		    "comments": [
		        {
		            "user": {
		                "name": "1MUJPfNEEeVUox15",
		                "nickname": "Yoko",
		                "type": "cacoo",
		                "imageUrl": "https://cacoo.com/account/1MUJPfNEEeVUox15/image/32x32"
		            },
		            "content": "comment 1",
		            "created": "Mon, 10 Aug 2009 17:00:00 +0900",
		            "updated": "Mon, 10 Aug 2009 17:00:00 +0900"
		        }
		    ]
		}`)
	})

	diagram, _, err := client.Diagram.GetDiagram(context.Background(), "00e77f4dc9973517")
	if err != nil {
		t.Errorf("diagram.getdiagram returns error: %v", err)
	}

	want := testDiagramResultWithSheetAndComment(t)

	if !reflect.DeepEqual(diagram, want) {
		t.Errorf("want: %v, but got: %v", want, diagram)
	}
}

func TestDiagramCreate(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/diagrams/create.json", func(w http.ResponseWriter, r *http.Request) {
		method(t, r, "POST")
		fmt.Fprint(w, `{
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
		    "updated": "Mon, 10 Aug 2009 17:00:00 +0900",
		    "sheets": [
		        {
		            "url": "https://cacoo.com/diagrams/00e77f4dc9973517#aaaa",
		            "imageUrl": "https://cacoo.com/diagrams/00e77f4dc9973517-aaaa.png",
		            "imageUrlForApi": "https://cacoo.com/api/v1/diagrams/00e77f4dc9973517-aaaa.png",
		            "uid": "aaaa",
		            "name": "sheet1",
		            "width": 200,
		            "height": 100
		        }
		    ],
		    "comments": [
		        {
		            "user": {
		                "name": "1MUJPfNEEeVUox15",
		                "nickname": "Yoko",
		                "type": "cacoo",
		                "imageUrl": "https://cacoo.com/account/1MUJPfNEEeVUox15/image/32x32"
		            },
		            "content": "comment 1",
		            "created": "Mon, 10 Aug 2009 17:00:00 +0900",
		            "updated": "Mon, 10 Aug 2009 17:00:00 +0900"
		        }
		    ]
		}`)
	})

	generatedDiagram := testGenerateDiagramOption(t)
	diagram, _, err := client.Diagram.Create(context.Background(), generatedDiagram)
	if err != nil {
		t.Errorf("diagram.create returns err: %q", err)
	}

	want := testDiagramResultWithSheetAndComment(t)

	if !reflect.DeepEqual(diagram, want) {
		t.Errorf("want: %v, but got: %v", want, diagram)
	}
}

func TestDiagramGet(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/diagrams/create.json", func(w http.ResponseWriter, r *http.Request) {
		method(t, r, "GET")
		fmt.Fprint(w, `{
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
		    "updated": "Mon, 10 Aug 2009 17:00:00 +0900",
		    "sheets": [
		        {
		            "url": "https://cacoo.com/diagrams/00e77f4dc9973517#aaaa",
		            "imageUrl": "https://cacoo.com/diagrams/00e77f4dc9973517-aaaa.png",
		            "imageUrlForApi": "https://cacoo.com/api/v1/diagrams/00e77f4dc9973517-aaaa.png",
		            "uid": "aaaa",
		            "name": "sheet1",
		            "width": 200,
		            "height": 100
		        }
		    ],
		    "comments": [
		        {
		            "user": {
		                "name": "1MUJPfNEEeVUox15",
		                "nickname": "Yoko",
		                "type": "cacoo",
		                "imageUrl": "https://cacoo.com/account/1MUJPfNEEeVUox15/image/32x32"
		            },
		            "content": "comment 1",
		            "created": "Mon, 10 Aug 2009 17:00:00 +0900",
		            "updated": "Mon, 10 Aug 2009 17:00:00 +0900"
		        }
		    ]
		}`)
	})

	diagram, _, err := client.Diagram.Get(context.Background())
	if err != nil {
		t.Errorf("diagram.create returns err: %q", err)
	}

	want := testDiagramResultWithSheetAndComment(t)

	if !reflect.DeepEqual(diagram, want) {
		t.Errorf("want: %v, but got: %v", want, diagram)
	}
}

func TestDiagramDelete(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/diagrams/00e77f4dc9973517/delete.json", func(w http.ResponseWriter, r *http.Request) {
		method(t, r, "GET")
		fmt.Fprint(w, `{}`)
	})

	resp, err := client.Diagram.Delete(context.Background(), "00e77f4dc9973517")
	if err != nil {
		t.Errorf("diagram.create returns err: %q", err)
	}
	if resp == nil {
		t.Error("response should be nil, but got nil")
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("status code should equal to %d, but got: %d", http.StatusOK, resp.StatusCode)
	}
}
func TestDiagramDeleteWithWrongID(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/diagrams/00e77f4dc9973517/delete.json", func(w http.ResponseWriter, r *http.Request) {
		method(t, r, "GET")
		fmt.Fprint(w, `{}`)
	})

	resp, err := client.Diagram.Delete(context.Background(), "12345")
	if err == nil {
		t.Error("error should not be nil")
	}

	if resp.StatusCode != 404 {
		t.Errorf("response code should be:%d, but got:%d", 404, resp.StatusCode)
	}
}

func TestDiagramGetToken(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/diagrams/00e77f4dc9973517/editor/token.json", func(w http.ResponseWriter, r *http.Request) {
		method(t, r, "GET")
		fmt.Fprint(w, `{
		    "token": "nITfHr0rxfeAxdCn"
		}`)
	})

	token, _, err := client.Diagram.GetToken(context.Background(), "00e77f4dc9973517")
	if err != nil {
		t.Errorf("token.gettoken returns err: %q", err)
	}

	want := testToken(t)

	if !reflect.DeepEqual(token, want) {
		t.Errorf("want: %v, but got: %v", want, token)
	}
}
func TestDiagramGetTokenWithWrongID(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/diagrams/00e77f4dc9973517/editor/token.json", func(w http.ResponseWriter, r *http.Request) {
		method(t, r, "GET")
		fmt.Fprint(w, `{
		    "token": "nITfHr0rxfeAxdCn"
		}`)
	})

	token, _, err := client.Diagram.GetToken(context.Background(), "12345600e77f4dc97")
	if err == nil {
		t.Error("error should not be nil", err)
	}

	if token != nil {
		t.Errorf("token should be nil, but got: %v", token)
	}

}

func testAccount() *Account {
	return &Account{
		Name:     String("1MUJPfNEEeVUox15"),
		Nickname: String("Yoko"),
		Type:     String("cacoo"),
		ImageURL: String("https://cacoo.com/account/1MUJPfNEEeVUox15/image/32x32"),
	}
}
func testDiagramResult(t *testing.T) DiagramResult {
	return DiagramResult{
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
		Owner:          testAccount(),
		Editing:        Bool(true),
		Own:            Bool(true),
		Shared:         Bool(false),
		FolderID:       Int(10001),
		FolderName:     String("Design"),
		SheetCount:     Int(3),
		Created:        testTime(t),
		Updated:        testTime(t),
	}
}
func testSheet() Sheet {
	return Sheet{
		Height:         Int(100),
		ImageURL:       String("https://cacoo.com/diagrams/00e77f4dc9973517-aaaa.png"),
		ImageURLForAPI: String("https://cacoo.com/api/v1/diagrams/00e77f4dc9973517-aaaa.png"),
		Name:           String("sheet1"),
		UID:            String("aaaa"),
		URL:            String("https://cacoo.com/diagrams/00e77f4dc9973517#aaaa"),
		Width:          Int(200),
	}
}
func testDiagramResultWithSheetAndComment(t *testing.T) *DiagramResult {
	comment := testComment(t)

	tdr := testDiagramResult(t)
	dr := &DiagramResult{}
	dr = &tdr
	dr.Comments = append(dr.Comments, comment)
	dr.Sheets = append(dr.Sheets, testSheet())

	return dr
}

func testComment(t *testing.T) Comment {
	return Comment{
		User:    testAccount(),
		Content: String("comment 1"),
		Updated: testTime(t),
		Created: testTime(t),
	}
}

func testTime(t *testing.T) *CacooTime {
	testTime, err := time.Parse(dateFormat, "Mon, 10 Aug 2009 17:00:00 +0900")
	if err != nil {
		t.Errorf("error should be nil but got: %q", err)
	}

	return &CacooTime{testTime}
}

func testDiagrams(t *testing.T) *Diagrams {
	return &Diagrams{
		Result: []DiagramResult{testDiagramResult(t)},
		Count:  Int(1),
	}
}
func testToken(t *testing.T) *Editor {
	return &Editor{
		Token: String("nITfHr0rxfeAxdCn"),
	}
}

func testGenerateDiagramOption(t *testing.T) *DiagramOption {
	return &DiagramOption{
		FolderID:    10001,
		Title:       "Wireframe",
		Description: "Current project wireframe",
		Security:    "url",
	}
}
