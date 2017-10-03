// autogenerated: do not edit!
// generated from gentemplate [gentemplate -d Package=ip -id nextHopVec -d VecType=nextHopVec -d Type=nextHop github.com/platinasystems/go/elib/vec.tmpl]

// Copyright 2016 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ip

import (
	"github.com/platinasystems/go/elib"
)

type nextHopVec []nextHop

func (p *nextHopVec) Resize(n uint) {
	old_cap := uint(cap(*p))
	new_len := uint(len(*p)) + n
	if new_len > old_cap {
		new_cap := elib.NextResizeCap(new_len)
		q := make([]nextHop, new_len, new_cap)
		copy(q, *p)
		*p = q
	}
	*p = (*p)[:new_len]
}

func (p *nextHopVec) validate(new_len uint, zero nextHop) *nextHop {
	old_cap := uint(cap(*p))
	old_len := uint(len(*p))
	if new_len <= old_cap {
		// Need to reslice to larger length?
		if new_len > old_len {
			*p = (*p)[:new_len]
			for i := old_len; i < new_len; i++ {
				(*p)[i] = zero
			}
		}
		return &(*p)[new_len-1]
	}
	return p.validateSlowPath(zero, old_cap, new_len, old_len)
}

func (p *nextHopVec) validateSlowPath(zero nextHop, old_cap, new_len, old_len uint) *nextHop {
	if new_len > old_cap {
		new_cap := elib.NextResizeCap(new_len)
		q := make([]nextHop, new_cap, new_cap)
		copy(q, *p)
		for i := old_len; i < new_cap; i++ {
			q[i] = zero
		}
		*p = q[:new_len]
	}
	if new_len > old_len {
		*p = (*p)[:new_len]
	}
	return &(*p)[new_len-1]
}

func (p *nextHopVec) Validate(i uint) *nextHop {
	var zero nextHop
	return p.validate(i+1, zero)
}

func (p *nextHopVec) ValidateInit(i uint, zero nextHop) *nextHop {
	return p.validate(i+1, zero)
}

func (p *nextHopVec) ValidateLen(l uint) (v *nextHop) {
	if l > 0 {
		var zero nextHop
		v = p.validate(l, zero)
	}
	return
}

func (p *nextHopVec) ValidateLenInit(l uint, zero nextHop) (v *nextHop) {
	if l > 0 {
		v = p.validate(l, zero)
	}
	return
}

func (p *nextHopVec) ResetLen() {
	if *p != nil {
		*p = (*p)[:0]
	}
}

func (p nextHopVec) Len() uint { return uint(len(p)) }
