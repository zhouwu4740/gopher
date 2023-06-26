package example

import (
	"log"
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	got := Split("a/b/c", "/")
	want := []string{"a", "b", "c"}
	if !reflect.DeepEqual(want, got) {
		log.Fatalf("expected: %v, got %v\n", want, got)
	}
}
