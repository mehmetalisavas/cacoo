package cacoo

import "testing"

func TestCacooTime(t *testing.T) {
	tTime := []byte(`"Mon, 10 Aug 2009 17:00:00 +0900"`)

	ct := &CacooTime{}
	err := ct.UnmarshalJSON(tTime)
	if err != nil {
		t.Errorf("error while parsing cacoo time, got: %q", err)
	}
}

func TestCacooTimeString(t *testing.T) {
	tTime := []byte(`"Mon, 10 Aug 2009 17:00:00 +0900"`)

	ct := &CacooTime{}
	err := ct.UnmarshalJSON(tTime)
	if err != nil {
		t.Errorf("error while parsing cacoo time, got: %q", err)
	}

	if ct.String() == "" {
		t.Error("time string should not be empty")
	}
}
