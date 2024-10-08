// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mlflowmodel


type MlflowModelTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.53.0/docs/resources/mlflow_model#key MlflowModel#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.53.0/docs/resources/mlflow_model#value MlflowModel#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

