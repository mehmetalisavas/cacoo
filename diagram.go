package cacoo

import (
	"context"
	"fmt"
)

// DiagramService handles diagrams related methods
type DiagramService service

// Diagrams respresents the information about diagram list
type Diagrams struct {
	Count  *int            `json:"count,omitempty"`
	Result []DiagramResult `json:"result,omitempty"`
}

// DiagramResult respresents the inforamation about the diagram and its owner
type DiagramResult struct {
	Description      *string  `json:"description,omitempty"`
	DiagramID        *string  `json:"diagramId,omitempty"`
	Editing          *bool    `json:"editing,omitempty"`
	FolderID         *int     `json:"folderId,omitempty"`
	FolderName       *string  `json:"folderName,omitempty"`
	ImageURL         *string  `json:"imageUrl,omitempty"`
	ImageURLForAPI   *string  `json:"imageUrlForApi,omitempty"`
	OrganizationKey  *string  `json:"organizationKey,omitempty"`
	OrganizationName *string  `json:"organizationName,omitempty"`
	Own              *bool    `json:"own,omitempty"`
	Owner            *Account `json:"owner,omitempty"`
	OwnerName        *string  `json:"ownerName,omitempty"`
	OwnerNickname    *string  `json:"ownerNickname,omitempty"`
	// string ? or any integer value ?
	ProjectID   interface{} `json:"projectId,omitempty"`
	ProjectName *string     `json:"projectName,omitempty"`
	Security    *string     `json:"security,omitempty"`
	Shared      *bool       `json:"shared,omitempty"`
	Sheets      []Sheet     `json:"sheets,omitempty"`
	SheetCount  *int        `json:"sheetCount,omitempty"`
	Title       *string     `json:"title,omitempty"`
	Type        *string     `json:"type,omitempty"`
	Updated     *CacooTime  `json:"updated,omitempty"`
	Created     *CacooTime  `json:"created,omitempty"`
	URL         *string     `json:"url,omitempty"`
	Comments    []Comment   `json:"comments,omitempty"`
}

// Sheet represents the information about hte sheets
type Sheet struct {
	Height         *int    `json:"height,omitempty"`
	ImageURL       *string `json:"imageUrl,omitempty"`
	ImageURLForAPI *string `json:"imageUrlForApi,omitempty"`
	Name           *string `json:"name,omitempty"`
	UID            *string `json:"uid,omitempty"`
	URL            *string `json:"url,omitempty"`
	Width          *int    `json:"width,omitempty"`
}

// ListDiagrams lists the diagrams of the account
func (ds *DiagramService) ListDiagrams(ctx context.Context) (*Diagrams, *Response, error) {
	u := fmt.Sprintf("%s", "diagrams.json")

	d := new(Diagrams)

	resp, err := ds.client.Get(ctx, u, &d)
	if err != nil {
		return nil, resp, err
	}

	return d, resp, nil
}

// GetDiagram fetches information about the given diagram name
func (ds *DiagramService) GetDiagram(ctx context.Context, id string) (*DiagramResult, *Response, error) {
	u := fmt.Sprintf("diagrams/%s.json", id)

	d := new(DiagramResult)

	resp, err := ds.client.Get(ctx, u, &d)
	if err != nil {
		return nil, resp, err
	}

	return d, resp, nil
}

// DiagramOption represents the options for the diagram creation
type DiagramOption struct {
	FolderID    int    `json:"folderId,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Security    string `json:"security,omitempty"`
}

// Create creates the diagram with given parameter
func (ds *DiagramService) Create(ctx context.Context, opt *DiagramOption) (*DiagramResult, *Response, error) {
	u := fmt.Sprintf("%s", "diagrams/create.json")

	d := new(DiagramResult)

	resp, err := ds.client.Post(ctx, u, opt, &d)
	if err != nil {
		return nil, resp, err
	}

	return d, resp, nil
}

// Get gives information about the diagram
func (ds *DiagramService) Get(ctx context.Context) (*DiagramResult, *Response, error) {
	u := fmt.Sprintf("%s", "diagrams/create.json")

	d := new(DiagramResult)

	resp, err := ds.client.Get(ctx, u, &d)
	if err != nil {
		return nil, resp, err
	}

	return d, resp, nil
}

// Delete deletes the diagram with given id
func (ds *DiagramService) Delete(ctx context.Context, diagramID string) (*Response, error) {
	u := fmt.Sprintf("diagrams/%s/delete.json", diagramID)

	resp, err := ds.client.Get(ctx, u, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Editor represents the diagram editor token
type Editor struct {
	Token *string `json:"token,omitempty"`
}

// GetToken fetches the token for given diagram
func (ds *DiagramService) GetToken(ctx context.Context, diagramID string) (*Editor, *Response, error) {
	u := fmt.Sprintf("diagrams/%s/editor/token.json", diagramID)

	e := new(Editor)

	resp, err := ds.client.Get(ctx, u, &e)
	if err != nil {
		return e, resp, err
	}

	return e, resp, nil
}
