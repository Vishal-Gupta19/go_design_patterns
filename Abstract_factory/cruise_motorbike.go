package abstractfactory

// CruiseMotorbike ...
type CruiseMotorbike struct{}

// NumWheels ...
func (s *CruiseMotorbike) NumWheels() int {
	return 2
}

// NumSeats ...
func (s *CruiseMotorbike) NumSeats() int {
	return 1
}

// GetMotorbikeType ...
func (s *CruiseMotorbike) GetMotorbikeType() int {
	return CruiseMotorbikeType
}
