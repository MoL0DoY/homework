package calculators

type Calculator interface {
	Execute(operator string) (string, error)
}
