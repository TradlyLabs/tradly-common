package db

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type supportedArrayType interface {
	~string | ~int | ~int64 | ~uint | ~uint64 | ~float64 | ~float32
}

// FlatArray is a generic type for handling array values from database
// It supports string, int, int64, uint, uint64, float64, and float32

type FlatArray[T supportedArrayType] []T

// Scan implements the sql.Scanner interface
func (a *FlatArray[T]) Scan(value any) error {
	return scanArray(value, a)
}

// Value implements the driver.Valuer interface
func (a FlatArray[T]) Value() (driver.Value, error) {
	if len(a) == 0 {
		return "{}", nil
	}

	// Create PostgreSQL array format: {value1,value2,...}
	var builder strings.Builder
	builder.WriteString("{")

	for i, v := range a {
		if i > 0 {
			builder.WriteString(",")
		}

		// Convert each element to string
		var s string
		switch any(v).(type) {
		case string:
			s = fmt.Sprintf("'%s'", strings.ReplaceAll(any(v).(string), "'", "''"))
		default:
			s = fmt.Sprintf("%v", v)
		}

		builder.WriteString(s)
	}

	builder.WriteString("}")
	return builder.String(), nil
}

// scanArray converts database value to FlatArray[T]
func scanArray[T supportedArrayType](value any, dest *FlatArray[T]) error {
	if value == nil {
		*dest = nil
		return nil
	}

	var data []byte

	// Convert value to byte slice
	switch v := value.(type) {
	case string:
		data = []byte(v)
	case []byte:
		data = v
	default:
		return fmt.Errorf("unsupported type for FlatArray: %T", value)
	}

	// Handle PostgreSQL array format (e.g., '{value1,value2}')
	if len(data) > 0 && data[0] == '{' && data[len(data)-1] == '}' {
		// Remove curly braces
		data = data[1 : len(data)-1]

		// Split by comma
		items := strings.Split(string(data), ",")
		result := make(FlatArray[T], 0, len(items))

		for _, item := range items {
			item = strings.TrimSpace(item)
			if item == "" {
				continue
			}

			var val T
			// For string types, remove quotes and unescape single quotes
			var itemToConvert string
			if _, ok := any(val).(string); ok {
				// Check if item is quoted
				if len(item) >= 2 && item[0] == '\'' && item[len(item)-1] == '\'' {
					// Remove quotes
					itemToConvert = item[1 : len(item)-1]
					// Unescape single quotes ('' -> ')
					itemToConvert = strings.ReplaceAll(itemToConvert, "''", "'")
				} else {
					itemToConvert = item
				}
			} else {
				itemToConvert = item
			}

			if err := convertToType(itemToConvert, &val); err != nil {
				return fmt.Errorf("failed to convert item '%s' to type %T: %w", item, val, err)
			}

			result = append(result, val)
		}

		*dest = result
		return nil
	}

	// Try to unmarshal as JSON array
	var result FlatArray[T]
	if err := json.Unmarshal(data, &result); err != nil {
		// If JSON unmarshaling fails, try to parse as single value
		var singleVal T
		if err := convertToType(string(data), &singleVal); err != nil {
			return fmt.Errorf("failed to parse value as array or single value: %w", err)
		}
		result = FlatArray[T]{singleVal}
	}

	*dest = result
	return nil
}

// convertToType converts string to the target type T
func convertToType[T supportedArrayType](s string, dest *T) error {
	var zero T

	switch any(zero).(type) {
	case string:
		*dest = any(s).(T)

	case int:
		val, err := strconv.Atoi(s)
		if err != nil {
			return err
		}
		*dest = any(val).(T)

	case int64:
		val, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return err
		}
		*dest = any(val).(T)

	case uint:
		val, err := strconv.ParseUint(s, 10, 64)
		if err != nil {
			return err
		}
		*dest = any(uint(val)).(T)

	case uint64:
		val, err := strconv.ParseUint(s, 10, 64)
		if err != nil {
			return err
		}
		*dest = any(val).(T)

	case float64:
		val, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return err
		}
		*dest = any(val).(T)

	case float32:
		val, err := strconv.ParseFloat(s, 32)
		if err != nil {
			return err
		}
		*dest = any(float32(val)).(T)

	default:
		return errors.New("unsupported type for conversion")
	}

	return nil
}
