package tosca_v1_2

import (
	"milkyway/pkg/tosca"
	"milkyway/pkg/tosca/grammars/tosca_v2_0"
)

//
// InterfaceDefinition
//
// [TOSCA-Simple-Profile-YAML-v1.2] @ 3.6.16
// [TOSCA-Simple-Profile-YAML-v1.1] @ 3.5.14
// [TOSCA-Simple-Profile-YAML-v1.0] @ 3.5.14
//

// tosca.Reader signature
func ReadInterfaceDefinition(context *tosca.Context) tosca.EntityPtr {
	context.SetReadTag("OperationDefinitions", "?,OperationDefinition")
	context.SetReadTag("NotificationDefinitions", "")

	self := tosca_v2_0.NewInterfaceDefinition(context)
	context.ReadFields(self)
	return self
}
