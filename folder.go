package cacoo

import (
	"context"
	"fmt"
)

// FolderService handles folder related methods
type FolderService service

// Folder represents the folder list of the account
type Folder struct {
	Results []FolderResult `json:"result"`
}

// FolderResult represents the list of the folder information
type FolderResult struct {
	Created    CacooTime `json:"created"`
	FolderID   int       `json:"folderId"`
	FolderName string    `json:"folderName"`
	Type       string    `json:"type"`
	Updated    CacooTime `json:"updated"`
}

// MyFolders returns the folder list of the authenticated user
func (fs *FolderService) MyFolders(ctx context.Context) (*Folder, *Response, error) {
	u := fmt.Sprintf("%s", "folders.json")

	f := new(Folder)

	resp, err := fs.client.Get(ctx, u, &f)
	if err != nil {
		return nil, resp, err
	}

	return f, resp, nil
}
