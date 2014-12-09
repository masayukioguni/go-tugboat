package command

import (
	"reflect"
	"testing"
)

func TestCommond_Account(t *testing.T) {
	command := &AccountCommand{
		Client: nil,
	}

	wantSynopsis := "Show account information"
	if !reflect.DeepEqual(command.Synopsis(), wantSynopsis) {
		t.Errorf("AccountCommand.Synopsis returned %+v, want %+v", command.Synopsis(), wantSynopsis)
	}

	if !reflect.DeepEqual(command.Help(), "") {
		t.Errorf("AccountCommand.Help returned %+v, want %+v", command.Help(), "")
	}

}
