package hermes

import (
	"fmt"
	"strings"

	Utils "github.com/realTristan/Hermes/utils"
)

// NewTempStorage is a function that creates a new TempStorage object for a given FullText object.
// Parameters:
//   - ft (*FullText): A pointer to the FullText object to create the TempStorage object for.
//
// Returns:
//   - (*TempStorage): A pointer to the newly created TempStorage object.
type TempStorage struct {
	data    map[string]any
	indices map[int]string
	index   int
	keys    map[string]int
}

// NewTempStorage is a function that creates a new TempStorage object for a given FullText object.
// Parameters:
//   - ft (*FullText): A pointer to the FullText object to create the TempStorage object for.
//
// Returns:
//   - (*TempStorage): A pointer to the newly created TempStorage object.
func NewTempStorage(ft *FullText) *TempStorage {
	var ts = &TempStorage{
		data:    ft.storage,
		indices: ft.indices,
		index:   ft.index,
		keys:    make(map[string]int),
	}

	// Loop through the data
	for k, v := range ts.indices {
		ts.keys[v] = k
	}
	return ts
}

// updateFullText is a method of the TempStorage struct that updates the FullText object with the data in the TempStorage object.
// Parameters:
//   - ft (*FullText): A pointer to the FullText object to update.
//
// Returns:
//   - None.
func (ts *TempStorage) updateFullText(ft *FullText) {
	ft.storage = ts.data
	ft.indices = ts.indices
	ft.index = ts.index
}

// cleanSingleArrays is a method of the TempStorage struct that replaces single-element integer arrays with their single integer value.
// Parameters:
//   - None.
//
// Returns:
//   - None.
func (ts *TempStorage) cleanSingleArrays() {
	for k, v := range ts.data {
		if v, ok := v.([]int); ok && len(v) == 1 {
			ts.data[k] = v[0]
		}
	}
}

// error is a method of the TempStorage struct that checks if the storage limit has been reached and returns an error if it has.
// Parameters:
//   - ft (*FullText): A pointer to the FullText object to check the storage limit against.
//
// Returns:
//   - (error): An error if the storage limit has been reached, nil otherwise.
func (ts *TempStorage) error(ft *FullText) error {
	// Check if the storage limit has been reached
	if ft.maxLength > 0 {
		if len(ts.data) > ft.maxLength {
			return fmt.Errorf("full-text storage limit reached (%d/%d keys). load cancelled", len(ts.data), ft.maxLength)
		}
	}
	if ft.maxBytes > 0 {
		if cacheSize, err := Utils.Size(ts.data); err != nil {
			return err
		} else if cacheSize > ft.maxBytes {
			return fmt.Errorf("full-text byte-size limit reached (%d/%d bytes). load cancelled", cacheSize, ft.maxBytes)
		}
	}
	return nil
}

// update is a method of the TempStorage struct that updates the TempStorage object with the given words and cache key.
// Parameters:
//   - ft (*FullText): A pointer to the FullText object to update.
//   - words ([]string): A slice of strings representing the words to update.
//   - cacheKey (string): A string representing the cache key to update.
//
// Returns:
//   - None.
func (ts *TempStorage) update(ft *FullText, words []string, cacheKey string) {
	// Loop through the words
	for i := 0; i < len(words); i++ {
		var word string = words[i]

		// Check if the word is valid
		if len(word) < ft.minWordLength {
			continue
		}
		if temp, ok := ts.data[word]; !ok {
			ts.data[word] = []int{ts.index}
		} else if v, ok := temp.([]int); !ok {
			ts.data[word] = []int{temp.(int), ts.keys[cacheKey]}
		} else {
			if Utils.SliceContains(v, ts.keys[cacheKey]) {
				continue
			}
			ts.data[word] = append(v, ts.keys[cacheKey])
		}
	}
}

// updateKeys is a method of the TempStorage struct that sets the given cache key in the temp storage keys
// Parameters:
//   - cacheKey (string): A string representing the cache key to set.
//
// Returns:
//   - None.
func (ts *TempStorage) updateKeys(cacheKey string) {
	if _, ok := ts.keys[cacheKey]; !ok {
		ts.index++
		ts.indices[ts.index] = cacheKey
		ts.keys[cacheKey] = ts.index
	}
}

// insertWords is a method of the TempStorage struct that inserts data into the temp storage.
// Parameters:
//   - ft (*FullText): A pointer to the FullText object to check the storage limit against.
//   - cacheKey (string): A string representing the cache key to insert.
//   - ftv (string): A string representing the value to insert.
//
// Returns:
//   - (error): An error if the storage limit has been reached, nil otherwise.
func (ts *TempStorage) insert(ft *FullText, cacheKey string, ftv string) error {
	// Set the cache key in the temp storage keys
	ts.updateKeys(cacheKey)

	// Clean the string value
	ftv = strings.TrimSpace(ftv)
	ftv = Utils.RemoveDoubleSpaces(ftv)
	ftv = strings.ToLower(ftv)

	// Loop through the words
	for _, word := range strings.Split(ftv, " ") {
		if len(word) == 0 {
			continue
		} else if len(word) < ft.minWordLength {
			continue
		} else if err := ts.error(ft); err != nil {
			return err
		}

		// Trim the word
		word = Utils.TrimNonAlphaNum(word)
		var words []string = Utils.SplitByAlphaNum(word)

		// Update the temp storage
		ts.update(ft, words, cacheKey)
	}

	// Return no error
	return nil
}
