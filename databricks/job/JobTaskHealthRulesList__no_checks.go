// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package job

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_JobTaskHealthRulesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (j *jsiiProxy_JobTaskHealthRulesList) validateGetParameters(index *float64) error {
	return nil
}

func (j *jsiiProxy_JobTaskHealthRulesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_JobTaskHealthRulesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_JobTaskHealthRulesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_JobTaskHealthRulesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_JobTaskHealthRulesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewJobTaskHealthRulesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

