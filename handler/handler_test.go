package handler

import (
	"bytes"
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_USIHandler_HandleCommand_FirstUSI(t *testing.T) {
	cmd := "usi"
	want := `id name weakest-shogi
id author zawawahoge
usiok
`

	stdout := new(bytes.Buffer)
	h := NewUSIHandler(stdout)
	h.HandleCommand(context.Background(), cmd)

	got := string(stdout.Bytes())

	if got != want {
		t.Errorf("h.HandleCommand(ctx,%s); got=%s; want=%s\ncmp.Diff(got,want)=%s", cmd, got, want, cmp.Diff(got, want))
	}
}

func Test_USIHandler_HandleCommand_IsReady(t *testing.T) {
	cmd := "isready"
	want := "readyok\n"

	stdout := new(bytes.Buffer)
	h := NewUSIHandler(stdout)
	h.HandleCommand(context.Background(), cmd)

	got := string(stdout.Bytes())

	if got != want {
		t.Errorf("h.HandleCommand(ctx,%s); got=%s; want=%s\ncmp.Diff(got,want)=%s", cmd, got, want, cmp.Diff(got, want))
	}
}
