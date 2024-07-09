package options

type options struct {
	a int
	b string
}

func (o *options) Configure(mods ...OptionModifier) *options {
	for _, mod := range mods {
		mod(o)
	}
	return o
}

type OptionModifier func(opt *options)

func New() *options {
	return &options{}
}

func WithDefaults() OptionModifier {
	return func(opt *options) {
		opt.a = 0
		opt.b = ""
	}
}

func WithA(a int) OptionModifier {
	return func(opt *options) {
		opt.a = a
	}
}

func WithB(b string) OptionModifier {
	return func(opt *options) {
		opt.b = b
	}
}
