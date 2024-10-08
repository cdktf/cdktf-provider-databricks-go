// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package librarypluginframework

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LibraryPluginframeworkConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.53.0/docs/resources/library_pluginframework#cluster_id LibraryPluginframework#cluster_id}.
	ClusterId *string `field:"required" json:"clusterId" yaml:"clusterId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.53.0/docs/resources/library_pluginframework#cran LibraryPluginframework#cran}.
	Cran *LibraryPluginframeworkCran `field:"optional" json:"cran" yaml:"cran"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.53.0/docs/resources/library_pluginframework#egg LibraryPluginframework#egg}.
	Egg *string `field:"optional" json:"egg" yaml:"egg"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.53.0/docs/resources/library_pluginframework#jar LibraryPluginframework#jar}.
	Jar *string `field:"optional" json:"jar" yaml:"jar"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.53.0/docs/resources/library_pluginframework#maven LibraryPluginframework#maven}.
	Maven *LibraryPluginframeworkMaven `field:"optional" json:"maven" yaml:"maven"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.53.0/docs/resources/library_pluginframework#pypi LibraryPluginframework#pypi}.
	Pypi *LibraryPluginframeworkPypi `field:"optional" json:"pypi" yaml:"pypi"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.53.0/docs/resources/library_pluginframework#requirements LibraryPluginframework#requirements}.
	Requirements *string `field:"optional" json:"requirements" yaml:"requirements"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.53.0/docs/resources/library_pluginframework#whl LibraryPluginframework#whl}.
	Whl *string `field:"optional" json:"whl" yaml:"whl"`
}

