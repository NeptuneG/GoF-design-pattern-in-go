package singleton

// Singleton is the tyoe that to be exposed
type Singleton struct {
}

var instance *Singleton

// GetInstance return the single instance of type Singleton
func GetInstance() *Singleton {
	return instance
}
