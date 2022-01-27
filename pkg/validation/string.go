package validation

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/yndc/recon/pkg/utils"
)

const (
	StringLengthBoundaryMin = 1 << iota
	StringLengthBoundaryMax
)

type StringLengthRule struct {
	Flags utils.Flags
	Max   int
	Min   int
}

func (b *StringLengthRule) SetMin(value int) {
	if value > b.Min || !b.Flags.Has(StringLengthBoundaryMin) {
		b.Min = value
	}
	b.Flags |= StringLengthBoundaryMin
}

func (b *StringLengthRule) SetMax(value int) {
	if value < b.Max || !b.Flags.Has(StringLengthBoundaryMax) {
		b.Max = value
	}
	b.Flags |= StringLengthBoundaryMax
}

func (b *StringLengthRule) CreateValidationFunction() ValidationFunction {
	return func(value interface{}) error {
		v := value.(string)
		if b.Flags.Has(NumericBoundaryMin) {
			if len(v) < b.Min {
				return fmt.Errorf("string length is less than the minimum length (%d)", b.Min)
			}
		}
		if b.Flags.Has(NumericBoundaryMax) {
			if len(v) > b.Max {
				return fmt.Errorf("string length is greater than the maximum length (%d)", b.Max)
			}
		}
		return nil
	}
}

func (b *StringLengthRule) Merge(o StringLengthRule) StringLengthRule {
	new := StringLengthRule{}
	if b.Flags.Has(NumericBoundaryMin) {
		new.SetMin(b.Min)
	}
	if b.Flags.Has(NumericBoundaryMax) {
		new.SetMax(b.Max)
	}
	if o.Flags.Has(NumericBoundaryMin) {
		new.SetMin(o.Min)
	}
	if o.Flags.Has(NumericBoundaryMax) {
		new.SetMax(o.Max)
	}
	return new
}

func createStringLengthRuleFromTags(tag reflect.StructTag) StringLengthRule {
	rule := StringLengthRule{}
	if len(tag.Get("minLength")) > 0 {
		min, err := strconv.Atoi(tag.Get("minLength"))
		if err != nil {
			panic(err)
		}
		rule.SetMin(min)
	}
	if len(tag.Get("maxLength")) > 0 {
		max, err := strconv.Atoi(tag.Get("maxLength"))
		if err != nil {
			panic(err)
		}
		rule.SetMax(max)
	}
	return rule
}
