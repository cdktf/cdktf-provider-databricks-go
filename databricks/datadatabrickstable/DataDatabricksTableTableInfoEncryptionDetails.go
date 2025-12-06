// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabrickstable


type DataDatabricksTableTableInfoEncryptionDetails struct {
	// sse_encryption_details block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/data-sources/table#sse_encryption_details DataDatabricksTable#sse_encryption_details}
	SseEncryptionDetails *DataDatabricksTableTableInfoEncryptionDetailsSseEncryptionDetails `field:"optional" json:"sseEncryptionDetails" yaml:"sseEncryptionDetails"`
}

