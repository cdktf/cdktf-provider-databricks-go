// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package databaseinstance

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DatabaseInstanceChildInstanceRefsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DatabaseInstanceChildInstanceRefsList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DatabaseInstanceChildInstanceRefsList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DatabaseInstanceChildInstanceRefsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DatabaseInstanceChildInstanceRefsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DatabaseInstanceChildInstanceRefsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DatabaseInstanceChildInstanceRefsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDatabaseInstanceChildInstanceRefsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

