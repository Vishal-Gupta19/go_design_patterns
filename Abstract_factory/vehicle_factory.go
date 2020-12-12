package abstractfactory

import (
	"errors"
	"fmt"
)

// Factory ...
type Factory interface {
	Build(v int) (Vehicle, error)
}

const (
	// CarFactoryType ...
	CarFactoryType = 1
	// MotorbikeFactoryType ...
	MotorbikeFactoryType = 2
)

// BuildFactory ...
func BuildFactory(f int) (Factory, error) {
	switch f {
	case CarFactoryType:
		return new(CarFactory), nil
	case MotorbikeFactoryType:
		return new(MotorbikeFactory), nil
	default:
		return nil, errors.New(fmt.Sprintf("Factory with id %d not recognized\n", f))
	}
}
