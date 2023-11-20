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
