package config

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestConfig_Path(t *testing.T) {
	path, err := GetConfigPath()
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	expected := filepath.Join(os.Getenv("HOME"), ".go-tugboat/config.yaml")
	if !reflect.DeepEqual(path, expected) {
		t.Fatalf("bad: %#v %#v", path, expected)
	}
}

func TestConfig_Load(t *testing.T) {
	c, err := LoadConfig(filepath.Join("./test-fixtures", "config.yaml"))
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	want := &Config{
		Authentication{
			APIKey: "APIKey",
		},
		Defaluts{
			Region: "nyc3",
		},
	}

	if !reflect.DeepEqual(c, want) {
		t.Errorf("LoadConfig is %v, want %v", c, want)
	}
}

func TestConfig_GetDefaultDirectory(t *testing.T) {
	want := ".go-tugboat"
	if !reflect.DeepEqual(GetDefaultDirectory(), want) {
		t.Errorf("LoadConfig is %v, want %v", GetDefaultDirectory(), want)
	}

}

func TestConfig_GetDefaultConfigName(t *testing.T) {
	want := "config.yaml"
	if !reflect.DeepEqual(GetDefaultConfigName(), want) {
		t.Errorf("LoadConfig is %v, want %v", GetDefaultConfigName(), want)
	}

}
