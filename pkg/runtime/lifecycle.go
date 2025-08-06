package runtime

import (
	"context"
	"fmt"
	"sync"

	"github.com/TradlyLabs/tradly-common/pkg/runtime/service"
)

var DefaultManager *Manager

func init() {
	DefaultManager = NewManager()
}

type Manager struct {
	services       []service.Service
	mu             sync.Mutex
	serviceIndexes map[string]int
	serviceNames   map[int]string
}

func NewManager() *Manager {
	m := &Manager{
		serviceIndexes: make(map[string]int),
		serviceNames:   make(map[int]string),
	}
	return m
}

func (lm *Manager) Register(name string, svc service.Service) {
	lm.mu.Lock()
	defer lm.mu.Unlock()
	index, exists := lm.serviceIndexes[name]
	if exists {
		lm.services[index] = svc
	}
	lm.serviceIndexes[name] = len(lm.services)
	lm.serviceNames[len(lm.services)] = name
	lm.services = append(lm.services, svc)
}

func Register(name string, svc service.Service) {
	DefaultManager.Register(name, svc)
}

func (lm *Manager) Start(ctx context.Context) error {
	var errors []error
	for index, svc := range lm.services {
		if err := svc.Start(ctx); err != nil {
			errors = append(errors, fmt.Errorf("failed to start service %s: %w", lm.serviceNames[index], err))
		}
	}
	if len(errors) == 1 {
		return errors[0]
	} else if len(errors) > 0 {
		return MultiError(errors)
	}
	return nil
}

func (lm *Manager) Stop(ctx context.Context) error {
	var errors []error
	for i := len(lm.services) - 1; i >= 0; i-- {
		if err := lm.services[i].Stop(ctx); err != nil {
			errors = append(errors, err)
		}
	}

	if len(errors) > 0 {
		return MultiError(errors)
	}
	return nil
}
