package apis

// APIVersions is a public varbiable that holds all the available API versions in the application.
type APIVersion struct {
	Pattern string `json:"pattern"`
}

var APIVersions = []APIVersion{
	{
		Pattern: "v1alpha1",
	},
}
