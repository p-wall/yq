package yqlib

import "gopkg.in/yaml.v3"

func equalsOperator(d *dataTreeNavigator, context Context, expressionNode *ExpressionNode) (Context, error) {
	log.Debugf("-- equalsOperation")
	return crossFunction(d, context, expressionNode, isEquals)
}

func isEquals(d *dataTreeNavigator, context Context, lhs *CandidateNode, rhs *CandidateNode) (*CandidateNode, error) {
	value := false

	lhsNode := unwrapDoc(lhs.Node)
	rhsNode := unwrapDoc(rhs.Node)

	if lhsNode.Tag == "!!null" {
		value = (rhsNode.Tag == "!!null")
	} else if lhsNode.Kind == yaml.ScalarNode && rhsNode.Kind == yaml.ScalarNode {
		value = matchKey(lhsNode.Value, rhsNode.Value)
	}
	log.Debugf("%v == %v ? %v", NodeToString(lhs), NodeToString(rhs), value)
	return createBooleanCandidate(lhs, value), nil
}
