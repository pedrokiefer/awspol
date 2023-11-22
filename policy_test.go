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

func TestS3UserAccessPolicy(t *testing.T) {
	pStr := `{
	"Version": "2012-10-17",
	"Statement": [
		{
			"Effect": "Allow",
			"Action": [
				"s3:ListAllMyBuckets",
				"s3:GetBucketLocation"
			],
			"Resource": "arn:aws:s3:::*"
		},
		{
			"Effect": "Allow",
			"Action": "s3:ListBucket",
			"Resource": "arn:aws:s3:::BUCKET-NAME",
			"Condition": {
				"StringLike": {
					"s3:prefix": [
						"",
						"home/",
						"home/${aws:username}/"
					]
				}
			}
		},
		{
			"Effect": "Allow",
			"Action": "s3:*",
			"Resource": [
				"arn:aws:s3:::BUCKET-NAME/home/${aws:username}",
				"arn:aws:s3:::BUCKET-NAME/home/${aws:username}/*"
			]
		}
	]
}
`

	pol, err := UnmarshalPolicy([]byte(pStr))
	assert.NoError(t, err)
	assert.Equal(t, Version2012, pol.Version)
	assert.Equal(t, 3, pol.Statements.Len())

	mPol, err := pol.MarshalPretty()
	assert.NoError(t, err)
	assert.Equal(t, pStr, string(mPol))

}
