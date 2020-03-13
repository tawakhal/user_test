package constanta

// ErrCode is list of Error Code
type ErrCode int32

// List constanta for ErrCode
const (
	ErrCodeDatabase  ErrCode = 100
	ErrCodeDepedency ErrCode = 200
	ErrCodeRedis     ErrCode = 300
	ErrCodeLibrary   ErrCode = 400
)

// RespCode is list of response code
type RespCode int32

// List conttanta for RespCode
const (
	RespCodeSuccess       RespCode = 1000
	RespCodeSystemError   RespCode = 2000
	RespCodeBusinessError RespCode = 3000
)
