package subscriber

import (
	"context"

	"github.com/jinbanglin/go-log"

	example "github.com/jinbanglin/examples/template/fnc/proto/example"
)

type Example struct{}

func (e *Example) Handle(ctx context.Context, msg *example.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}
