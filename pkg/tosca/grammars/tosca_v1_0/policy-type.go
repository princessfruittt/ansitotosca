package tosca_v1_0

import (
	"milkyway/pkg/tosca"
	"milkyway/pkg/tosca/grammars/tosca_v2_0"
)

//
// PolicyType
//
// [TOSCA-Simple-Profile-YAML-v1.0] @ 3.6.11
//

// tosca.Reader signature
func ReadPolicyType(context *tosca.Context) tosca.EntityPtr {
	context.SetReadTag("TriggerDefinitions", "")

	return tosca_v2_0.ReadPolicyType(context)
}
