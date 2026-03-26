package goresty

// MapType converts an OpenAPI type and format pair to a Go type string.
func MapType(openAPIType, format string) string {
	switch openAPIType {
	case "string":
		switch format {
		case "date-time", "date":
			return "time.Time"
		}
		return "string"
	case "integer":
		switch format {
		case "int32":
			return "int32"
		case "int64":
			return "int64"
		}
		return "int"
	case "number":
		if format == "float" {
			return "float32"
		}
		return "float64"
	case "boolean":
		return "bool"
	case "object":
		return "map[string]any"
	case "array":
		return "[]any"
	default:
		return "any"
	}
}
