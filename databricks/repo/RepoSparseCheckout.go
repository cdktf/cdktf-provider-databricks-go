package repo


type RepoSparseCheckout struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/repo#patterns Repo#patterns}.
	Patterns *[]*string `field:"required" json:"patterns" yaml:"patterns"`
}

