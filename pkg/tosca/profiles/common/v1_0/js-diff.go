// This file was auto-generated from a YAML file

package v1_0

func init() {
	Profile["/tosca/common/1.0/js/diff.js"] = `

clout.exec('tosca.lib.traversal');

if (!puccini.arguments.base) {
    throw 'must provide "base" argument';
}

var base = clout.load(puccini.arguments.base);

tosca.coerce();
tosca.coerce(base);

puccini.write(diff(clout, base));

function diff(clout, base) {
    var nodes = gatherNodeTemplates(clout);
    var baseNodes = gatherNodeTemplates(base);
    
    var diff = {
        added: [],
        removed: []
    };
    
    for (var n = 0, l = nodes.length; n < l; n++) {
        var node = nodes[n];
        if (!hasNode(baseNodes, node))
            diff.added.push(node.name);
    }
    
    for (var n = 0, l = baseNodes.length; n < l; n++) {
        var node = baseNodes[n];
        if (!hasNode(nodes, node))
            diff.removed.push(node.name);
    }

    return diff;
}

function gatherNodeTemplates(clout) {
    var nodeTemplates = [];
    for (var vertexId in clout.vertexes) {
        var vertex = clout.vertexes[vertexId];
        if (tosca.isNodeTemplate(vertex)) {
            var nodeTemplate = vertex.properties;
            nodeTemplates.push(nodeTemplate);
        }
    }
    return nodeTemplates;
}

function hasNode(nodes, node) {
    for (var n = 0, l = nodes.length; n < l; n++)
        if (nodes[n].name === node.name)
            return true;
    return false;
}
`
}
