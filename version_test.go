package awspol

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckValidVersion(t *testing.T) {
	tests := []struct {
		name        string
		policy      string
		wantErr     bool
		expectedErr error
	}{
		{
			name:        "valid version",
			policy:      "{\"Version\":\"2012-10-17\",\"Statement\":[{\"Sid\":\"Stmt123456789\",\"Effect\":\"Allow\",\"Action\":\"sts:GetCallerIdentity\",\"Resource\":\"*\"}]}",
			wantErr:     false,
			expectedErr: nil,
		},
		{
			name:        "valid legacy version",
			policy:      "{\"Version\":\"2008-10-17\",\"Statement\":[{\"Sid\":\"Stmt123456789\",\"Effect\":\"Allow\",\"Action\":\"sts:GetCallerIdentity\",\"Resource\":\"*\"}]}",
			wantErr:     false,
			expectedErr: nil,
		},
		{
			name:        "valid version",
			policy:      "{\"Version\":\"2020-10-17\",\"Statement\":[{\"Sid\":\"Stmt123456789\",\"Effect\":\"Allow\",\"Action\":\"sts:GetCallerIdentity\",\"Resource\":\"*\"}]}",
			wantErr:     true,
			expectedErr: PolicyVersionError{Version: "2020-10-17"},
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
