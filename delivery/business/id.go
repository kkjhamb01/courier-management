package business

import (
	"github.com/kkjhamb01/courier-management/common/stringutil"
)

func requestHumanReadableId() string {
	return "request_" + stringutil.RandHumanReadableString(10)
}
