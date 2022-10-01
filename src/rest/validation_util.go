package rest

import (
	v "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

func rules(fieldPtr any, rules []v.Rule, optionRules ...v.Rule) *v.FieldRules {
	for _, r := range optionRules {
		rules = append(rules, r)
	}
	return v.Field(fieldPtr, rules...)
}

func requiredRule(fieldPtr any, optionRules ...v.Rule) *v.FieldRules {
	return rules(fieldPtr, []v.Rule{v.Required}, optionRules...)
}

func ymdRule(fieldPtr any, min, max int, optionRules ...v.Rule) *v.FieldRules {
	return rules(fieldPtr, []v.Rule{v.Min(min), v.Max(max)}, optionRules...)
}

func yearRule(fieldPtr any, optionRules ...v.Rule) *v.FieldRules {
	return ymdRule(fieldPtr, 1900, 3000, optionRules...)
}

func monthRule(fieldPtr any, optionRules ...v.Rule) *v.FieldRules {
	return ymdRule(fieldPtr, 1, 12, optionRules...)
}

func dayRule(fieldPtr any, optionRules ...v.Rule) *v.FieldRules {
	return ymdRule(fieldPtr, 1, 31, optionRules...)
}

func stringRule(fieldPtr any, min, max int, optionRules ...v.Rule) *v.FieldRules {
	return rules(fieldPtr, []v.Rule{v.RuneLength(min, max)}, optionRules...)
}

func urlRule(fieldPtr any, optionRules ...v.Rule) *v.FieldRules {
	optionRules = append(optionRules, is.URL)
	return stringRule(fieldPtr, 1, 4096, optionRules...)
}
