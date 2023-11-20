package awspol

import "fmt"

type PolicyVersion string

var (
	Version2008 = PolicyVersion("2008-10-17")
	Version2012 = PolicyVersion("2012-10-17")
)

var validVersions = map[PolicyVersion]bool{
	Version2008: true,
	Version2012: true,
}

type PolicyVersionError struct {
	Version PolicyVersion
}

func (e PolicyVersionError) Error() string {
	return fmt.Sprintf("invalid policy version: %s", e.Version)
}
