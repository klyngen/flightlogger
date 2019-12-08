package presentation

import (
	"strings"
)

func (f *FlightLogApi) 

func (f *FlightLogApi) onlyEditOwn(userid string, requestObject string, resourceObject string, action string) bool {

	// IF ACTION == GET -> RETURN TRUE
	if strings.ToUpper(action) == "GET" {
		return true
	}

	slug := getSlug(resourceObject, requestObject)

	if len(slug) > 0 { // We have an ID
		if slug == userid {
			return true
		}
		return false
	}
}

func getSlug(resourcePath string, requestPath string) string {
	// SPLIT THE PATH BY / AND FIND THE LOCATION OF :id
	// Then return that segment of the requestpath

	segments := strings.Split(resourcePath, "/")

	for i, seg := range segments {
		if seg == ":id" {
			return strings.Split(requestPath, "/")[i]
		}
	}
	return ""
}

