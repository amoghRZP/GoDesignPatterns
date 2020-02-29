package creational

// Interface for increment method
type SingletonInterface interface {
	Increment() int64
}

// struct for singleton
type Singleton struct {
	count int64
}

var instance *Singleton

// Function to get singleton instance (equivalent to static method in java)
func GetInstance() *Singleton {
	if instance == nil {
		instance = new(Singleton)
	}
	return instance
}

func (s *Singleton) Increment() int64 {
	s.count++
	return s.count
}
