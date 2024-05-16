// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksexternallocation


type DataDatabricksExternalLocationExternalLocationInfoEncryptionDetails struct {
	// sse_encryption_details block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.44.0/docs/data-sources/external_location#sse_encryption_details DataDatabricksExternalLocation#sse_encryption_details}
	SseEncryptionDetails *DataDatabricksExternalLocationExternalLocationInfoEncryptionDetailsSseEncryptionDetails `field:"optional" json:"sseEncryptionDetails" yaml:"sseEncryptionDetails"`
}

