package foobar

type Foo struct {
	Bar string
	Baz int
}

func (*Foo) Foo(s string) error {
	return nil
}

func (*Foo) FooBaz(s string, i int, b []byte) error {
	return nil
}

type Bar struct {
	Foo string
}

func (*Bar) bar(s string) error {
	return nil
}

func Qux(s string) error {
	return nil
}

func Quux(s string) (string, error) {
	return "", nil
}

func Corge(s string, i int) (int, string, error) {
	return 123, "", nil
}
