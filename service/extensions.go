package service

import "strings"

func isOwner(reqSub string, reqObj string, polObj string) bool {
	// First find out if there is a UID
	if strings.Contains(polObj, "{uid}") {
		comparable := strings.Replace(polObj, "{uid}", reqSub, 1)

		if comparable == reqObj {
			return true
		}
	}

	return false
}

func isOwnerWrapper(args ...interface{}) (interface{}, error) {
	reqSub := args[0].(string)
	reqObj := args[1].(string)
	polObj := args[2].(string)

	return bool(isOwner(reqSub, reqObj, polObj)), nil
}
