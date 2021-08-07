package handler

import (
	"bytes"
	"context"
	"errors"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func runSingleCommand(cmd string) (string, error) {
	stdout := new(bytes.Buffer)
	h := NewUSIHandler(stdout)
	_, err := h.HandleCommand(context.Background(), cmd)

	return stdout.String(), err
}

func Test_USIHandler_HandleCommand_FirstUSI(t *testing.T) {
	cmd := "usi"
	want := `id name weakest-shogi
id author zawawahoge
usiok
`
	got, err := runSingleCommand(cmd)
	if err != nil {
		t.Fatalf("h.HandleCommand(ctx,%s)=_, %#v; want nil", cmd, err)
	}
	if got != want {
		t.Errorf("h.HandleCommand(ctx,%s); got=%s; want=%s\ncmp.Diff(got,want)=%s", cmd, got, want, cmp.Diff(got, want))
	}
}

func Test_USIHandler_HandleCommand_IsReady(t *testing.T) {
	cmd := "isready"
	want := "readyok\n"

	got, err := runSingleCommand(cmd)
	if err != nil {
		t.Fatalf("h.HandleCommand(ctx,%s)=_, %#v; want nil", cmd, err)
	}
	if got != want {
		t.Errorf("h.HandleCommand(ctx,%s); got=%s; want=%s\ncmp.Diff(got,want)=%s", cmd, got, want, cmp.Diff(got, want))
	}
}

func Test_USIHandler_HandleCommand_NewGame(t *testing.T) {
	cmd := "usinewgame"
	want := ""

	got, err := runSingleCommand(cmd)
	if err != nil {
		t.Fatalf("h.HandleCommand(ctx,%s)=_, %#v; want nil", cmd, err)
	}
	if got != want {
		t.Errorf("h.HandleCommand(ctx,%s); got=%s; want=%s\ncmp.Diff(got,want)=%s", cmd, got, want, cmp.Diff(got, want))
	}
}

func Test_USIHandler_HandleCommand_Go(t *testing.T) {
	cmd := "go"
	want := "bestmove resign\n"

	got, err := runSingleCommand(cmd)
	if err != nil {
		t.Fatalf("h.HandleCommand(ctx,%s)=_, %#v; want nil", cmd, err)
	}
	if got != want {
		t.Errorf("h.HandleCommand(ctx,%s); got=%s; want=%s\ncmp.Diff(got,want)=%s", cmd, got, want, cmp.Diff(got, want))
	}
}

func Test_USIHandler_HandleCommand_Quit(t *testing.T) {
	cmd := "quit"
	want := false

	h := NewUSIHandler(os.Stdout)
	got, err := h.HandleCommand(context.Background(), cmd)
	if err != nil {
		t.Fatalf("h.HandleCommand(ctx,%s)=_, %#v; want nil", cmd, err)
	}
	if got != want {
		t.Errorf("h.HandleCommand(ctx,%s)=%#v, nil; want %v", cmd, got, want)
	}
}

func Test_USIHandler_HandleCommand_Empty(t *testing.T) {
	cmd := ""
	wantErr := errors.New("command is empty")

	h := NewUSIHandler(os.Stdout)
	_, err := h.HandleCommand(context.Background(), cmd)
	if err.Error() != wantErr.Error() {
		t.Errorf("h.HandleCommand(ctx,%s)=_, %#v; want=%v", cmd, err, wantErr)
	}
}
