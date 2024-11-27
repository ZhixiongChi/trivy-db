package elxr

type Advisory struct {
	VulnerabilityID string
	Platform        string
	PkgName         string

	VendorIDs    []string
	State        string
	Severity     string
	FixedVersion string
	Title        string
}

type VulnerabilityDetail struct {
	Description string
}

type bucket struct {
	codeName string
	pkgName  string
	vulnID   string // CVE-ID, DLA-ID, DSA-ID or ESA-ID
	severity string
}

type header struct {
	ID          string `json:"ID"`
	Description string `json:"Description"`
}

type annotation struct {
	Type        string
	Release     string
	Package     string
	Kind        string
	Version     string
	Description string
	Severity    string
	Bugs        []string
}

type bug struct {
	Header      header
	Annotations []annotation
}
