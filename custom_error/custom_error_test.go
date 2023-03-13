package custom_error

import (
	"net/http"
	"testing"
)

func TestNewNoContentFoundError(t *testing.T) {
	want := CustomError{
		Code: http.StatusNoContent,
		Msg:  MsgNoDataFound,
	}
	custom_err := NewNoContentFoundError("test", nil, MsgNoDataFound)

	got, ok := custom_err.(*CustomError)
	if !ok {
		t.Fatal("Failed in type assertion.")
	}
	if got.Code != want.Code {
		t.Fatalf("want \"%d\", but got \"%d\"", want.Code, got.Code)
	}
	if got.Msg != want.Msg {
		t.Fatalf("want \"%s\", but got \"%s\"", want.Msg, got.Msg)
	}
}
