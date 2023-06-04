package cache

// WFT is a struct that represents a value to be set in the cache and in the full-text cache.
type WFT struct {
	value string
}

// WithFT is a function that creates a new WFT struct with the specified value.
//
// Parameters:
//   - value: A string representing the value to set in the cache and in the full-text cache.
//
// Returns:
//   - A WFT struct with the specified value or the initial string
func (cache *Cache) WithFT(value string) any {
	if cache.ft == nil {
		return value
	}
	return WFT{value}
}

// ftFromMap is a function that gets the full-text value from a map.
//
// Parameters:
//   - value: any representing the value to get the full-text value from.
//
// Returns:
//   - A string representing the full-text value, or an empty string if the value is not a map or does not contain the correct keys.
func ftFromMap(value any) string {
	if _, ok := value.(map[string]any); !ok {
		return ""
	}

	// Verify that the map has the correct length
	var v map[string]any = value.(map[string]any)
	if len(v) != 2 {
		return ""
	}

	// Verify that the map has the correct keys
	if ft, ok := v["$hermes.full_text"]; ok {
		if ft, ok := ft.(bool); ok {
			if !ft {
				return ""
			} else if v, ok := v["$hermes.value"]; ok {
				if v, ok := v.(string); ok {
					return v
				}
			}
		}
	}
	return ""
}
