package errors

type CustomErrorI interface {
	Error() string
}

type CustomError string

func (e *CustomError) Error() string {
	return "Salario abaixo do minimo"
}
