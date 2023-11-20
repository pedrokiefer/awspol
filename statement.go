package awspol

type Effect string

var (
	Allow = Effect("Allow")
	Deny  = Effect("Deny")
)

type Statement struct {
	Sid          string                 `json:"Sid,omitempty"`
	Principal    *Principal             `json:"Principal,omitempty"`
	NotPrincipal *Principal             `json:"NotPrincipal,omitempty"`
	Effect       Effect                 `json:"Effect"`
	Action       *ArrayOrValue[string]  `json:"Action,omitempty"`
	NotAction    *ArrayOrValue[string]  `json:"NotAction,omitempty"`
	Resource     *ArrayOrValue[string]  `json:"Resource,omitempty"`
	NotResource  *ArrayOrValue[string]  `json:"NotResource,omitempty"`
	Condition    map[string]interface{} `json:"Condition,omitempty"`
}
