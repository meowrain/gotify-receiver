package toml

import "testing"

func TestLoadConfig(t *testing.T) {
	_, err := LoadConfig()
	if err != nil {
		t.Fatal(err)
	}
}
