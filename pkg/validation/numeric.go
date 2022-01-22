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

func (b IntBoundaryRule) Merge(o IntBoundaryRule) IntBoundaryRule {
	new := IntBoundaryRule{}
	new.SetMin(b.Min, b.Flags.Has(NumericBoundaryExclusiveMin))
	new.SetMax(b.Max, b.Flags.Has(NumericBoundaryExclusiveMax))
	new.SetMin(b.Min, o.Flags.Has(NumericBoundaryExclusiveMin))
	new.SetMax(b.Max, o.Flags.Has(NumericBoundaryExclusiveMax))
	return new
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

func (b UintBoundaryRule) Merge(o UintBoundaryRule) UintBoundaryRule {
	new := UintBoundaryRule{}
	new.SetMin(b.Min, b.Flags.Has(NumericBoundaryExclusiveMin))
	new.SetMax(b.Max, b.Flags.Has(NumericBoundaryExclusiveMax))
	new.SetMin(b.Min, o.Flags.Has(NumericBoundaryExclusiveMin))
	new.SetMax(b.Max, o.Flags.Has(NumericBoundaryExclusiveMax))
	return new
}

type FloatBoundaryRule struct {
	Flags utils.Flags
	Max   float64
	Min   float64
}

func (b FloatBoundaryRule) SetMin(value float64, exclusive bool) {
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

func (b FloatBoundaryRule) SetMax(value float64, exclusive bool) {
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

func (b FloatBoundaryRule) CreateValidationFunction() ValidationFunction {
	return func(value interface{}) error {
		v := utils.ForceFloat64(value)
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

func (b FloatBoundaryRule) Merge(o FloatBoundaryRule) FloatBoundaryRule {
	new := FloatBoundaryRule{}
	new.SetMin(b.Min, b.Flags.Has(NumericBoundaryExclusiveMin))
	new.SetMax(b.Max, b.Flags.Has(NumericBoundaryExclusiveMax))
	new.SetMin(b.Min, o.Flags.Has(NumericBoundaryExclusiveMin))
	new.SetMax(b.Max, o.Flags.Has(NumericBoundaryExclusiveMax))
	return new
}

func createIntegerBoundaryRule(kind reflect.Kind) IntBoundaryRule {
	rule := IntBoundaryRule{}
	switch kind {
	case reflect.Int:
		rule.SetMin(math.MinInt, false)
		rule.SetMax(math.MaxInt, false)
	case reflect.Int8:
		rule.SetMin(math.MinInt8, false)
		rule.SetMax(math.MaxInt8, false)
	case reflect.Int16:
		rule.SetMin(math.MinInt16, false)
		rule.SetMax(math.MaxInt16, false)
	case reflect.Int32:
		rule.SetMin(math.MinInt32, false)
		rule.SetMax(math.MaxInt32, false)
	}
	return rule
}

func createUnsignedIntegerBoundaryRule(kind reflect.Kind) UintBoundaryRule {
	rule := UintBoundaryRule{}
	switch kind {
	case reflect.Uint:
		rule.SetMin(0, false)
		rule.SetMax(math.MaxUint, false)
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
	return rule
}

func createIntegerBoundaryRuleFromTags(tag reflect.StructTag) IntBoundaryRule {
	rule := IntBoundaryRule{}
	if len(tag.Get("minimum")) > 0 {
		rule.SetMin(utils.ForceConvertInt(tag.Get("minimum")), false)
	} else if len(tag.Get("exclusiveMinimum")) > 0 {
		rule.SetMin(utils.ForceConvertInt(tag.Get("exclusiveMinimum")), true)
	}
	if len(tag.Get("maximum")) > 0 {
		rule.SetMax(utils.ForceConvertInt(tag.Get("maximum")), false)
	} else if len(tag.Get("exclusiveMaximum")) > 0 {
		rule.SetMax(utils.ForceConvertInt(tag.Get("exclusiveMaximum")), true)
	}
	return rule
}

func createUnsignedIntegerBoundaryRuleFromTags(tag reflect.StructTag) UintBoundaryRule {
	rule := UintBoundaryRule{}
	if len(tag.Get("minimum")) > 0 {
		rule.SetMin(utils.ForceConvertUint(tag.Get("minimum")), false)
	} else if len(tag.Get("exclusiveMinimum")) > 0 {
		rule.SetMin(utils.ForceConvertUint(tag.Get("exclusiveMinimum")), true)
	}
	if len(tag.Get("maximum")) > 0 {
		rule.SetMax(utils.ForceConvertUint(tag.Get("maximum")), false)
	} else if len(tag.Get("exclusiveMaximum")) > 0 {
		rule.SetMax(utils.ForceConvertUint(tag.Get("exclusiveMaximum")), true)
	}
	return rule
}

func createFloatBoundaryRuleFromTags(tag reflect.StructTag) FloatBoundaryRule {
	rule := FloatBoundaryRule{}
	if len(tag.Get("minimum")) > 0 {
		rule.SetMin(utils.ForceConvertFloat(tag.Get("minimum")), false)
	} else if len(tag.Get("exclusiveMinimum")) > 0 {
		rule.SetMin(utils.ForceConvertFloat(tag.Get("exclusiveMinimum")), true)
	}
	if len(tag.Get("maximum")) > 0 {
		rule.SetMax(utils.ForceConvertFloat(tag.Get("maximum")), false)
	} else if len(tag.Get("exclusiveMaximum")) > 0 {
		rule.SetMax(utils.ForceConvertFloat(tag.Get("exclusiveMaximum")), true)
	}
	return rule
}
