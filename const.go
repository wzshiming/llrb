package llrb

import (
	"golang.org/x/exp/constraints"
)

type Ordered interface {
	constraints.Ordered
}

type Any any
