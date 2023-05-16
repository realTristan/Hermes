package cache

// Values is a method of the Cache struct that gets all the values in the cache.
// This function is thread-safe.
//
// Returns:
//   - A slice of map[string]interface{} representing all the values in the cache.
func (c *Cache) Values() []map[string]interface{} {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.values()
}

// values is a method of the Cache struct that returns all the values in the cache.
// This function is not thread-safe, and should only be called from an exported function.
//
// Returns:
//   - A slice of map[string]interface{} representing all the values in the cache.
func (c *Cache) values() []map[string]interface{} {
	var values []map[string]interface{} = []map[string]interface{}{}
	for _, value := range c.data {
		values = append(values, value)
	}
	return values
}
