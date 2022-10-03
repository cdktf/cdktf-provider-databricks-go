package sqlendpoint


type SqlEndpointTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/sql_endpoint#create SqlEndpoint#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
}

