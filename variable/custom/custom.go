package custom

import (
	"github.com/grafana-tools/sdk"
)

type Option func(constant *Custom)

// Values represent a "label" to "value" map of options for a custom variable.
type Values map[string]string

type Custom struct {
	Builder sdk.TemplateVar
}

func New(name string, options ...Option) *Custom {
	constant := &Custom{Builder: sdk.TemplateVar{
		Name:  name,
		Label: name,
		Type:  "custom",
	}}

	for _, opt := range options {
		opt(constant)
	}

	return constant
}

func WithValues(values Values) Option {
	return func(constant *Custom) {
		for label, value := range values {
			constant.Builder.Options = append(constant.Builder.Options, sdk.Option{
				Text:  label,
				Value: value,
			})
		}
	}
}

func WithDefault(value string) Option {
	return func(constant *Custom) {
		constant.Builder.Current = sdk.Current{
			Text: value,
		}
	}
}

func WithLabel(label string) Option {
	return func(constant *Custom) {
		constant.Builder.Label = label
	}
}

func HideLabel() Option {
	return func(constant *Custom) {
		constant.Builder.Hide = 1
	}
}

func Hide() Option {
	return func(constant *Custom) {
		constant.Builder.Hide = 2
	}
}

func Multi() Option {
	return func(constant *Custom) {
		constant.Builder.Multi = true
	}
}

func IncludeAll() Option {
	return func(constant *Custom) {
		constant.Builder.IncludeAll = true
		constant.Builder.Options = append(constant.Builder.Options, sdk.Option{
			Text:  "All",
			Value: "$__all",
		})
	}
}
