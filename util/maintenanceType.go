package util

// Valid Maintenance Types that are allowed to be logged so far, can add more later
const (
	OilChange     = "oil change"
	TireRotations = "tire rotations"
	Checkup       = "checkup"
)

func IsSupportedMaintenanceType(maintenanceType string) bool {
	switch maintenanceType {
	case OilChange, TireRotations, Checkup:
		return true
	}
	return false
}
