package awspol

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanDeserializeLegacySchema(t *testing.T) {
	serializedPol := `{
	"Version": "2008-10-17",
	"Id": "Policy123456789",
	"Statement": {
		"Sid": "Stmt123456789",
		"Effect": "Allow",
		"Action": "sts:GetCallerIdentity",
		"Resource": "*"
	}
}
`

	pol, err := UnmarshalPolicy([]byte(serializedPol))
	assert.NoError(t, err)
	assert.Equal(t, Version2008, pol.Version)
	assert.Equal(t, "Policy123456789", pol.Id)
	assert.Equal(t, 1, pol.Statements.Len())

	mPol, err := pol.MarshalPretty()
	assert.NoError(t, err)
	assert.Equal(t, serializedPol, string(mPol))
}
