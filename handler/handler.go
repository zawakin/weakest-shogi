package handler

import (
	"context"
	"errors"
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"

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

func handleFirstUSI(ctx context.Context, cmds []string) {
	log.Debugf("FirstUSI called; cmds=%#v", cmds)

	res := &usi.FirstUSIResponse{
		Name:   "weakest-shogi",
		Author: "zawawahoge",
	}
	fmt.Printf("id name %s\n", res.Name)
	fmt.Printf("id author %s\n", res.Author)
	fmt.Println("usiok")
}

func handleIsReady(ctx context.Context, cmds []string) {
	log.Debugf("IsReady called; cmds=%#v", cmds)

	fmt.Println("readyok")
}

func handleNewGame(ctx context.Context, cmds []string) {
	log.Debugf("NewGame called; cmds=%#v", cmds)

	fmt.Println("readyok")
}
