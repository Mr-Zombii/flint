package typechecker

import "maps"

type Env struct {
	vars    map[string]*Type
	parent  *Env
	modules map[string]*Env
}

func NewEnv(parent *Env) *Env {
	modules := make(map[string]*Env)
	if parent != nil && parent.modules != nil {
		maps.Copy(modules, parent.modules)
	}
	return &Env{
		vars:    make(map[string]*Type),
		parent:  parent,
		modules: modules,
	}
}

func (e *Env) Get(name string) (*Type, bool) {
	if ty, ok := e.vars[name]; ok {
		return ty, true
	}
	if e.parent != nil {
		return e.parent.Get(name)
	}
	return nil, false
}

func (e *Env) Set(name string, ty *Type) {
	e.vars[name] = ty
}
