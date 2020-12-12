package abstractfactory

import (
	"testing"
)

func TestMotorbikeFactory(t *testing.T) {
	motorbikeF, err := BuildFactory(MotorbikeFactoryType)
	if err != nil {
		t.Fatal(err)
	}
	motorbikeVehicle, err := motorbikeF.Build(SportMotorbikeType)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Motorbike vehicle has %d wheels\n", motorbikeVehicle.NumWheels())
	// Example for 'Type Assertion' or type-casting
	sportBike, ok := motorbikeVehicle.(Motorbike)
	if !ok {
		t.Fatal("Struct assertion has failed")
	}
	t.Logf("Sport motorbike has type %d\n", sportBike.GetMotorbikeType())

	motorbikeVehicle, err = motorbikeF.Build(CruiseMotorbikeType)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Motorbike vehicle has %d wheels\n", motorbikeVehicle.NumWheels())
	// Example for 'Type Assertion' or type-casting
	cruiseBike, ok := motorbikeVehicle.(Motorbike)
	if !ok {
		t.Fatal("Struct assertion has failed")
	}
	t.Logf("Cruise motorbike has type %d\n", cruiseBike.GetMotorbikeType())

}

func TestCarFactory(t *testing.T) {
	carF, err := BuildFactory(CarFactoryType)
	if err != nil {
		t.Fatal(err)
	}
	carVehicle, err := carF.Build(LuxuryCarType)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Car vehicle has %d seats\n", carVehicle.NumWheels())
	luxuryCar, ok := carVehicle.(Car)
	if !ok {
		t.Fatal("Struct assertion has failed")
	}
	t.Logf("Luxury car has %d doors.\n", luxuryCar.NumDoors())

	carVehicle, err = carF.Build(FamilyCarType)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Car vehicle has %d seats\n", carVehicle.NumWheels())
	familyCar, ok := carVehicle.(Car)
	if !ok {
		t.Fatal("Struct assertion has failed")
	}
	t.Logf("Family car has %d doors.\n", familyCar.NumDoors())
}
