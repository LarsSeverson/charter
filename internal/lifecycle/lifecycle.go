package lifecycle

import (
	"context"
	"errors"
	"log/slog"
	"time"
)

type Component interface {
	Name() string
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
}

type Lifecycle struct {
	components  []Component
	startCount  int
	failChannel chan error
	logger      *slog.Logger
}

func New() *Lifecycle {
	l := Lifecycle{
		failChannel: make(chan error, 1),
		logger:      slog.Default(),
	}
	return &l
}

func (l *Lifecycle) Add(c Component) error {
	l.components = append(l.components, c)
	return nil
}

func (l *Lifecycle) Fail(err error) {
	select {
	case l.failChannel <- err:
	default:
	}
}

func (l *Lifecycle) Failed() <-chan error {
	return l.failChannel
}

func (l *Lifecycle) Start(ctx context.Context) error {

	for _, c := range l.components {
		startedAt := time.Now()

		err := c.Start(ctx)
		if err != nil {
			startErr := errors.Join(err, ErrComponentStart)
			stopErr := c.Stop(ctx)
			err := errors.Join(startErr, stopErr)
			return err
		}

		l.startCount += 1

		l.logger.InfoContext(
			ctx,
			"component started",
			slog.String("component", c.Name()),
			slog.Duration("duration", time.Since(startedAt)),
		)
	}

	return nil
}

func (l *Lifecycle) Stop(ctx context.Context) error {
	var errs []error

	for i := l.startCount - 1; i >= 0; i -= 1 {
		c := l.components[i]
		stoppedAt := time.Now()

		err := c.Stop(ctx)
		if err != nil {
			errs = append(errs, err, ErrComponentStop)
			continue
		}

		l.logger.InfoContext(
			ctx,
			"component stopped",
			slog.String("component", c.Name()),
			slog.Duration("duration", time.Since(stoppedAt)),
		)
	}

	l.startCount = 0

	err := errors.Join(errs...)
	return err
}

func (l *Lifecycle) Run(ctx context.Context, stopTimeout time.Duration) error {
	startErr := l.Start(ctx)
	if startErr != nil {
		stopErr := l.stopWithTimeout(stopTimeout)
		err := errors.Join(startErr, stopErr)
		return err
	}

	var runErr error

	select {
	case <-ctx.Done():
	case runErr = <-l.Failed():
	}

	stopErr := l.stopWithTimeout(stopTimeout)
	err := errors.Join(runErr, stopErr)
	return err
}

func (l *Lifecycle) stopWithTimeout(timeout time.Duration) error {
	stopCtx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	err := l.Stop(stopCtx)
	return err
}
