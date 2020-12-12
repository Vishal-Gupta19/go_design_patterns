package abstractfactory

import (
	"errors"
	"fmt"
)

const (
	// LuxuryCarType ...
	LuxuryCarType = 1
	// FamilyCarType ...
	FamilyCarType = 2
)

// CarFactory ...
type CarFactory struct{}

// Build ...
func (c *CarFactory) Build(v int) (Vehicle, error) {
	switch v {
	case LuxuryCarType:
		return new(LuxuryCar), nil
	case FamilyCarType:
		return new(FamilyCar), nil
	default:
		return nil, errors.New(fmt.Sprintf("vehicle of type %d not recognized", v))
	}
}
