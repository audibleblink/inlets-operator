package main


// InfraConfig is the configuration for
// creating Infrastructure Resources
type InfraConfig struct {
	Provider        string
	Region          string
	Zone            string
	AccessKey       string
	SecretKey       string
	OrganizationID  string
	SubscriptionID  string
	VpcID           string
	SubnetID        string
	AccessKeyFile   string
	SecretKeyFile   string
	ProjectID       string
	AnnotatedOnly   bool
	MaxClientMemory string
	Plan            string
	ProConfig       InletsProConfig
}

type InletsProConfig struct {
	License       string
	LicenseFile   string
	ClientImage   string
	InletsRelease string
}

func (c InletsProConfig) GetLicenseKey() (string, error) {
	return "asdf", nil
}
