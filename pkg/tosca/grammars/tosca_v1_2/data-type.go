package tosca_v1_2

import (
	"milkyway/pkg/tosca"
	"milkyway/pkg/tosca/grammars/tosca_v2_0"
)

//
// DataType
//
// [TOSCA-Simple-Profile-YAML-v1.2] @ 3.7.6
// [TOSCA-Simple-Profile-YAML-v1.1] @ 3.6.6
// [TOSCA-Simple-Profile-YAML-v1.0] @ 3.6.5
//

// tosca.Reader signature
func ReadDataType(context *tosca.Context) tosca.EntityPtr {
	context.SetReadTag("KeySchema", "")
	context.SetReadTag("EntrySchema", "")

	return tosca_v2_0.ReadDataType(context)
}
