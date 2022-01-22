package validation

import (
	"fmt"
	"math"
	"reflect"

	"github.com/yndc/recon/pkg/utils"
)

const (
	NumericBoundaryMin = 1 << iota
	NumericBoundaryMax
	NumericBoundaryExclusiveMin
	NumericBoundaryExclusiveMax
)

type IntBoundaryRule struct {
	Flags utils.Flags
	Max   int64
	Min   int64
}

func (b IntBoundaryRule) SetMin(value int64, exclusive bool) {
	b.Flags |= NumericBoundaryMin
	if value < b.Min {
		b.Min = value
		if exclusive {
			b.Flags |= NumericBoundaryExclusiveMin
		}
	} else if value == b.Min && exclusive {
		b.Flags |= NumericBoundaryExclusiveMin
	}
}

func (b IntBoundaryRule) SetMax(value int64, exclusive bool) {
	b.Flags |= NumericBoundaryMax
	if value > b.Max {
		b.Max = value
		if exclusive {
			b.Flags |= NumericBoundaryExclusiveMax
		}
	} else if value == b.Max && exclusive {
		b.Flags |= NumericBoundaryExclusiveMax
	}
}

func (b IntBoundaryRule) CreateValidationFunction() ValidationFunction {
	return func(value interface{}) error {
		v := utils.ForceInt64(value)
		if b.Flags.Has(NumericBoundaryMin) {
			if b.Flags.Has(NumericBoundaryExclusiveMin) {
				if v <= b.Min {
					return fmt.Errorf("value is less than or equal to the minimum value (%d)", b.Min)
				}
			} else {
				if v < b.Min {
					return fmt.Errorf("value is less than the minimum value (%d)", b.Min)
				}
			}
		}
		if b.Flags.Has(NumericBoundaryMax) {
			if b.Flags.Has(NumericBoundaryExclusiveMax) {
				if v >= b.Max {
					return fmt.Errorf("value is greater than or equal to the maximum value (%d)", b.Max)
				}
			} else {
				if v > b.Max {
					return fmt.Errorf("value is greater than the maximum value (%d)", b.Max)
				}
			}
		}
		return nil
	}
}

type UintBoundaryRule struct {
	Flags utils.Flags
	Max   uint64
	Min   uint64
}

func (b UintBoundaryRule) SetMin(value uint64, exclusive bool) {
	b.Flags |= NumericBoundaryMin
	if value < b.Min {
		b.Min = value
		if exclusive {
			b.Flags |= NumericBoundaryExclusiveMin
		}
	} else if value == b.Min && exclusive {
		b.Flags |= NumericBoundaryExclusiveMin
	}
}

func (b UintBoundaryRule) SetMax(value uint64, exclusive bool) {
	b.Flags |= NumericBoundaryMax
	if value > b.Max {
		b.Max = value
		if exclusive {
			b.Flags |= NumericBoundaryExclusiveMax
		}
	} else if value == b.Max && exclusive {
		b.Flags |= NumericBoundaryExclusiveMax
	}
}

func (b UintBoundaryRule) CreateValidationFunction() ValidationFunction {
	return func(value interface{}) error {
		v := utils.ForceUint64(value)
		if b.Flags.Has(NumericBoundaryMin) {
			if b.Flags.Has(NumericBoundaryExclusiveMin) {
				if v <= b.Min {
					return fmt.Errorf("value is less than or equal to the minimum value (%d)", b.Min)
				}
			} else {
				if v < b.Min {
					return fmt.Errorf("value is less than the minimum value (%d)", b.Min)
				}
			}
		}
		if b.Flags.Has(NumericBoundaryMax) {
			if b.Flags.Has(NumericBoundaryExclusiveMax) {
				if v >= b.Max {
					return fmt.Errorf("value is greater than or equal to the maximum value (%d)", b.Max)
				}
			} else {
				if v > b.Max {
					return fmt.Errorf("value is greater than the maximum value (%d)", b.Max)
				}
			}
		}
		return nil
	}
}

func createIntegerBoundaryValidator(kind reflect.Kind) ValidationFunction {
	rule := IntBoundaryRule{}
	switch kind {
	case reflect.Int8:
		rule.SetMin(math.MinInt8, false)
		rule.SetMax(math.MaxInt8, false)
	case reflect.Int16:
		rule.SetMin(math.MinInt16, false)
		rule.SetMax(math.MaxInt16, false)
	case reflect.Int32:
		rule.SetMin(math.MinInt32, false)
		rule.SetMax(math.MaxInt32, false)
	case reflect.Uint8:
		rule.SetMin(0, false)
		rule.SetMax(math.MaxUint8, false)
	case reflect.Uint16:
		rule.SetMin(0, false)
		rule.SetMax(math.MaxUint16, false)
	case reflect.Uint32:
		rule.SetMin(0, false)
		rule.SetMax(math.MaxUint32, false)
	}
	return rule.CreateValidationFunction()
}
