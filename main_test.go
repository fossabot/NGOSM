package main

import (
	"testing"
)

func TestVersion(t *testing.T) {
	var version = getVersion()

	if version != "0.0.1" {
		t.Error("Version ", version, " isn't correct")
	}

}
