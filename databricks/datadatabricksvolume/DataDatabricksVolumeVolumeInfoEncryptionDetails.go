// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksvolume


type DataDatabricksVolumeVolumeInfoEncryptionDetails struct {
	// sse_encryption_details block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.70.0/docs/data-sources/volume#sse_encryption_details DataDatabricksVolume#sse_encryption_details}
	SseEncryptionDetails *DataDatabricksVolumeVolumeInfoEncryptionDetailsSseEncryptionDetails `field:"optional" json:"sseEncryptionDetails" yaml:"sseEncryptionDetails"`
}

