package cacoo

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"
)

func TestDiagramImage(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	file, err := os.Open("testdata/cacoo_test.png")
	if err != nil {
		t.Errorf("image not found, image should exists as: %s", "cacoo_test.png")
	}
	defer file.Close()

	mux.HandleFunc("/diagrams/00e77f4dc9973517.png", func(w http.ResponseWriter, r *http.Request) {
		method(t, r, "GET")
		fmt.Fprint(w)
		io.Copy(w, file)
	})

	filePath := "testdata/00e77f4dc9973517.png"
	resp, err := client.Image.DownloadDiagram(context.Background(), "00e77f4dc9973517", filePath)
	if err != nil {
		t.Errorf("image.downloaddigram returned error: %v", err)
	}

	if resp == nil {
		t.Error("Downloaded diagram should not be nil")
	}

	downloadedFile, err := os.Open(filePath)
	if err != nil {
		t.Error("error should be nil while opening downloaded file path")
	}
	fs, err := file.Stat()
	if err != nil {
		t.Error("error while getting stats for file")
	}
	dfs, err := downloadedFile.Stat()
	if err != nil {
		t.Error("error while getting stats for downloaded file")
	}

	// sizes of 2 file should be the same
	if fs.Size() != dfs.Size() {
		t.Errorf("sizes should be the same; but file stats: %d , downloaded file stats: %d", fs.Size(), dfs.Size())
	}
}
