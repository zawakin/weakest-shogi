package handler

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"
)

func handleNewGame(ctx context.Context, cmds []string) {
	log.Debugf("NewGame called; cmds=%#v", cmds)

	fmt.Println("readyok")
}
