package modelserving


type ModelServingConfigA struct {
	// served_models block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/model_serving#served_models ModelServing#served_models}
	ServedModels interface{} `field:"required" json:"servedModels" yaml:"servedModels"`
	// traffic_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/model_serving#traffic_config ModelServing#traffic_config}
	TrafficConfig *ModelServingConfigTrafficConfig `field:"optional" json:"trafficConfig" yaml:"trafficConfig"`
}

