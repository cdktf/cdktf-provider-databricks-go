// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job


type JobPythonWheelTask struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/resources/job#entry_point Job#entry_point}.
	EntryPoint *string `field:"optional" json:"entryPoint" yaml:"entryPoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/resources/job#named_parameters Job#named_parameters}.
	NamedParameters *map[string]*string `field:"optional" json:"namedParameters" yaml:"namedParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/resources/job#package_name Job#package_name}.
	PackageName *string `field:"optional" json:"packageName" yaml:"packageName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/resources/job#parameters Job#parameters}.
	Parameters *[]*string `field:"optional" json:"parameters" yaml:"parameters"`
}

