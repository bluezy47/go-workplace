package internal

import "fmt"

// Employee struct
type Employee struct {
	Name       string `json:"name"`
	Age        int    `json:"age"`
	Department string `json:"department"`
	Score      int    `json:"score"`
}

// ScoreCalculationHook interface.
type ScoreCalculationHook interface {
	BeforeCalculation(employee Employee)
	AfterCalculation(employee Employee, score float64) float64
}

// DepartmentScoreHook struct. implements ScoreCalculationHook
type DepartmentScoreHook struct {
	department string
}

func NewDepartmentScoreHook(department string) *DepartmentScoreHook {
	return &DepartmentScoreHook{department: department}
}

func (dsh *DepartmentScoreHook) BeforeCalculation(employee Employee) {
	// Feature-specific logic before calculation
	if employee.Department == dsh.department {
		fmt.Println("Processing employee from department:", dsh.department)
	}
}

func (dsh *DepartmentScoreHook) AfterCalculation(employee Employee, score float64) float64 {
	// Feature-specific logic after calculation
	if employee.Department == dsh.department {
		score += 10.0 // Example adjustment for department
	}
	return score
}

// RoleScoreHook struct. implements ScoreCalculationHook
type RoleScoreHook struct {
	role string
}

func NewRoleScoreHook(role string) *RoleScoreHook {
	return &RoleScoreHook{role: role}
}

func (rsh *RoleScoreHook) BeforeCalculation(employee Employee) {
	// Feature-specific logic before calculation
	if employee.Department == rsh.role {
		fmt.Println("Processing employee with role:", rsh.role)
	}
}

func (rsh *RoleScoreHook) AfterCalculation(employee Employee, score float64) float64 {
	// Feature-specific logic after calculation
	if employee.Department == rsh.role {
		score += 5.0 // Example adjustment for role
	}
	return score
}
