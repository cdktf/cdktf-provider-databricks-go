package mlflowmodel


type MlflowModelTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.15.0/docs/resources/mlflow_model#key MlflowModel#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.15.0/docs/resources/mlflow_model#value MlflowModel#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

