package modelserving


type ModelServingTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/model_serving#create ModelServing#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/model_serving#update ModelServing#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

