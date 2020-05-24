package main

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/zawawahoge/weakest-shogi/handler"
)

func main() {
	var err error

	handler := handler.NewUSIHandler()

	isContinued := true
	for isContinued {
		var command string
		fmt.Scanf("%s", &command)
		isContinued, err = handler.HandleCommand(context.Background(), command)
		if err != nil {
			log.Warnf(err.Error())
		}
	}
}
