package cloudify_v1_3

import (
	"milkyway/pkg/tosca"
)

//
// PropertyDefinition
//

type PropertyDefinition struct {
	*ParameterDefinition `name:"property definition"`

	Required *bool `read:"required"`
}

func NewPropertyDefinition(context *tosca.Context) *PropertyDefinition {
	return &PropertyDefinition{
		ParameterDefinition: NewParameterDefinition(context),
	}
}

// tosca.Reader signature
func ReadPropertyDefinition(context *tosca.Context) tosca.EntityPtr {
	self := NewPropertyDefinition(context)
	context.ValidateUnsupportedFields(context.ReadFields(self))
	return self
}

// tosca.Mappable interface
func (self *PropertyDefinition) GetKey() string {
	return self.Name
}

func (self *PropertyDefinition) Inherit(parentDefinition *PropertyDefinition) {
	logInherit.Debugf("property definition: %s", self.Name)

	self.ParameterDefinition.Inherit(parentDefinition.ParameterDefinition)

	if (self.Required == nil) && (parentDefinition.Required != nil) {
		self.Required = parentDefinition.Required
	}
}

func (self *PropertyDefinition) IsRequired() bool {
	// defaults to true
	return (self.Required == nil) || *self.Required
}

//
// PropertyDefinitions
//

type PropertyDefinitions map[string]*PropertyDefinition

func (self PropertyDefinitions) Inherit(parentDefinitions PropertyDefinitions) {
	for name, definition := range parentDefinitions {
		if _, ok := self[name]; !ok {
			self[name] = definition
		}
	}

	for name, definition := range self {
		if parentDefinition, ok := parentDefinitions[name]; ok {
			if definition != parentDefinition {
				definition.Inherit(parentDefinition)
			}
		}
	}
}
