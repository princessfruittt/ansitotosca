package tosca_v1_2

import (
	"github.com/tliron/puccini/tosca/grammars/tosca_v2_0"
	"milkyway/pkg/tosca"
)

//
// GroupType
//
// [TOSCA-Simple-Profile-YAML-v1.2] @ 3.7.11
// [TOSCA-Simple-Profile-YAML-v1.1] @ 3.6.11
// [TOSCA-Simple-Profile-YAML-v1.0] @ 3.6.10
//

// tosca.Reader signature
func ReadGroupType(context *tosca.Context) tosca.EntityPtr {
	context.SetReadTag("InterfaceDefinitions", "interfaces,InterfaceDefinition")

	return tosca_v2_0.ReadGroupType(context)
}