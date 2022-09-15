package business

import (
	"gitlab.artin.ai/backend/courier-management/common/stringutil"
)

func requestHumanReadableId() string {
	return "request_" + stringutil.RandHumanReadableString(10)
}
