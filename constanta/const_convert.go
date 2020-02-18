package constanta

import "fmt"

const (
	notFoundInt32 = `Not Found : %d`
)

// Int32 is convert from type to int32
func (et Status) Int32() int32 {
	return int32(et)
}

// String is convert to string message
func (et Status) String() string {
	switch et {
	case StatusActive:
		return "Active"
	case StatusNotUsed:
		return "Not Used"
	case StatusNotActive:
		return "Not Active"
	case StatusDeleted:
		return "Deleted"
	case StatusBlock:
		return "Block"
	default:
		return fmt.Sprintf(notFoundInt32, et)
	}
}
