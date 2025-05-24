// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package artifactallowlist


type ArtifactAllowlistArtifactMatcher struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.1/docs/resources/artifact_allowlist#artifact ArtifactAllowlist#artifact}.
	Artifact *string `field:"required" json:"artifact" yaml:"artifact"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.1/docs/resources/artifact_allowlist#match_type ArtifactAllowlist#match_type}.
	MatchType *string `field:"required" json:"matchType" yaml:"matchType"`
}

