package modelserving


type ModelServingConfigTrafficConfig struct {
	// routes block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/model_serving#routes ModelServing#routes}
	Routes interface{} `field:"optional" json:"routes" yaml:"routes"`
}

