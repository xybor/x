package x

import (
	"fmt"
	"strconv"
)

func any2int(a any, strict bool) (int64, error) {
	switch t := a.(type) {
	case int:
		return int64(t), nil
	case int8:
		return int64(t), nil
	case int16:
		return int64(t), nil
	case int32:
		return int64(t), nil
	case int64:
		return int64(t), nil
	case float32:
		if t == float32(int64(t)) {
			return int64(t), nil
		}

		return 0, fmt.Errorf("%wexpected int, got float", ErrHTTPBadRequest)
	case float64:
		if t == float64(int64(t)) {
			return int64(t), nil
		}

		return 0, fmt.Errorf("%wexpected int, got float", ErrHTTPBadRequest)
	case string:
		if !strict {
			if n, err := strconv.ParseInt(t, 10, 64); err == nil {
				return n, nil
			}
		}

		return 0, fmt.Errorf("%wexpected int, got string", ErrHTTPBadRequest)
	case bool:
		return 0, fmt.Errorf("%wexpected int, got bool", ErrHTTPBadRequest)
	default:
		return 0, fmt.Errorf("not handle for %T", t)
	}
}

func any2float(a any, strict bool) (float64, error) {
	switch t := a.(type) {
	case int:
		return float64(t), nil
	case int8:
		return float64(t), nil
	case int16:
		return float64(t), nil
	case int32:
		return float64(t), nil
	case int64:
		return float64(t), nil
	case float32:
		return float64(t), nil
	case float64:
		return float64(t), nil
	case string:
		if !strict {
			if n, err := strconv.ParseFloat(t, 64); err == nil {
				return n, nil
			}
		}

		return 0, fmt.Errorf("%wexpected float, got string", ErrHTTPBadRequest)
	case bool:
		return 0, fmt.Errorf("%wexpected float, got bool", ErrHTTPBadRequest)
	default:
		return 0, fmt.Errorf("not handle for %T", t)
	}
}

func any2bool(a any, strict bool) (bool, error) {
	switch t := a.(type) {
	case int:
		if !strict {
			switch t {
			case 0:
				return false, nil
			case 1:
				return true, nil
			}
		}
		return false, fmt.Errorf("%wexpected bool, got number", ErrHTTPBadRequest)
	case int8:
		if !strict {
			switch t {
			case 0:
				return false, nil
			case 1:
				return true, nil
			}
		}
		return false, fmt.Errorf("%wexpected bool, got number", ErrHTTPBadRequest)
	case int16:
		if !strict {
			switch t {
			case 0:
				return false, nil
			case 1:
				return true, nil
			}
		}
		return false, fmt.Errorf("%wexpected bool, got number", ErrHTTPBadRequest)
	case int32:
		if !strict {
			switch t {
			case 0:
				return false, nil
			case 1:
				return true, nil
			}
		}
		return false, fmt.Errorf("%wexpected bool, got number", ErrHTTPBadRequest)
	case int64:
		if !strict {
			switch t {
			case 0:
				return false, nil
			case 1:
				return true, nil
			}
		}
		return false, fmt.Errorf("%wexpected bool, got number", ErrHTTPBadRequest)
	case float32:
		if !strict {
			switch t {
			case 0:
				return false, nil
			case 1:
				return true, nil
			}
		}
		return false, fmt.Errorf("%wexpected bool, got number", ErrHTTPBadRequest)
	case float64:
		if !strict {
			switch t {
			case 0:
				return false, nil
			case 1:
				return true, nil
			}
		}
		return false, fmt.Errorf("%wexpected bool, got number", ErrHTTPBadRequest)
	case string:
		if !strict {
			if b, err := strconv.ParseBool(t); err == nil {
				return b, nil
			}
		}

		return false, fmt.Errorf("%wexpected bool, got string", ErrHTTPBadRequest)
	case bool:
		return t, nil
	default:
		return false, fmt.Errorf("not handle for %T", t)
	}
}

func any2string(a any, strict bool) (string, error) {
	switch t := a.(type) {
	case int:
		if !strict {
			return strconv.FormatInt(int64(t), 10), nil
		}
		return "", fmt.Errorf("%wexpected string, got number", ErrHTTPBadRequest)
	case int8:
		if !strict {
			return strconv.FormatInt(int64(t), 10), nil
		}
		return "", fmt.Errorf("%wexpected string, got number", ErrHTTPBadRequest)
	case int16:
		if !strict {
			return strconv.FormatInt(int64(t), 10), nil
		}
		return "", fmt.Errorf("%wexpected string, got number", ErrHTTPBadRequest)
	case int32:
		if !strict {
			return strconv.FormatInt(int64(t), 10), nil
		}
		return "", fmt.Errorf("%wexpected string, got number", ErrHTTPBadRequest)
	case int64:
		if !strict {
			return strconv.FormatInt(int64(t), 10), nil
		}
		return "", fmt.Errorf("%wexpected string, got number", ErrHTTPBadRequest)
	case float32:
		if !strict {
			return strconv.FormatFloat(float64(t), 'f', -1, 32), nil
		}
		return "", fmt.Errorf("%wexpected string, got number", ErrHTTPBadRequest)
	case float64:
		if !strict {
			return strconv.FormatFloat(float64(t), 'f', -1, 64), nil
		}
		return "", fmt.Errorf("%wexpected string, got number", ErrHTTPBadRequest)
	case string:
		return t, nil
	case bool:
		if !strict {
			return strconv.FormatBool(t), nil
		}
		return "", fmt.Errorf("%wexpected string, got bool", ErrHTTPBadRequest)
	default:
		return "", fmt.Errorf("not handle for %T", t)
	}
}
