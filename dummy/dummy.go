package dummy

type DummyDoAComponent struct {
	FieldDoA *DummyDoA
}

type DummyDoA struct {
}

func (doa *DummyDoA) A () string {
	return "Dummy A"
}
