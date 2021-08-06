package handler

import (
	"context"

	log "github.com/sirupsen/logrus"
)

func (h *baseUSIHandler) handleGo(ctx context.Context, cmds []string) {
	log.Debugf("Go called; cmds=%#v", cmds)

	h.Println("bestmove resign")
}
