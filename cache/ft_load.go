package cache

import (
	"fmt"
	"strings"
)

/*
loadCacheData() loops through a map of data and extracts relevant information to populate the wordCache map in the FullText struct.

Parameters:
- data (map[string]map[string]interface{}): a map of data to be loaded into the cache
- schema (map[string]bool): a map representing the schema of the data; only keys in this map will be loaded into the cache

Returns:
- error: if an error occurs during the loading process, it is returned. Otherwise, returns nil.
*/
func (ft *FullText) loadCache(data map[string]map[string]interface{}, schema map[string]bool) (map[string][]string, error) {
	var temp map[string][]string = ft.wordCache

	// Loop through the json data
	for itemKey, itemValue := range data {
		// Loop through the map
		for key, value := range itemValue {
			// Check if the key is in the schema
			if !schema[key] {
				continue
			}

			// Check if the value is a string
			if v, ok := value.(string); ok {
				// Clean the value
				v = strings.TrimSpace(v)
				v = removeDoubleSpaces(v)
				v = strings.ToLower(v)

				// Loop through the words
				for _, word := range strings.Split(v, " ") {
					if ft.maxWords > 0 {
						if len(temp) > ft.maxWords {
							return ft.wordCache, fmt.Errorf("full text cache key limit reached (%d/%d keys). load cancelled", len(temp), ft.maxWords)
						}
					}
					if ft.maxSizeBytes > 0 {
						if cacheSize, err := size(temp); err != nil {
							return ft.wordCache, err
						} else if cacheSize > ft.maxSizeBytes {
							return ft.wordCache, fmt.Errorf("full text cache size limit reached (%d/%d bytes). load cancelled", cacheSize, ft.maxSizeBytes)
						}
					}
					switch {
					case len(word) <= 1:
						continue
					case !isAlphaNum(word):
						word = removeNonAlphaNum(word)
					}
					if _, ok := temp[word]; !ok {
						temp[word] = []string{itemKey}
						continue
					}
					if containsString(temp[word], itemKey) {
						continue
					}
					temp[word] = append(temp[word], itemKey)
				}
			}
		}
	}
	return temp, nil
}
