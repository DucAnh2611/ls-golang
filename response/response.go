package response

import (
	"github.com/DucAnh2611/ls-golang/constants"
)

func Success[K comparable](data K) map[string]any {
	res := map[string]any{
		"success": true,
		"code":    constants.SuccessOk,
	}

	var nilValue K
	if data != nilValue {
		res["data"] = data
	}

	return res
}
func Error(code constants.ResponseCode) map[string]any {
	res := map[string]any{
		"success": false,
		"code":    code,
	}

	return res
}
