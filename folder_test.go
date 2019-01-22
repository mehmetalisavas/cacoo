package cacoo

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestMyFolders(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/folders.json", func(w http.ResponseWriter, r *http.Request) {
		method(t, r, "GET")
		fmt.Fprint(w, `{
		    "result": [
		        {
		            "folderId": 10001,
		            "folderName": "Folder 1",
		            "type": "normal",
		            "created": "Mon, 10 Aug 2009 17:00:00 +0900",
		            "updated": "Mon, 10 Aug 2009 17:00:00 +0900"
		        }
				]
		}`)
	})

	folders, _, err := client.Folders.MyFolders(context.Background())
	if err != nil {
		t.Errorf("folders.myfolders has error: %v :", err)
	}

	want := testFolder(t)

	if !reflect.DeepEqual(folders, want) {
		t.Errorf("want: %v, but got: %v", want, folders)
	}
}

func testFolder(t *testing.T) *Folder {
	f := &Folder{}
	f.Results = append(f.Results, testFolderResult(t))

	return f
}

func testFolderResult(t *testing.T) FolderResult {
	return FolderResult{
		FolderID:   Int(10001),
		FolderName: String("Folder 1"),
		Type:       String("normal"),
		Created:    testTime(t),
		Updated:    testTime(t),
	}
}
