// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package modelserving


type ModelServingConfigA struct {
	// auto_capture_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.67.0/docs/resources/model_serving#auto_capture_config ModelServing#auto_capture_config}
	AutoCaptureConfig *ModelServingConfigAutoCaptureConfig `field:"optional" json:"autoCaptureConfig" yaml:"autoCaptureConfig"`
	// served_entities block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.67.0/docs/resources/model_serving#served_entities ModelServing#served_entities}
	ServedEntities interface{} `field:"optional" json:"servedEntities" yaml:"servedEntities"`
	// served_models block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.67.0/docs/resources/model_serving#served_models ModelServing#served_models}
	ServedModels interface{} `field:"optional" json:"servedModels" yaml:"servedModels"`
	// traffic_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.67.0/docs/resources/model_serving#traffic_config ModelServing#traffic_config}
	TrafficConfig *ModelServingConfigTrafficConfig `field:"optional" json:"trafficConfig" yaml:"trafficConfig"`
}

