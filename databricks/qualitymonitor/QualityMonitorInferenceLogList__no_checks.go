// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package qualitymonitor

// Building without runtime type checking enabled, so all the below just return nil

func (q *jsiiProxy_QualityMonitorInferenceLogList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (q *jsiiProxy_QualityMonitorInferenceLogList) validateGetParameters(index *float64) error {
	return nil
}

func (q *jsiiProxy_QualityMonitorInferenceLogList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_QualityMonitorInferenceLogList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_QualityMonitorInferenceLogList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_QualityMonitorInferenceLogList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_QualityMonitorInferenceLogList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewQualityMonitorInferenceLogListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

