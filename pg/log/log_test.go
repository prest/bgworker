package log

import "testing"

func TestLog(t *testing.T) {
	if err := Info("Test Message"); err != nil {
		t.Error(err)
	}
}
