package x

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/go-chi/chi/v5"
)

var (
	ErrHTTPBadRequest = errors.New("")
)

const (
	ContentTypeApplicationJSON    = "application/json"
	ContentTypeXWWWFormUrlEncoded = "application/x-www-form-urlencoded"
)

func WriteHTTPResponseJSON(w http.ResponseWriter, code int, obj any) error {
	jsonString, err := json.Marshal(obj)
	if err != nil {
		return err
	}

	w.Header().Add("Content-Type", ContentTypeApplicationJSON)
	w.WriteHeader(code)
	_, err = w.Write(jsonString)
	return err
}

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

func parseURLQuery[T any](obj T, req *http.Request) error {
	return parse(obj, req, false, "query", func(r *http.Request, s string) any {
		return req.URL.Query()[s][0]
	})
}

func parseURLParameter[T any](obj T, req *http.Request) error {
	return parse(obj, req, false, "param", func(r *http.Request, s string) any {
		return chi.URLParam(r, s)
	})
}

func parseJSONBody[T any](obj T, req *http.Request) error {
	m := map[string]any{}
	if err := json.NewDecoder(req.Body).Decode(&m); err != nil {
		return fmt.Errorf("%winvalid json", ErrHTTPBadRequest)
	}

	return parse(obj, req, true, "json", func(r *http.Request, s string) any {
		return m[s]
	})
}

func parseURLEncodedFormData[T any](obj T, req *http.Request) error {
	return parse(obj, req, false, "form", func(req *http.Request, fieldName string) any {
		return req.FormValue(fieldName)
	})
}

func parse[T any](obj T, req *http.Request, strict bool, tagName string, fieldVal func(*http.Request, string) any) error {
	objType := reflect.TypeOf(obj).Elem()
	objVal := reflect.ValueOf(obj).Elem()

	for i := 0; i < objType.NumField(); i++ {
		field := objType.Field(i)
		fieldName := field.Name

		tagValue := field.Tag.Get(tagName)
		tagValue, _, _ = strings.Cut(tagValue, ",")
		if tagValue == "" {
			continue
		}

		fieldValue := fieldVal(req, tagValue)
		if fieldValue == "" {
			continue
		}

		fieldVal := objVal.FieldByName(fieldName)

		switch field.Type.Kind() {
		case reflect.String:
			s, err := any2string(fieldValue, strict)
			if err != nil {
				return fmt.Errorf("%s: %w", tagValue, err)
			}

			fieldVal.SetString(s)

		case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
			intFormVal, err := any2int(fieldValue, strict)
			if err != nil {
				return fmt.Errorf("%s: %w", tagValue, err)
			}
			fieldVal.SetInt(intFormVal)

		case reflect.Float32, reflect.Float64:
			floatFormVal, err := any2float(fieldValue, strict)
			if err != nil {
				return fmt.Errorf("%s: %w", tagValue, err)
			}
			fieldVal.SetFloat(floatFormVal)

		case reflect.Bool:
			boolFormVal, err := any2bool(fieldValue, strict)
			if err != nil {
				return fmt.Errorf("%s: %w", tagValue, err)
			}
			fieldVal.SetBool(boolFormVal)

		default:
			return fmt.Errorf("not support type %s", field.Type.Kind())
		}
	}

	return nil
}
