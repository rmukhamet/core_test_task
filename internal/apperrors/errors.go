package apperrors

import "fmt"

var (
	ErrorWrongVersion       = fmt.Errorf("error wrong version")
	ErrorRetailerNotValid   = fmt.Errorf("error retailer not valid")
	ErrorUnknownDataToQueue = fmt.Errorf("error unknown data to queue")
	ErrorRetailerIDRequired = fmt.Errorf("error retailer id field required")
)
