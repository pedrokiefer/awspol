package awspol

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrincipalWildcard(t *testing.T) {
	p := &Principal{}
	err := json.Unmarshal([]byte(`"*"`), p)
	assert.NoError(t, err)
	assert.True(t, p.Any)

	m, err := json.Marshal(p)
	assert.NoError(t, err)
	assert.Equal(t, []byte(`"*"`), m)
}

func TestPrincipalAWS(t *testing.T) {
	p := &Principal{}
	err := json.Unmarshal([]byte(`{"AWS": "*"}`), p)
	assert.NoError(t, err)
	assert.False(t, p.Any)
	assert.Equal(t, 1, len(p.Map))
	assert.Equal(t, "*", p.Map["AWS"].Value)

	m, err := json.Marshal(p)
	assert.NoError(t, err)
	assert.Equal(t, []byte(`{"AWS":"*"}`), m)
}

func TestPrincipalInvalid(t *testing.T) {
	p := &Principal{}
	err := json.Unmarshal([]byte(`{"Something": "*"}`), p)
	assert.Error(t, err)
	assert.Equal(t, err, PrincipalError{Principal: "Something"})
}

func TestMultiplePrincipals(t *testing.T) {
	pData := `{ 
	"AWS": [
		"arn:aws:iam::123456789012:root",
		"999999999999"
	],
	"CanonicalUser": "79a59df900b949e55d96a1e698fbacedfd6e09d98eacf8f8d5218e7cd47ef2be"
}`

	p := &Principal{}
	err := json.Unmarshal([]byte(pData), p)
	assert.NoError(t, err)
	assert.False(t, p.Any)
	assert.Equal(t, 2, len(p.Map["AWS"].Array))
	assert.Equal(t, "79a59df900b949e55d96a1e698fbacedfd6e09d98eacf8f8d5218e7cd47ef2be", p.Map["CanonicalUser"].Value)
}
