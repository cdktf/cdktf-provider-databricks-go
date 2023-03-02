package mount


type MountTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mount#default Mount#default}.
	Default *string `field:"optional" json:"default" yaml:"default"`
}

