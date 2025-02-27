package internal

type EmployeeScoreCalculator struct {
	hooks []ScoreCalculationHook
}

func (esc *EmployeeScoreCalculator) AddHook(hook ScoreCalculationHook) {
	esc.hooks = append(esc.hooks, hook)
}

func (esc *EmployeeScoreCalculator) CalculateScore(employee Employee) float64 {
	// Call "BeforeCalculation" hooks
	for _, hook := range esc.hooks {
		hook.BeforeCalculation(employee)
	}

	// Core logic for calculating employee score
	score := 0.0
	// Add your conditions and loops here

	// Call "AfterCalculation" hooks
	for _, hook := range esc.hooks {
		score = hook.AfterCalculation(employee, score)
	}

	return score
}
