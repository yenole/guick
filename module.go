package guick

import "errors"

var ErrNotInterface = errors.New("interface conversion failed")

type instantiater interface {
	Initialize(cfg *Config) error
}

func tryInitialize(module interface{}, cfg *Config) error {
	if v, ok := module.(instantiater); ok {
		return v.Initialize(cfg)
	}
	return ErrNotInterface
}

type injecter interface {
	Inject(guick *Guick) error
}

func tryInject(module interface{}, guick *Guick) error {
	if v, ok := module.(injecter); ok {
		return v.Inject(guick)
	}
	return ErrNotInterface
}
