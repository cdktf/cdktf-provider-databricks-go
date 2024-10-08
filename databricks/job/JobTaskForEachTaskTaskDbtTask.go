// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job


type JobTaskForEachTaskTaskDbtTask struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.53.0/docs/resources/job#commands Job#commands}.
	Commands *[]*string `field:"required" json:"commands" yaml:"commands"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.53.0/docs/resources/job#catalog Job#catalog}.
	Catalog *string `field:"optional" json:"catalog" yaml:"catalog"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.53.0/docs/resources/job#profiles_directory Job#profiles_directory}.
	ProfilesDirectory *string `field:"optional" json:"profilesDirectory" yaml:"profilesDirectory"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.53.0/docs/resources/job#project_directory Job#project_directory}.
	ProjectDirectory *string `field:"optional" json:"projectDirectory" yaml:"projectDirectory"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.53.0/docs/resources/job#schema Job#schema}.
	Schema *string `field:"optional" json:"schema" yaml:"schema"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.53.0/docs/resources/job#source Job#source}.
	Source *string `field:"optional" json:"source" yaml:"source"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.53.0/docs/resources/job#warehouse_id Job#warehouse_id}.
	WarehouseId *string `field:"optional" json:"warehouseId" yaml:"warehouseId"`
}

