package awspol

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCondition(t *testing.T) {
	cStr := `{"StringLike":{"s3:prefix":["","home/","home/${aws:username}/"]}}`
	c := ConditionOperator{}
	err := json.Unmarshal([]byte(cStr), &c)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(c))
	assert.Equal(t, 1, len(c["StringLike"]))
	assert.Equal(t, 3, len(c["StringLike"]["s3:prefix"].Array))

	m, err := json.Marshal(c)
	assert.NoError(t, err)
	assert.Equal(t, cStr, string(m))
}
