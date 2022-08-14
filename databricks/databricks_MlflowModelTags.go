// Prebuilt databricks Provider for Terraform CDK (cdktf)
package databricks


type MlflowModelTags struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mlflow_model#key MlflowModel#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mlflow_model#value MlflowModel#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

