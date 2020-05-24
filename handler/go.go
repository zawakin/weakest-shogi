package handler

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"
)

func handleGo(ctx context.Context, cmds []string) {
	log.Debugf("Go called; cmds=%#v", cmds)

	fmt.Println("bestmove resign")
}
