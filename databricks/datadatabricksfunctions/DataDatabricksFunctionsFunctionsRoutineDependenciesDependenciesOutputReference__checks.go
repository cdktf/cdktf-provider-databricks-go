// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package datadatabricksfunctions

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func (d *jsiiProxy_DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesOutputReference) validateInterpolationForAttributeParameters(property *string) error {
	if property == nil {
		return fmt.Errorf("parameter property is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesOutputReference) validatePutConnectionParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesConnection:
		value := value.(*[]*DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesConnection)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesConnection:
		value_ := value.([]*DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesConnection)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesConnection; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (d *jsiiProxy_DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesOutputReference) validatePutCredentialParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesCredential:
		value := value.(*[]*DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesCredential)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesCredential:
		value_ := value.([]*DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesCredential)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesCredential; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (d *jsiiProxy_DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesOutputReference) validatePutFunctionParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesFunction:
		value := value.(*[]*DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesFunction)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesFunction:
		value_ := value.([]*DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesFunction)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesFunction; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (d *jsiiProxy_DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesOutputReference) validatePutTableParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesTable:
		value := value.(*[]*DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesTable)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesTable:
		value_ := value.([]*DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesTable)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesTable; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (d *jsiiProxy_DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesOutputReference) validateResolveParameters(_context cdktf.IResolveContext) error {
	if _context == nil {
		return fmt.Errorf("parameter _context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
	switch val.(type) {
	case *string:
		// ok
	case string:
		// ok
	case *float64:
		// ok
	case float64:
		// ok
	case *int:
		// ok
	case int:
		// ok
	case *uint:
		// ok
	case uint:
		// ok
	case *int8:
		// ok
	case int8:
		// ok
	case *int16:
		// ok
	case int16:
		// ok
	case *int32:
		// ok
	case int32:
		// ok
	case *int64:
		// ok
	case int64:
		// ok
	case *uint8:
		// ok
	case uint8:
		// ok
	case *uint16:
		// ok
	case uint16:
		// ok
	case *uint32:
		// ok
	case uint32:
		// ok
	case *uint64:
		// ok
	case uint64:
		// ok
	default:
		return fmt.Errorf("parameter val must be one of the allowed types: *string, *float64; received %#v (a %T)", val, val)
	}

	return nil
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktf.IResolvable:
		// ok
	case *DataDatabricksFunctionsFunctionsRoutineDependenciesDependencies:
		val := val.(*DataDatabricksFunctionsFunctionsRoutineDependenciesDependencies)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case DataDatabricksFunctionsFunctionsRoutineDependenciesDependencies:
		val_ := val.(DataDatabricksFunctionsFunctionsRoutineDependenciesDependencies)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktf.IResolvable, *DataDatabricksFunctionsFunctionsRoutineDependenciesDependencies; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesOutputReference) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewDataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesOutputReferenceParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	if complexObjectIndex == nil {
		return fmt.Errorf("parameter complexObjectIndex is required, but nil was provided")
	}

	if complexObjectIsFromSet == nil {
		return fmt.Errorf("parameter complexObjectIsFromSet is required, but nil was provided")
	}

	return nil
}

