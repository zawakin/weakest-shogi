package handler

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"
)

func handleIsReady(ctx context.Context, cmds []string) {
	log.Debugf("IsReady called; cmds=%#v", cmds)

	fmt.Println("readyok")
}
