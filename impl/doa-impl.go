package impl

type RealDoAComponent struct {
	FieldDoA *RealDoA
}

type RealDoA struct {
}

func (doa *RealDoA) A () string {
	return "Real A"
}
