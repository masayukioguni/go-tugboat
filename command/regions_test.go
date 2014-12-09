package command

import (
	"reflect"
	"testing"
)

func TestCommond_Region(t *testing.T) {
	command := &RegionsCommand{
		Client: nil,
	}

	wantSynopsis := "Show regions"
	if !reflect.DeepEqual(command.Synopsis(), wantSynopsis) {
		t.Errorf("RegionsCommand.Synopsis returned %+v, want %+v", command.Synopsis(), wantSynopsis)
	}

	if !reflect.DeepEqual(command.Help(), "") {
		t.Errorf("RegionsCommand.Help returned %+v, want %+v", command.Help(), "")
	}

}
