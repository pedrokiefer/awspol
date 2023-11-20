package awspol

import "encoding/json"

type ArrayOrValue[T any] struct {
	Value T
	Array []T
}

func (aov ArrayOrValue[T]) Len() int {
	if len(aov.Array) == 0 {
		return 1
	}
	return len(aov.Array)
}

func (aov ArrayOrValue[T]) MarshalJSON() ([]byte, error) {
	if len(aov.Array) > 0 {
		return json.Marshal(aov.Array)
	}
	return json.Marshal(aov.Value)
}

func (aov *ArrayOrValue[T]) UnmarshalJSON(data []byte) error {
	if data[0] == '[' {
		return json.Unmarshal(data, &aov.Array)
	}
	return json.Unmarshal(data, &aov.Value)
}
