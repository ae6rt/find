package find

import (
	"regexp"
	"testing"
)

func TestFind(t *testing.T) {
	arr, err := Find("testdata", regexp.MustCompile(`cert\.pem`), 2, []string{".dot"})
	if err != nil {
		t.Fatal(err)
	}
	if len(arr) != 1 {
		t.Fatalf("Want 1 but got %d\n", len(arr))
	}
	if arr[0] != "testdata/acme.com/cert.pem" {
		t.Fatalf("Want testdata/acme.com/cert.pem but got %s\n", arr[0])
	}
}

func TestBadRoot(t *testing.T) {
	_, err := Find("testdata/", regexp.MustCompile(`cert\.pem`), 2, []string{".dot"})
	if err == nil {
		t.Fatalf("Want an error because root ends in /")
	}
}

func TestNoDepth(t *testing.T) {
	arr, err := Find("testdata", regexp.MustCompile(`cert\.pem`), -1, []string{".dot"})
	if err != nil {
		t.Fatal(err)
	}
	if len(arr) != 3 {
		t.Fatalf("Want 3 but got %d\n", len(arr))
	}
}
