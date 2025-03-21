package application

import (
	"context"
	"fmt"
	"github.com/Rasikrr/learning_platform_core/brokers/nats"
	"log"
)

func (a *App) initNats(_ context.Context) error {
	cfg := a.Config().NATS
	if !cfg.Required {
		log.Println("NATS do not required, skip")
		return nil
	}
	var err error
	a.publisher, err = nats.NewPublisher(cfg.DSN)
	if err != nil {
		return fmt.Errorf("init NATS error: %w", err)
	}
	a.subscriber, err = nats.NewSubscriber(cfg.DSN, nats.WithQueue(cfg.Queue))
	if err != nil {
		return fmt.Errorf("init NATS error: %w", err)
	}

	a.starters.Add(a.subscriber)
	a.closers.Add(a.subscriber)

	a.closers.Add(a.publisher)
	return nil
}
