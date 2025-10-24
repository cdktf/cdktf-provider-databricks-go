// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package qualitymonitor

// Building without runtime type checking enabled, so all the below just return nil

func (q *jsiiProxy_QualityMonitorProviderConfigList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (q *jsiiProxy_QualityMonitorProviderConfigList) validateGetParameters(index *float64) error {
	return nil
}

func (q *jsiiProxy_QualityMonitorProviderConfigList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_QualityMonitorProviderConfigList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_QualityMonitorProviderConfigList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_QualityMonitorProviderConfigList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_QualityMonitorProviderConfigList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewQualityMonitorProviderConfigListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

