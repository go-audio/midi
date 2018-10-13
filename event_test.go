package midi

import (
	"errors"
	"reflect"
	"testing"
)

func TestEventCopy(t *testing.T) {
	event := &Event{}
	event.TimeSignature = &TimeSignature{}
	event.SmpteOffset = &SmpteOffset{}
	if !reflect.DeepEqual(event, event.Copy()) {
		t.Fatal(errors.New("Expected copy to be equal"))
	}
}

func TestEventString(t *testing.T) {
	event := Event{}
	expect := "Ch 0 @ 0 (0) \t0X0"
	str := event.String()
	if str != expect {
		t.Errorf("Expected '%s' got '%s'", expect, str)
	}
}
