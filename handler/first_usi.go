package handler

import (
	"context"

	log "github.com/sirupsen/logrus"

	"github.com/zawawahoge/weakest-shogi/usi"
)

func (h *baseUSIHandler) handleFirstUSI(ctx context.Context, cmds []string) {
	log.Debugf("FirstUSI called; cmds=%#v", cmds)

	res := &usi.FirstUSIResponse{
		Name:   "weakest-shogi",
		Author: "zawawahoge",
	}
	h.Printf("id name %s\n", res.Name)
	h.Printf("id author %s\n", res.Author)
	h.Println("usiok")
}
