package grammars

import (
	"fmt"

	_ "github.com/tliron/puccini/tosca/profiles"
	"milkyway/pkg/tosca"
	"milkyway/pkg/tosca/grammars/cloudify_v1_3"
	"milkyway/pkg/tosca/grammars/hot"
	"milkyway/pkg/tosca/grammars/tosca_v1_0"
	"milkyway/pkg/tosca/grammars/tosca_v1_1"
	"milkyway/pkg/tosca/grammars/tosca_v1_2"
	"milkyway/pkg/tosca/grammars/tosca_v1_3"
	"milkyway/pkg/tosca/grammars/tosca_v2_0"
)

func init() {
	initGrammar(&tosca_v1_0.Grammar)
	initGrammar(&tosca_v1_1.Grammar)
	initGrammar(&tosca_v1_2.Grammar)
	initGrammar(&tosca_v1_3.Grammar)
	initGrammar(&tosca_v2_0.Grammar)
	initGrammar(&cloudify_v1_3.Grammar)
	initGrammar(&hot.Grammar)
}

func initGrammar(grammar *tosca.Grammar) {
	for keyword, versions := range grammar.Versions {
		var grammars map[string]*tosca.Grammar
		var ok bool

		if grammars, ok = Grammars[keyword]; !ok {
			grammars = make(map[string]*tosca.Grammar)
			Grammars[keyword] = grammars
		}

		for _, version := range versions {
			if _, ok := grammars[version.Version]; !ok {
				grammars[version.Version] = grammar
			} else {
				panic(fmt.Sprintf("initGrammar version conflict: %s = %s", keyword, version.Version))
			}

			if version.ImplicitProfilePath != "" {
				var paths map[string]string
				if paths, ok = ImplicitProfilePaths[keyword]; !ok {
					paths = make(map[string]string)
					ImplicitProfilePaths[keyword] = paths
				}

				if _, ok := paths[version.Version]; !ok {
					paths[version.Version] = version.ImplicitProfilePath
				} else {
					panic(fmt.Sprintf("initGrammar implicit profile path conflict: %s = %s -> %s", keyword, version.Version, version.ImplicitProfilePath))
				}
			}
		}
	}
}
