package command

import (
	"reflect"
	"testing"
)

func TestCommond_Droplets(t *testing.T) {
	command := &DropletsCommand{
		Client: nil,
	}

	wantSynopsis := "Retrieve a droplet list"
	if !reflect.DeepEqual(command.Synopsis(), wantSynopsis) {
		t.Errorf("DropletsCommand.Synopsis returned %+v, want %+v", command.Synopsis(), wantSynopsis)
	}

	if !reflect.DeepEqual(command.Help(), "") {
		t.Errorf("DropletsCommand.Help returned %+v, want %+v", command.Help(), "")
	}

}
