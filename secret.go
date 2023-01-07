package secret

type Secret struct {
	v string
}

func New(v string) Secret {
	return Secret{v: v}
}

func (s Secret) GoString() string {
	return s.String()
}

func (s Secret) String() string {
	if len(s.v) == 0 {
		return "*secret is empty*"
	}
	return "******"
}

func (s Secret) Value() string {
	return s.v
}
