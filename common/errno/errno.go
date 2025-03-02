package errno

const (
	//Predictable errors
	ErrUnknown         = "00000"
	ErrProductNotFound = "10001"
	ErrStockNotEnough  = "10002"
	ErrUserNotFound    = "10003"
	//Consistent errors
	ErrUnknownConsistent  = "20000"
	ErrDatabaseConsistent = "20001"

	//System errors
	ErrUnknownSystem  = "30000"
	ErrRedisSystem    = "30001"
	ErrDatabaseSystem = "30002"
)
