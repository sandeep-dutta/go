// autogenerated: do not edit!
// generated from gentemplate [gentemplate -d Package=vnet -id countersVec -d VecType=CountersVec -d Type=Counters github.com/platinasystems/go/elib/vec.tmpl]

package vnet

import (
	"github.com/platinasystems/go/elib"
)

type CountersVec []Counters

func (p *CountersVec) Resize(n uint) {
	c := elib.Index(cap(*p))
	l := elib.Index(len(*p)) + elib.Index(n)
	if l > c {
		c = elib.NextResizeCap(l)
		q := make([]Counters, l, c)
		copy(q, *p)
		*p = q
	}
	*p = (*p)[:l]
}

func (p *CountersVec) validate(new_len uint, zero *Counters) *Counters {
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

func (p *CountersVec) validateSlowPath(zero *Counters,
	c, l, lʹ elib.Index) *Counters {
	if l > c {
		cNext := elib.NextResizeCap(l)
		q := make([]Counters, cNext, cNext)
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

func (p *CountersVec) Validate(i uint) *Counters {
	return p.validate(i+1, (*Counters)(nil))
}

func (p *CountersVec) ValidateInit(i uint, zero Counters) *Counters {
	return p.validate(i+1, &zero)
}

func (p *CountersVec) ValidateLen(l uint) (v *Counters) {
	if l > 0 {
		v = p.validate(l, (*Counters)(nil))
	}
	return
}

func (p *CountersVec) ValidateLenInit(l uint, zero Counters) (v *Counters) {
	if l > 0 {
		v = p.validate(l, &zero)
	}
	return
}

func (p CountersVec) Len() uint { return uint(len(p)) }
