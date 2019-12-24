package component

type DoA interface {
	A() string
}

type DoAComponent struct {
	FieldDoA *DoA
}
