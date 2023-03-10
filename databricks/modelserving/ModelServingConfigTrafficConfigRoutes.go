package modelserving


type ModelServingConfigTrafficConfigRoutes struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/model_serving#served_model_name ModelServing#served_model_name}.
	ServedModelName *string `field:"required" json:"servedModelName" yaml:"servedModelName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/model_serving#traffic_percentage ModelServing#traffic_percentage}.
	TrafficPercentage *float64 `field:"required" json:"trafficPercentage" yaml:"trafficPercentage"`
}

