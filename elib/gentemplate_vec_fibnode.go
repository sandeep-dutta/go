// autogenerated: do not edit!
// generated from gentemplate [gentemplate -d Package=elib -id fibNode -d VecType=fibNodeVec -d Type=fibNode vec.tmpl]

package elib

type fibNodeVec []fibNode

func (p *fibNodeVec) Resize(n uint) {
	c := Index(cap(*p))
	l := Index(len(*p)) + Index(n)
	if l > c {
		c = NextResizeCap(l)
		q := make([]fibNode, l, c)
		copy(q, *p)
		*p = q
	}
	*p = (*p)[:l]
}

func (p *fibNodeVec) validate(new_len uint, zero *fibNode) *fibNode {
	c := Index(cap(*p))
	lʹ := Index(len(*p))
	l := Index(new_len)
	if l <= c {
		// Need to reslice to larger length?
		if l >= lʹ {
			*p = (*p)[:l]
		}
		return &(*p)[l-1]
	}
	return p.validateSlowPath(zero, c, l, lʹ)
}

func (p *fibNodeVec) validateSlowPath(zero *fibNode,
	c, l, lʹ Index) *fibNode {
	if l > c {
		cNext := NextResizeCap(l)
		q := make([]fibNode, cNext, cNext)
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

func (p *fibNodeVec) Validate(i uint) *fibNode {
	return p.validate(i+1, (*fibNode)(nil))
}

func (p *fibNodeVec) ValidateInit(i uint, zero fibNode) *fibNode {
	return p.validate(i+1, &zero)
}

func (p *fibNodeVec) ValidateLen(l uint) (v *fibNode) {
	if l > 0 {
		v = p.validate(l, (*fibNode)(nil))
	}
	return
}

func (p *fibNodeVec) ValidateLenInit(l uint, zero fibNode) (v *fibNode) {
	if l > 0 {
		v = p.validate(l, &zero)
	}
	return
}

func (p fibNodeVec) Len() uint { return uint(len(p)) }
