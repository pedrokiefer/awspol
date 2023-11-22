package awspol

import (
	"bytes"
	"encoding/json"
)

type Policy struct {
	Version    PolicyVersion           `json:"Version"`
	Id         string                  `json:"Id,omitempty"`
	Statements ArrayOrValue[Statement] `json:"Statement"`
}

// NewPolicy creates a new Policy object with default values.
func NewPolicy() Policy {
	return Policy{
		Version:    Version2012,
		Statements: NewArrayOrValue[Statement](),
	}
}

// UnmarshalPolicy unmarshals the given byte slice into a Policy struct.
// It returns the unmarshaled Policy and any error encountered during the unmarshaling process.
func UnmarshalPolicy(data []byte) (*Policy, error) {
	pol := Policy{}
	err := json.Unmarshal(data, &pol)

	if pol.Version == "" || !validVersions[pol.Version] {
		return nil, PolicyVersionError{pol.Version}
	}

	return &pol, err
}

func (pol *Policy) AddStatement(stmt Statement) {
	pol.Statements.Array = append(pol.Statements.Array, stmt)
}

func (pol Policy) Marshal() ([]byte, error) {
	return json.Marshal(pol)
}

func (pol Policy) MarshalPretty() ([]byte, error) {
	buf := bytes.Buffer{}
	enc := json.NewEncoder(&buf)
	enc.SetIndent("", "\t")
	err := enc.Encode(pol)
	return buf.Bytes(), err
}
