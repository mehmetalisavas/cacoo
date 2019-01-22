package cacoo

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestLicense(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/account/license.json", func(w http.ResponseWriter, r *http.Request) {
		method(t, r, "GET")
		fmt.Fprint(w, `{
		    "plan": "free",
		    "remainingSheets": 25,
		    "remainingSharedFolders": 1,
		    "maxNumberOfSharersPerDiagram": 15,
		    "maxNumberOfSharersPerSharedFolder": 3,
		    "canCreateSheet": true,
		    "canCreateSharedFolder": true,
		    "canExportVectorFormat": false
		}`)
	})

	license, _, err := client.License.GetLicense(context.Background())
	if err != nil {
		t.Errorf("license.getlicense returned error: %v", err)
	}

	want := testLicense()

	if !reflect.DeepEqual(license, want) {
		t.Errorf("want: %v, but got: %v", want, license)
	}
}

func testLicense() *License {
	return &License{
		CanCreateSharedFolder:             Bool(true),
		CanCreateSheet:                    Bool(true),
		CanExportVectorFormat:             Bool(false),
		MaxNumberOfSharersPerDiagram:      Int(15),
		MaxNumberOfSharersPerSharedFolder: Int(3),
		Plan:                              String("free"),
		RemainingSharedFolders:            Int(1),
		RemainingSheets:                   Int(25),
	}
}
