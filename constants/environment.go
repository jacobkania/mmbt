package constants

import (
	"fmt"
)

type environment struct {
	Development Constant
	Production  Constant
}

// Environment is the type of environment the code will run in
var Environment = &environment{
	Development: "development",
	Production:  "production",
}

func (e *environment) Validate(environmentName string) (Constant, error) {
	if environmentName == "development" || environmentName == "production" {
		return Constant(environmentName), nil
	}
	return "", fmt.Errorf("Invalid environment")
}
