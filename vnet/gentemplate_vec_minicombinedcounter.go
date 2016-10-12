// autogenerated: do not edit!
// generated from gentemplate [gentemplate -d Package=vnet -id miniCombinedCounter -d VecType=miniCombinedCounterVec -d Type=miniCombinedCounter github.com/platinasystems/go/elib/vec.tmpl]

package vnet

import (
	"github.com/platinasystems/go/elib"
)

type miniCombinedCounterVec []miniCombinedCounter

func (p *miniCombinedCounterVec) Resize(n uint) {
	c := elib.Index(cap(*p))
	l := elib.Index(len(*p)) + elib.Index(n)
	if l > c {
		c = elib.NextResizeCap(l)
		q := make([]miniCombinedCounter, l, c)
		copy(q, *p)
		*p = q
	}
	*p = (*p)[:l]
}

func (p *miniCombinedCounterVec) validate(new_len uint, zero *miniCombinedCounter) *miniCombinedCounter {
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

func (p *miniCombinedCounterVec) validateSlowPath(zero *miniCombinedCounter,
	c, l, lʹ elib.Index) *miniCombinedCounter {
	if l > c {
		cNext := elib.NextResizeCap(l)
		q := make([]miniCombinedCounter, cNext, cNext)
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

func (p *miniCombinedCounterVec) Validate(i uint) *miniCombinedCounter {
	return p.validate(i+1, (*miniCombinedCounter)(nil))
}

func (p *miniCombinedCounterVec) ValidateInit(i uint, zero miniCombinedCounter) *miniCombinedCounter {
	return p.validate(i+1, &zero)
}

func (p *miniCombinedCounterVec) ValidateLen(l uint) (v *miniCombinedCounter) {
	if l > 0 {
		v = p.validate(l, (*miniCombinedCounter)(nil))
	}
	return
}

func (p *miniCombinedCounterVec) ValidateLenInit(l uint, zero miniCombinedCounter) (v *miniCombinedCounter) {
	if l > 0 {
		v = p.validate(l, &zero)
	}
	return
}

func (p miniCombinedCounterVec) Len() uint { return uint(len(p)) }
