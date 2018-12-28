package apis

// APIVersions is a public varbiable that holds all the available API versions in the application.
type APIVersion struct {
	Version string `json:"version"`
}

var APIVersions = []APIVersion{
	{
		Version: "v1alpha1",
	},
}
