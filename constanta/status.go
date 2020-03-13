package constanta

// Status is list of status in table
type Status int32

// List of status
const (
	StatusActive    Status = 1
	StatusNotUsed   Status = 0
	StatusNotActive Status = -1
	StatusDeleted   Status = -2
	StatusBlock     Status = -3
)
