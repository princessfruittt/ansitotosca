package tosca_v1_3

import (
	"milkyway/pkg/tosca"
	"milkyway/pkg/tosca/grammars/tosca_v2_0"
)

//
// NodeTemplate
//
// [TOSCA-Simple-Profile-YAML-v1.3] @ 3.8.3
// [TOSCA-Simple-Profile-YAML-v1.2] @ 3.8.3
// [TOSCA-Simple-Profile-YAML-v1.1] @ 3.7.3
// [TOSCA-Simple-Profile-YAML-v1.0] @ 3.7.3
//

// tosca.Reader signature
func ReadNodeTemplate(context *tosca.Context) tosca.EntityPtr {
	self := tosca_v2_0.NewNodeTemplate(context)
	context.ValidateUnsupportedFields(context.ReadFields(self))
	switch self.Name {
	case "SELF", "SOURCE", "TARGET", "HOST":
		context.Clone(self.Name).ReportValueInvalid("node template name", "reserved")
	}
	return self
}
