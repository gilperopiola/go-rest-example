package common

type LogSource int

const (
	Unknown    LogSource = iota // 0
	NewRelic                    // 1
	Prometheus                  // 2
	Gin                         // 3
	Gorm                        // 4
)

func (ls LogSource) String() string {
	return [...]string{"unknown", "new_relic", "prometheus", "gin", "gorm"}[ls]
}
