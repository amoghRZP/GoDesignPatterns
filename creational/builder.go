package creational

// Product to be built by different builders
type VehicleProduct struct {
	seats  int
	wheels int
	engine string
}

// Interface to be implemented by builders
type BuildProcess interface {
	setSeats() BuildProcess
	setWheels() BuildProcess
	setEngine() BuildProcess
	build() VehicleProduct
}

// Manufacturer can have different kind of builders
type Manufacturer struct {
	b BuildProcess
}

func (m *Manufacturer) construct() {
	m.b.setEngine().setSeats().setWheels()
}

// Car Builder implementation
type CarBuilder struct {
	car VehicleProduct
}

func (cb *CarBuilder) setSeats() BuildProcess {
	cb.car.seats = 4
	return cb
}

func (cb *CarBuilder) setWheels() BuildProcess {
	cb.car.wheels = 4
	return cb
}

func (cb *CarBuilder) setEngine() BuildProcess {
	cb.car.engine = "Car Engine"
	return cb
}

func (cb *CarBuilder) build() VehicleProduct {
	return cb.car
}

// Bike Builder implementation
type BikeBuilder struct {
	bike VehicleProduct
}

func (bb *BikeBuilder) setSeats() BuildProcess {
	bb.bike.seats = 1
	return bb
}

func (bb *BikeBuilder) setWheels() BuildProcess {
	bb.bike.wheels = 2
	return bb
}

func (bb *BikeBuilder) setEngine() BuildProcess {
	bb.bike.engine = "Bike Engine"
	return bb
}

func (bb *BikeBuilder) build() VehicleProduct {
	return bb.bike
}
