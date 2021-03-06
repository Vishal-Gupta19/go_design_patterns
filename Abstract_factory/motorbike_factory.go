package abstractfactory

import (
	"errors"
	"fmt"
)

const (
	// SportMotorbikeType ...
	SportMotorbikeType = 1
	// CruiseMotorbikeType ...
	CruiseMotorbikeType = 2
)

// MotorbikeFactory ...
type MotorbikeFactory struct{}

// Build ..
func (m *MotorbikeFactory) Build(v int) (Vehicle, error) {
	switch v {
	case SportMotorbikeType:
		return new(SportMotorbike), nil
	case CruiseMotorbikeType:
		return new(CruiseMotorbike), nil
	default:
		return nil, errors.New(fmt.Sprintf("Vehicle of type %d not recognized\n", v))

	}
}
