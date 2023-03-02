package datadatabrickssparkversion

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-databricks.dataDatabricksSparkVersion.DataDatabricksSparkVersion",
		reflect.TypeOf((*DataDatabricksSparkVersion)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "beta", GoGetter: "Beta"},
			_jsii_.MemberProperty{JsiiProperty: "betaInput", GoGetter: "BetaInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberProperty{JsiiProperty: "genomics", GoGetter: "Genomics"},
			_jsii_.MemberProperty{JsiiProperty: "genomicsInput", GoGetter: "GenomicsInput"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "gpu", GoGetter: "Gpu"},
			_jsii_.MemberProperty{JsiiProperty: "gpuInput", GoGetter: "GpuInput"},
			_jsii_.MemberProperty{JsiiProperty: "graviton", GoGetter: "Graviton"},
			_jsii_.MemberProperty{JsiiProperty: "gravitonInput", GoGetter: "GravitonInput"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "latest", GoGetter: "Latest"},
			_jsii_.MemberProperty{JsiiProperty: "latestInput", GoGetter: "LatestInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "longTermSupport", GoGetter: "LongTermSupport"},
			_jsii_.MemberProperty{JsiiProperty: "longTermSupportInput", GoGetter: "LongTermSupportInput"},
			_jsii_.MemberProperty{JsiiProperty: "ml", GoGetter: "Ml"},
			_jsii_.MemberProperty{JsiiProperty: "mlInput", GoGetter: "MlInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "photon", GoGetter: "Photon"},
			_jsii_.MemberProperty{JsiiProperty: "photonInput", GoGetter: "PhotonInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetBeta", GoMethod: "ResetBeta"},
			_jsii_.MemberMethod{JsiiMethod: "resetGenomics", GoMethod: "ResetGenomics"},
			_jsii_.MemberMethod{JsiiMethod: "resetGpu", GoMethod: "ResetGpu"},
			_jsii_.MemberMethod{JsiiMethod: "resetGraviton", GoMethod: "ResetGraviton"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetLatest", GoMethod: "ResetLatest"},
			_jsii_.MemberMethod{JsiiMethod: "resetLongTermSupport", GoMethod: "ResetLongTermSupport"},
			_jsii_.MemberMethod{JsiiMethod: "resetMl", GoMethod: "ResetMl"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPhoton", GoMethod: "ResetPhoton"},
			_jsii_.MemberMethod{JsiiMethod: "resetScala", GoMethod: "ResetScala"},
			_jsii_.MemberMethod{JsiiMethod: "resetSparkVersion", GoMethod: "ResetSparkVersion"},
			_jsii_.MemberProperty{JsiiProperty: "scala", GoGetter: "Scala"},
			_jsii_.MemberProperty{JsiiProperty: "scalaInput", GoGetter: "ScalaInput"},
			_jsii_.MemberProperty{JsiiProperty: "sparkVersion", GoGetter: "SparkVersion"},
			_jsii_.MemberProperty{JsiiProperty: "sparkVersionInput", GoGetter: "SparkVersionInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksSparkVersion{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformDataSource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-databricks.dataDatabricksSparkVersion.DataDatabricksSparkVersionConfig",
		reflect.TypeOf((*DataDatabricksSparkVersionConfig)(nil)).Elem(),
	)
}
