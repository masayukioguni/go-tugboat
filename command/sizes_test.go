package command

import (
	"reflect"
	"testing"
)

func TestCommond_Sizes(t *testing.T) {
	command := &SizesCommand{
		Client: nil,
	}

	wantSynopsis := "Show available droplet sizes"
	if !reflect.DeepEqual(command.Synopsis(), wantSynopsis) {
		t.Errorf("SizesCommand.Synopsis returned %+v, want %+v", command.Synopsis(), wantSynopsis)
	}

	if !reflect.DeepEqual(command.Help(), "") {
		t.Errorf("VersionCommand.Help returned %+v, want %+v", command.Help(), "")
	}

}
