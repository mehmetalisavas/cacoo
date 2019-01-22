package cacoo

import (
	"context"
	"fmt"
)

// LicenseService handles license related methods
type LicenseService service

// License represents the information about account license
type License struct {
	CanCreateSharedFolder             *bool   `json:"canCreateSharedFolder,omitempty"`
	CanCreateSheet                    *bool   `json:"canCreateSheet,omitempty"`
	CanExportVectorFormat             *bool   `json:"canExportVectorFormat,omitempty"`
	MaxNumberOfSharersPerDiagram      *int    `json:"maxNumberOfSharersPerDiagram,omitempty"`
	MaxNumberOfSharersPerSharedFolder *int    `json:"maxNumberOfSharersPerSharedFolder,omitempty"`
	Plan                              *string `json:"plan,omitempty"`
	RemainingSharedFolders            *int    `json:"remainingSharedFolders,omitempty"`
	RemainingSheets                   *int    `json:"remainingSheets,omitempty"`
}

// GetLicense gives information about the account license
func (ls *LicenseService) GetLicense(ctx context.Context) (*License, *Response, error) {
	u := fmt.Sprintf("%s", "account/license.json")

	l := new(License)

	resp, err := ls.client.Get(ctx, u, &l)
	if err != nil {
		return nil, resp, err
	}

	return l, resp, nil
}
