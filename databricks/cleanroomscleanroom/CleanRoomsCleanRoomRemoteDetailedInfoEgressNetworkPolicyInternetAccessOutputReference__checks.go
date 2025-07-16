// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package cleanroomscleanroom

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func (c *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) validateInterpolationForAttributeParameters(property *string) error {
	if property == nil {
		return fmt.Errorf("parameter property is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) validatePutAllowedInternetDestinationsParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessAllowedInternetDestinations:
		value := value.(*[]*CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessAllowedInternetDestinations)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessAllowedInternetDestinations:
		value_ := value.([]*CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessAllowedInternetDestinations)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessAllowedInternetDestinations; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (c *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) validatePutAllowedStorageDestinationsParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessAllowedStorageDestinations:
		value := value.(*[]*CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessAllowedStorageDestinations)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessAllowedStorageDestinations:
		value_ := value.([]*CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessAllowedStorageDestinations)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessAllowedStorageDestinations; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (c *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) validatePutLogOnlyModeParameters(value *CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessLogOnlyMode) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) validateResolveParameters(_context cdktf.IResolveContext) error {
	if _context == nil {
		return fmt.Errorf("parameter _context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktf.IResolvable:
		// ok
	case *CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccess:
		val := val.(*CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccess)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccess:
		val_ := val.(CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccess)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktf.IResolvable, *CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccess; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) validateSetRestrictionModeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewCleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReferenceParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

