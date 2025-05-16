package domain

type PackCalculatorService interface {
	CalculatePacks(amount int) (*CalculationResult, error)
}
