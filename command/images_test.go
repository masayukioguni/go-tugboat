package command

import (
	"reflect"
	"testing"
)

func TestCommond_Image(t *testing.T) {
	command := &ImagesCommand{
		Client: nil,
	}

	wantSynopsis := "Retrieve a image list"
	if !reflect.DeepEqual(command.Synopsis(), wantSynopsis) {
		t.Errorf("ImagesCommand.Synopsis returned %+v, want %+v", command.Synopsis(), wantSynopsis)
	}
}
