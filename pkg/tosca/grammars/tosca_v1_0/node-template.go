package tosca_v1_0

import (
	"github.com/tliron/puccini/tosca/grammars/tosca_v1_3"
	"milkyway/pkg/tosca"
)

//
// NodeTemplate
//
// [TOSCA-Simple-Profile-YAML-v1.0] @ 3.7.3
//

// tosca.Reader signature
func ReadNodeTemplate(context *tosca.Context) tosca.EntityPtr {
	context.SetReadTag("Metadata", "")

	return tosca_v1_3.ReadNodeTemplate(context)
}
