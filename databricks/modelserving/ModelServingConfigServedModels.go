package modelserving


type ModelServingConfigServedModels struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.22.0/docs/resources/model_serving#model_name ModelServing#model_name}.
	ModelName *string `field:"required" json:"modelName" yaml:"modelName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.22.0/docs/resources/model_serving#model_version ModelServing#model_version}.
	ModelVersion *string `field:"required" json:"modelVersion" yaml:"modelVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.22.0/docs/resources/model_serving#workload_size ModelServing#workload_size}.
	WorkloadSize *string `field:"required" json:"workloadSize" yaml:"workloadSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.22.0/docs/resources/model_serving#environment_vars ModelServing#environment_vars}.
	EnvironmentVars *map[string]*string `field:"optional" json:"environmentVars" yaml:"environmentVars"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.22.0/docs/resources/model_serving#name ModelServing#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.22.0/docs/resources/model_serving#scale_to_zero_enabled ModelServing#scale_to_zero_enabled}.
	ScaleToZeroEnabled interface{} `field:"optional" json:"scaleToZeroEnabled" yaml:"scaleToZeroEnabled"`
}

