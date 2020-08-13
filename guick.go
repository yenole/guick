package guick

type Guick struct {
	cfg     *Config
	modules []interface{}
}

func New() *Guick {
	return &Guick{}
}

func (g *Guick) Run(modules ...interface{}) (err error) {
	g.cfg, err = LoadConfig()
	if err != nil {
		return err
	}
	err = g.enroll(modules)
	if err != nil {
		return err
	}

	err = g.inject()
	if err != nil {
		return err
	}

	return nil
}

func (g *Guick) enroll(modules []interface{}) error {
	for _, module := range modules {
		err := tryInitialize(module, g.cfg)
		if err != nil {
			return err
		}
		g.modules = append(g.modules, module)
	}
	return nil
}

func (g *Guick) inject() error {
	miss := make([]interface{}, 0, len(g.modules))
	for _, module := range g.modules {
		err := tryInject(module, g)
		if err != nil && err != ErrNotInterface {
			miss = append(miss, module)
		}
	}
	return nil
}
