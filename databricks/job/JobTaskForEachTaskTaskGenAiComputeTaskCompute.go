// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job


type JobTaskForEachTaskTaskGenAiComputeTaskCompute struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/resources/job#num_gpus Job#num_gpus}.
	NumGpus *float64 `field:"required" json:"numGpus" yaml:"numGpus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/resources/job#gpu_node_pool_id Job#gpu_node_pool_id}.
	GpuNodePoolId *string `field:"optional" json:"gpuNodePoolId" yaml:"gpuNodePoolId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/resources/job#gpu_type Job#gpu_type}.
	GpuType *string `field:"optional" json:"gpuType" yaml:"gpuType"`
}

