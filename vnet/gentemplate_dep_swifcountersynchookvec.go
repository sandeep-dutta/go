// autogenerated: do not edit!
// generated from gentemplate [gentemplate -id SwIfCounterSyncHookVec -d Package=vnet -d DepsType=SwIfCounterSyncHookVec -d Type=SwIfCounterSyncHook -d Data=hooks github.com/platinasystems/go/elib/dep/dep.tmpl]

package vnet

import (
	"github.com/platinasystems/go/elib/dep"
)

type SwIfCounterSyncHookVec struct {
	deps  dep.Deps
	hooks []SwIfCounterSyncHook
}

func (t *SwIfCounterSyncHookVec) Len() int {
	return t.deps.Len()
}

func (t *SwIfCounterSyncHookVec) Get(i int) SwIfCounterSyncHook {
	return t.hooks[t.deps.Index(i)]
}

func (t *SwIfCounterSyncHookVec) Add(x SwIfCounterSyncHook, ds ...*dep.Dep) {
	if len(ds) == 0 {
		t.deps.Add(&dep.Dep{})
	} else {
		t.deps.Add(ds[0])
	}
	t.hooks = append(t.hooks, x)
}
