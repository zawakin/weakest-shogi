package handler

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/zawawahoge/weakest-shogi/usi"
)

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
