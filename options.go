package wasmhttp

// Option represents a configurable option for the Serve function.
type Option func(o *options)

type options struct {
	stripPrefix *bool
}

// WithStripPrefix allows configuring the handler to strip the prefix from the
// request path.
func WithStripPrefix(stripPrefix bool) Option {
	return func(o *options) {
		o.stripPrefix = &stripPrefix
	}
}

// setDefaults sets the default values for the options struct.
// If the stripPrefix field is nil, it sets it to true.
func (o *options) setDefaults() {
	if o.stripPrefix == nil {
		var stripPrefix = true
		o.stripPrefix = &stripPrefix
	}
}
