package command

import (
	"reflect"
	"testing"
)

// This is the directory where our test fixtures are.
func TestCommond_Version(t *testing.T) {
	command := &VersionCommand{
		Version:        "1.0",
		AppName:        "test app",
		LibraryVersion: "0.1",
	}

	wantSynopsis := "Prints the test app version"
	if !reflect.DeepEqual(command.Synopsis(), wantSynopsis) {
		t.Errorf("VersionCommand.Synopsis returned %+v, want %+v", command.Synopsis(), wantSynopsis)
	}

	if !reflect.DeepEqual(command.Help(), "") {
		t.Errorf("VersionCommand.Help returned %+v, want %+v", command.Help(), "")
	}

	args := []string{}

	if !reflect.DeepEqual(command.Run(args), 1) {
		t.Errorf("VersionCommand.Run returned %+v, want %+v", command.Run(args), 1)
	}
}
