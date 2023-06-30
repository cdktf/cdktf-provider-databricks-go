package sqlwidget

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v9/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v9/sqlwidget/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SqlWidgetPositionOutputReference interface {
	cdktf.ComplexObject
	AutoHeight() interface{}
	SetAutoHeight(val interface{})
	AutoHeightInput() interface{}
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *SqlWidgetPosition
	SetInternalValue(val *SqlWidgetPosition)
	PosX() *float64
	SetPosX(val *float64)
	PosXInput() *float64
	PosY() *float64
	SetPosY(val *float64)
	PosYInput() *float64
	SizeX() *float64
	SetSizeX(val *float64)
	SizeXInput() *float64
	SizeY() *float64
	SetSizeY(val *float64)
	SizeYInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetAutoHeight()
	ResetPosX()
	ResetPosY()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SqlWidgetPositionOutputReference
type jsiiProxy_SqlWidgetPositionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SqlWidgetPositionOutputReference) AutoHeight() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoHeight",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlWidgetPositionOutputReference) AutoHeightInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoHeightInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlWidgetPositionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlWidgetPositionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlWidgetPositionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlWidgetPositionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlWidgetPositionOutputReference) InternalValue() *SqlWidgetPosition {
	var returns *SqlWidgetPosition
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlWidgetPositionOutputReference) PosX() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"posX",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlWidgetPositionOutputReference) PosXInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"posXInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlWidgetPositionOutputReference) PosY() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"posY",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlWidgetPositionOutputReference) PosYInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"posYInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlWidgetPositionOutputReference) SizeX() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sizeX",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlWidgetPositionOutputReference) SizeXInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sizeXInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlWidgetPositionOutputReference) SizeY() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sizeY",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlWidgetPositionOutputReference) SizeYInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sizeYInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlWidgetPositionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlWidgetPositionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSqlWidgetPositionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SqlWidgetPositionOutputReference {
	_init_.Initialize()

	if err := validateNewSqlWidgetPositionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SqlWidgetPositionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.sqlWidget.SqlWidgetPositionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSqlWidgetPositionOutputReference_Override(s SqlWidgetPositionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.sqlWidget.SqlWidgetPositionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SqlWidgetPositionOutputReference)SetAutoHeight(val interface{}) {
	if err := j.validateSetAutoHeightParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoHeight",
		val,
	)
}

func (j *jsiiProxy_SqlWidgetPositionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SqlWidgetPositionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SqlWidgetPositionOutputReference)SetInternalValue(val *SqlWidgetPosition) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SqlWidgetPositionOutputReference)SetPosX(val *float64) {
	if err := j.validateSetPosXParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"posX",
		val,
	)
}

func (j *jsiiProxy_SqlWidgetPositionOutputReference)SetPosY(val *float64) {
	if err := j.validateSetPosYParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"posY",
		val,
	)
}

func (j *jsiiProxy_SqlWidgetPositionOutputReference)SetSizeX(val *float64) {
	if err := j.validateSetSizeXParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sizeX",
		val,
	)
}

func (j *jsiiProxy_SqlWidgetPositionOutputReference)SetSizeY(val *float64) {
	if err := j.validateSetSizeYParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sizeY",
		val,
	)
}

func (j *jsiiProxy_SqlWidgetPositionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SqlWidgetPositionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SqlWidgetPositionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlWidgetPositionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlWidgetPositionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlWidgetPositionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlWidgetPositionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlWidgetPositionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlWidgetPositionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlWidgetPositionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlWidgetPositionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlWidgetPositionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlWidgetPositionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlWidgetPositionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlWidgetPositionOutputReference) ResetAutoHeight() {
	_jsii_.InvokeVoid(
		s,
		"resetAutoHeight",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlWidgetPositionOutputReference) ResetPosX() {
	_jsii_.InvokeVoid(
		s,
		"resetPosX",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlWidgetPositionOutputReference) ResetPosY() {
	_jsii_.InvokeVoid(
		s,
		"resetPosY",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlWidgetPositionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := s.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlWidgetPositionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

