package cpf

type Cpf struct {
	Number string
}

func NewCpf(number string) Cpf {
	return Cpf{Number: number}
}

func (cpf Cpf) IsValid() bool {
	return Validate(cpf.Number)
}
