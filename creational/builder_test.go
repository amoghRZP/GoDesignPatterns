package creational

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCarBuilder(t *testing.T) {
	var (
		car          VehicleProduct
		manufacturer Manufacturer
		carBuilder   CarBuilder
	)

	car = VehicleProduct{}
	carBuilder = CarBuilder{car: car}

	assert.Equal(t, 0, car.seats)
	assert.Equal(t, 0, car.wheels)
	assert.Equal(t, "", car.engine)

	manufacturer = Manufacturer{b: &carBuilder}
	manufacturer.construct()
	builtCar := carBuilder.build()

	assert.Equal(t, 4, builtCar.seats)
	assert.Equal(t, 4, builtCar.wheels)
	assert.Equal(t, "Car Engine", builtCar.engine)
}

func TestBikeBuilder(t *testing.T) {
	var (
		bike         VehicleProduct
		manufacturer Manufacturer
		bikeBuilder  BikeBuilder
	)

	bike = VehicleProduct{}
	bikeBuilder = BikeBuilder{bike: bike}

	assert.Equal(t, 0, bike.seats)
	assert.Equal(t, 0, bike.wheels)
	assert.Equal(t, "", bike.engine)

	manufacturer = Manufacturer{b: &bikeBuilder}
	manufacturer.construct()
	builtBike := bikeBuilder.build()

	assert.Equal(t, 1, builtBike.seats)
	assert.Equal(t, 2, builtBike.wheels)
	assert.Equal(t, "Bike Engine", builtBike.engine)
}
