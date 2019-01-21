package cacoo

import (
	"context"
	"fmt"
	"testing"

	"github.com/bookerzzz/grok"
)

func TestFolder(t *testing.T) {
	c := NewClient(TestToken)
	// c := NewClient("asd")
	folders, _, err := c.Folders.MyFolders(context.Background())
	if err != nil {
		fmt.Println("err is:", err)
	}

	fmt.Println(folders)
	fmt.Println("folder id is:", folders.Results[0].FolderID)
	fmt.Println("folder id is:", folders.Results[0])
	grok.V(folders)
	_ = folders
}
