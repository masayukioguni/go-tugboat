package command

import (
	"reflect"
	"testing"
)

func TestCommond_Info(t *testing.T) {
	command := &InfoCommand{
		Client: nil,
	}

	wantSynopsis := "Show a droplet's information"
	if !reflect.DeepEqual(command.Synopsis(), wantSynopsis) {
		t.Errorf("InfoCommand.Synopsis returned %+v, want %+v", command.Synopsis(), wantSynopsis)
	}

	if reflect.DeepEqual(command.Help(), "") {
		t.Errorf("InfoCommand.Help returned %+v, want %+v", command.Help(), "")
	}
}
