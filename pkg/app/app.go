package app

import (
	"context"

	log "github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
)

type App struct {
	Services map[string]Service
}

func New() *App {
	return &App{
		Services: make(map[string]Service),
	}
}

func (a *App) AddService(name string, svc Service) error {
	if _, ok := a.Services[name]; ok {
		return xerrors.Errorf("Service %s already exists", name)
	}
	a.Services[name] = svc
	return nil
}

func (a *App) GetService(name string) (Service, error) {
	if svc, ok := a.Services[name]; ok {
		return svc, nil
	}
	return nil, xerrors.New("Service not registered")
}
func (a *App) Run() error {
	mainCtx, cancel := context.WithCancel(context.Background())
	errChan := make(chan error)
	for name, svc := range a.Services {
		go func(name string, svc Service, ctx context.Context, err chan<- error, cancel func()) {
			log.Infof("Starting service %s...", name)
			err <- svc.Run(ctx)
			cancel()
		}(name, svc, mainCtx, errChan, cancel)
	}
	return <-errChan
}
