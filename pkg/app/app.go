package app

import (
	"context"

	log "github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
)

//App represents Application-Object
type App struct {
	Services map[string]Service
}

//New creates new App
func New() *App {
	return &App{
		Services: make(map[string]Service),
	}
}

//AddService adds a background service to the application
func (a *App) AddService(name string, svc Service) error {
	if _, ok := a.Services[name]; ok {
		return xerrors.Errorf("Service %s already exists", name)
	}
	a.Services[name] = svc
	return nil
}

//GetService returns a Service for modification
func (a *App) GetService(name string) (Service, error) {
	if svc, ok := a.Services[name]; ok {
		return svc, nil
	}
	return nil, xerrors.New("Service not registered")
}

//Run starts all services in the background and returns as soon as one dies with its error message
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
