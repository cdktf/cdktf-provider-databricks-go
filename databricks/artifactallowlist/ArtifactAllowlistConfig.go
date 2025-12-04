// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package artifactallowlist

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ArtifactAllowlistConfig struct {
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
	// artifact_matcher block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/resources/artifact_allowlist#artifact_matcher ArtifactAllowlist#artifact_matcher}
	ArtifactMatcher interface{} `field:"required" json:"artifactMatcher" yaml:"artifactMatcher"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/resources/artifact_allowlist#artifact_type ArtifactAllowlist#artifact_type}.
	ArtifactType *string `field:"required" json:"artifactType" yaml:"artifactType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/resources/artifact_allowlist#created_at ArtifactAllowlist#created_at}.
	CreatedAt *float64 `field:"optional" json:"createdAt" yaml:"createdAt"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/resources/artifact_allowlist#created_by ArtifactAllowlist#created_by}.
	CreatedBy *string `field:"optional" json:"createdBy" yaml:"createdBy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/resources/artifact_allowlist#id ArtifactAllowlist#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/resources/artifact_allowlist#metastore_id ArtifactAllowlist#metastore_id}.
	MetastoreId *string `field:"optional" json:"metastoreId" yaml:"metastoreId"`
}

