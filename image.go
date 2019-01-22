package cacoo

import (
	"context"
	"fmt"
	"os"
)

// ImageService handles diagram image related methods
type ImageService service

// DownloadDiagram fetches the image with given diagram id
func (as *ImageService) DownloadDiagram(ctx context.Context, diagramID string, filePath string) (*Response, error) {

	u := fmt.Sprintf("diagrams/%s.png", diagramID)
	diagramFile := fmt.Sprintf("%s", filePath)

	// if there is and error on http level, remove created file
	var err error
	defer func() {
		if err != nil {
			os.Remove(diagramFile)
		}
	}()

	// Create the file
	out, err := os.Create(diagramFile)
	if err != nil {
		return nil, err
	}
	defer out.Close()

	resp, err := as.client.Get(ctx, u, out)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
