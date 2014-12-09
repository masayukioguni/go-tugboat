package command

import (
	"reflect"
	"testing"
)

func TestCommond_SSHKey(t *testing.T) {
	command := &SSHKeyCommand{
		Client: nil,
	}

	wantSynopsis := "Show available SSH keys"
	if !reflect.DeepEqual(command.Synopsis(), wantSynopsis) {
		t.Errorf("AccountCommand.Synopsis returned %+v, want %+v", command.Synopsis(), wantSynopsis)
	}

	if !reflect.DeepEqual(command.Help(), "") {
		t.Errorf("AccountCommand.Help returned %+v, want %+v", command.Help(), "")
	}

}
