// autogenerated: do not edit!
// generated from gentemplate [gentemplate -d Package=event -id actor -d VecType=ActorVec -d Type=Actor github.com/platinasystems/go/elib/vec.tmpl]

package event

import (
	"github.com/platinasystems/go/elib"
)

type ActorVec []Actor

func (p *ActorVec) Resize(n uint) {
	c := elib.Index(cap(*p))
	l := elib.Index(len(*p)) + elib.Index(n)
	if l > c {
		c = elib.NextResizeCap(l)
		q := make([]Actor, l, c)
		copy(q, *p)
		*p = q
	}
	*p = (*p)[:l]
}

func (p *ActorVec) validate(new_len uint, zero *Actor) *Actor {
	c := elib.Index(cap(*p))
	lʹ := elib.Index(len(*p))
	l := elib.Index(new_len)
	if l <= c {
		// Need to reslice to larger length?
		if l >= lʹ {
			*p = (*p)[:l]
		}
		return &(*p)[l-1]
	}
	return p.validateSlowPath(zero, c, l, lʹ)
}

func (p *ActorVec) validateSlowPath(zero *Actor,
	c, l, lʹ elib.Index) *Actor {
	if l > c {
		cNext := elib.NextResizeCap(l)
		q := make([]Actor, cNext, cNext)
		copy(q, *p)
		if zero != nil {
			for i := c; i < cNext; i++ {
				q[i] = *zero
			}
		}
		*p = q[:l]
	}
	if l > lʹ {
		*p = (*p)[:l]
	}
	return &(*p)[l-1]
}

func (p *ActorVec) Validate(i uint) *Actor {
	return p.validate(i+1, (*Actor)(nil))
}

func (p *ActorVec) ValidateInit(i uint, zero Actor) *Actor {
	return p.validate(i+1, &zero)
}

func (p *ActorVec) ValidateLen(l uint) (v *Actor) {
	if l > 0 {
		v = p.validate(l, (*Actor)(nil))
	}
	return
}

func (p *ActorVec) ValidateLenInit(l uint, zero Actor) (v *Actor) {
	if l > 0 {
		v = p.validate(l, &zero)
	}
	return
}

func (p ActorVec) Len() uint { return uint(len(p)) }
