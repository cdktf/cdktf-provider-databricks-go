// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package qualitymonitor

// Building without runtime type checking enabled, so all the below just return nil

func (q *jsiiProxy_QualityMonitorCustomMetricsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (q *jsiiProxy_QualityMonitorCustomMetricsList) validateGetParameters(index *float64) error {
	return nil
}

func (q *jsiiProxy_QualityMonitorCustomMetricsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_QualityMonitorCustomMetricsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_QualityMonitorCustomMetricsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_QualityMonitorCustomMetricsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_QualityMonitorCustomMetricsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewQualityMonitorCustomMetricsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

