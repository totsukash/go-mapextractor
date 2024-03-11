package extractor

/**
 * `json.Unmarshal`関数は数値をデフォルトで`float64`型として解釈するため、
 * 数値は`float64`型として取得する必要があります。
 */

func Get[T any](m map[string]any, keys ...string) T {
	var empty T
	if len(keys) == 0 {
		return empty
	}

	key := keys[0]
	value := m[key]

	if len(keys) == 1 {
		v, ok := value.(T)
		if !ok {
			return empty
		}
		return v
	}

	nextMap, ok := value.(map[string]any)
	if !ok {
		return empty
	}
	nextKeys := keys[1:]

	return Get[T](nextMap, nextKeys...)
}
