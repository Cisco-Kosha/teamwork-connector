package app

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strconv"
)

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func getPageRange(params url.Values, respHeaders http.Header, pageCount int) (int, int, error) {
	var err error
	pageStart := 1
	pageEnd := 1
	numPages := 1

	if val, ok := params["pageStart"]; ok {
		pageStart, err = strconv.Atoi(val[0])
		pageEnd = pageStart
		if err != nil {
			return 0, 0, err
		}
	}
	if val, ok := params["pageEnd"]; ok {
		pageEnd, err = strconv.Atoi(val[0])
		if err != nil {
			return 0, 0, err
		}
	}

	if pageCount != 0 {
		numPages = pageCount
	} else {
		if respHeaders.Get("X-Pages") != "" {
			numPages, err = strconv.Atoi(respHeaders.Get("X-Pages"))
			if err != nil {
				return 0, 0, err
			}
			if numPages == 0 {
				numPages = 1
			}
		}
	}

	if pageStart > numPages || pageEnd > numPages || pageStart < 1 {
		return 0, 0, errors.New("invalid pageStart or pageEnd value")
	}

	if val, ok := params["allPages"]; ok && val[0] == "true" {
		pageStart = 1
		pageEnd = numPages
	}
	return pageStart, pageEnd, nil
}
