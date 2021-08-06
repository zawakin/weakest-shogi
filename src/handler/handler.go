package handler

import (
	"context"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/zawawahoge/weakest-shogi/usicommand"
)

// USIHandler is a handler of USI command.
type USIHandler interface {
	HandleCommand(ctx context.Context, command string) (bool, error)
}

type baseUSIHandler struct {
	writer io.Writer
}

// NewUSIHandler is constructor of USI handler.
func NewUSIHandler(writer io.Writer) USIHandler {
	return &baseUSIHandler{
		writer: writer,
	}
}

// HandleUSICommand handles usi command.
func (h *baseUSIHandler) HandleCommand(ctx context.Context, command string) (bool, error) {
	cmds := strings.Fields(command)
	if len(cmds) == 0 {
		return true, errors.New("command is empty")
	}

	first := cmds[0]
	switch usicommand.Command(first) {
	case usicommand.CommandFirstUSI:
		h.handleFirstUSI(ctx, cmds)
	case usicommand.CommandIsReady:
		h.handleIsReady(ctx, cmds)
	case usicommand.CommandNewGame:
		h.handleNewGame(ctx, cmds)
	case usicommand.CommandGo:
		h.handleGo(ctx, cmds)
	case usicommand.CommandQuit:
		return false, nil
	default:
		return true, fmt.Errorf("invalid command; cmd = %s", command)
	}

	return true, nil
}

func (h *baseUSIHandler) Printf(format string, a ...interface{}) {
	io.WriteString(h.writer, fmt.Sprintf(format, a...))
}

func (h *baseUSIHandler) Println(s string) {
	io.WriteString(h.writer, fmt.Sprintf("%s\n", s))
}
