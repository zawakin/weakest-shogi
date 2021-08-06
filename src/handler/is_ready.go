package handler

import (
	"context"

	log "github.com/sirupsen/logrus"
)

func (h *baseUSIHandler) handleIsReady(ctx context.Context, cmds []string) {
	log.Debugf("IsReady called; cmds=%#v", cmds)

	h.Println("readyok")
}
