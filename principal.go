package awspol

import "encoding/json"

type Principal struct {
	Any bool
	Map map[string]ArrayOrValue[string]
}

type PrincipalError struct {
	Principal string
}

func (e PrincipalError) Error() string {
	return "invalid principal: " + e.Principal
}

func (p *Principal) UnmarshalJSON(data []byte) error {
	if string(data) == "null" || string(data) == `""` {
		return nil
	}

	var wildcard string
	if err := json.Unmarshal(data, &wildcard); err == nil {
		p.Any = true
		return nil
	}

	var m map[string]ArrayOrValue[string]
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	for k := range m {
		if k != "AWS" && k != "Service" && k != "Federated" && k != "CanonicalUser" {
			return PrincipalError{k}
		}
	}

	p.Map = m

	return nil
}

func (p Principal) MarshalJSON() ([]byte, error) {
	if p.Any {
		return json.Marshal("*")
	}
	return json.Marshal(p.Map)
}
