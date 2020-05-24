package handler

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/zawawahoge/weakest-shogi/usi"
)

// USIHandler is a handler of USI command.
type USIHandler interface {
	HandleCommand(ctx context.Context, command string) (bool, error)
}

type baseUSIHandler struct {
}

// NewUSIHandler is constructor of USI handler.
func NewUSIHandler() USIHandler {
	return &baseUSIHandler{}
}

// HandleUSICommand handles usi command.
func (h *baseUSIHandler) HandleCommand(ctx context.Context, command string) (bool, error) {
	cmds := strings.Fields(command)
	if len(cmds) == 0 {
		return true, errors.New("Command is empty")
	}

	first := cmds[0]
	switch usi.Command(first) {
	case usi.CommandFirstUSI:
		handleFirstUSI(ctx, cmds)
	case usi.CommandIsReady:
		handleIsReady(ctx, cmds)
	case usi.CommandQuit:
		return false, nil
	default:
		return true, fmt.Errorf("Invalid command; cmd = %s", command)
	}

	return true, nil
}
