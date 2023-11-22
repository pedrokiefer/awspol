package awspol

import "encoding/json"

type Effect string

var (
	Allow = Effect("Allow")
	Deny  = Effect("Deny")
)

type Statement struct {
	Sid          string                `json:"Sid,omitempty"`
	Principal    *Principal            `json:"Principal,omitempty"`
	NotPrincipal *Principal            `json:"NotPrincipal,omitempty"`
	Effect       Effect                `json:"Effect"`
	Action       *ArrayOrValue[string] `json:"Action,omitempty"`
	NotAction    *ArrayOrValue[string] `json:"NotAction,omitempty"`
	Resource     *ArrayOrValue[string] `json:"Resource,omitempty"`
	NotResource  *ArrayOrValue[string] `json:"NotResource,omitempty"`
	Condition    ConditionOperator     `json:"Condition,omitempty"`
}

type InvalidStatementError struct {
	Reason string
}

func (e InvalidStatementError) Error() string {
	return "invalid statement: " + e.Reason
}

func (s *Statement) UnmarshalJSON(data []byte) error {
	type lStatement Statement
	var l lStatement
	if err := json.Unmarshal(data, &l); err != nil {
		return err
	}

	if l.Principal != nil && l.NotPrincipal != nil {
		return InvalidStatementError{"Principal and NotPrincipal cannot both be set"}
	}

	if l.Action != nil && l.NotAction != nil {
		return InvalidStatementError{"Action and NotAction cannot both be set"}
	}

	if l.Resource != nil && l.NotResource != nil {
		return InvalidStatementError{"Resource and NotResource cannot both be set"}
	}

	s.Sid = l.Sid
	s.Effect = l.Effect
	s.Action = l.Action
	s.NotAction = l.NotAction
	s.Principal = l.Principal
	s.NotPrincipal = l.NotPrincipal
	s.Resource = l.Resource
	s.Condition = l.Condition

	return nil
}

func (s *Statement) MarshalJSON() ([]byte, error) {
	return json.Marshal(s)
}
