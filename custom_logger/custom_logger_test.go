package custom_logger

import "testing"

func TestInfo(t *testing.T) {
	wantString := "logging with string type."
	wantInt := 123
	// On testing, root directory is custom_logger, not box-ratings-lib.
	// So, output file will be created within custom_logger directory. Be aware.
	Info(wantString, wantInt)
}
