package awspol

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckValidStatements(t *testing.T) {
	tests := []struct {
		name        string
		policy      string
		wantErr     bool
		expectedErr error
	}{
		{
			name:        "valid version",
			policy:      "{\"Version\":\"2012-10-17\",\"Statement\":[{\"Sid\":\"Stmt123456789\",\"Effect\":\"Allow\",\"Action\":\"sts:GetCallerIdentity\",\"NotAction\":\"sts:GetCallerIdentity\",\"Resource\":\"*\"}]}",
			wantErr:     true,
			expectedErr: InvalidStatementError{"Action and NotAction cannot both be set"},
		},
		{
			name:        "valid version",
			policy:      "{\"Version\":\"2012-10-17\",\"Statement\":[{\"Sid\":\"Stmt123456789\",\"Effect\":\"Allow\",\"Resource\":\"sts:GetCallerIdentity\",\"NotResource\":\"sts:GetCallerIdentity\",\"Resource\":\"*\"}]}",
			wantErr:     true,
			expectedErr: InvalidStatementError{"Resource and NotResource cannot both be set"},
		},
		{
			name:        "valid version",
			policy:      "{\"Version\":\"2012-10-17\",\"Statement\":[{\"Sid\":\"Stmt123456789\",\"Effect\":\"Allow\",\"Principal\":\"*\",\"NotPrincipal\":\"*\",\"Resource\":\"*\"}]}",
			wantErr:     true,
			expectedErr: InvalidStatementError{"Principal and NotPrincipal cannot both be set"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := UnmarshalPolicy([]byte(tt.policy))
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}
