package lib

import (
	"testing"

	"github.com/WalkerEpps/quick_serve/lib"
)

func TestLoad(t *testing.T) {
	config := lib.LoadConfig()
	if config == nil {
		t.Error("Error while loading")
	}
	t.Log("config: ", config)
}
