package xhttp

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/xybor/x/xreflect"
)

const (
	ContentTypeApplicationJSON    = "application/json"
	ContentTypeXWWWFormUrlEncoded = "application/x-www-form-urlencoded"
)

func ParseHTTPRequest[T any](req *http.Request) (T, error) {
	var defaultT T
	var t T

	if err := parseURLParameter(&t, req); err != nil {
		return t, err
	}

	switch req.Method {
	case http.MethodGet:
		if err := parseURLQuery(&t, req); err != nil {
			return defaultT, err
		}

		return t, nil

	case http.MethodPost, http.MethodPut, http.MethodDelete:
		contentType := req.Header.Get("content-type")
		switch contentType {
		case ContentTypeApplicationJSON:
			if err := parseJSONBody(&t, req); err != nil {
				return defaultT, err
			}
			return t, nil

		case ContentTypeXWWWFormUrlEncoded:
			if err := parseURLEncodedFormData(&t, req); err != nil {
				return defaultT, err
			}

			return t, nil

		default:
			return defaultT, fmt.Errorf("%w%s", ErrHTTPBadRequest, fmt.Sprintf("not support content type %s", contentType))
		}

	default:
		return defaultT, fmt.Errorf("%w%s", ErrHTTPBadRequest, fmt.Sprintf("not support method %s", req.Method))
	}
}

func parseURLQuery(obj any, req *http.Request) error {
	return parse(obj, req, false, "query", func(r *http.Request, s string) any {
		query := req.URL.Query()[s]
		if len(query) == 0 {
			return ""
		}

		return query[0]
	})
}

func parseURLParameter(obj any, req *http.Request) error {
	return parse(obj, req, false, "param", func(r *http.Request, s string) any {
		return chi.URLParam(r, s)
	})
}

func parseJSONBody(obj any, req *http.Request) error {
	m := map[string]any{}
	if err := json.NewDecoder(req.Body).Decode(&m); err != nil {
		return fmt.Errorf("%winvalid json", ErrHTTPBadRequest)
	}

	return parse(obj, req, true, "json", func(r *http.Request, s string) any {
		return m[s]
	})
}

func parseURLEncodedFormData(obj any, req *http.Request) error {
	return parse(obj, req, false, "form", func(req *http.Request, fieldName string) any {
		return req.FormValue(fieldName)
	})
}

func parse(obj any, req *http.Request, strict bool, tagName string, fieldVal func(*http.Request, string) any) error {
	return xreflect.Parse(obj, strict, tagName, func(s string) any {
		return fieldVal(req, s)
	})
}
