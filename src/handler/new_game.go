package handler

import (
	"context"

	log "github.com/sirupsen/logrus"
)

func (h *baseUSIHandler) handleNewGame(ctx context.Context, cmds []string) {
	log.Debugf("NewGame called; cmds=%#v", cmds)
}
