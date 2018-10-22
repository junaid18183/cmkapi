package cmkapi

// StructGetHostResult holds the results of a Check_MK webAPI answer
type StructGetHostResult struct {
	Result struct {
		Attributes struct {
			NetworkScan struct {
				ScanInterval  int           `json:"scan_interval"`
				ExcludeRanges []interface{} `json:"exclude_ranges"`
				IPRanges      []interface{} `json:"ip_ranges"`
				RunAs         string        `json:"run_as"`
			} `json:"network_scan"`
			TagAgent          string        `json:"tag_agent"`
			SnmpCommunity     string        `json:"snmp_community"`
			Ipv6Address       string        `json:"ipv6address"`
			Alias             string        `json:"alias"`
			TagCriticality    string        `json:"tag_criticality"`
			Site              string        `json:"site"`
			TagAddressFamily  string        `json:"tag_address_family"`
			Contactgroups     contactGroups `json:"contactgroups"`
			NetworkScanResult struct {
				Start  interface{} `json:"start"`
				State  interface{} `json:"state"`
				End    interface{} `json:"end"`
				Output string      `json:"output"`
			} `json:"network_scan_result"`
			Parents       []interface{} `json:"parents"`
			Ipaddress     string        `json:"ipaddress"`
			TagNetworking string        `json:"tag_networking"`
		} `json:"attributes"`
		Hostname string `json:"hostname"`
		Path     string `json:"path"`
	} `json:"result"`
	ResultCode int `json:"result_code"`
}

// Hosts contactgroups settings
type contactGroups struct {
	UseForServices bool     `json:"use_for_services"`
	RecursePerms   bool     `json:"recurse_perms"`
	RecurseUse     bool     `json:"recurse_use"`
	Use            bool     `json:"use"`
	Groups         []string `json:"groups"`
}

// ActivateResult holds the result of a Check_MK webAPI update
type ActivateResult struct {
	Result     activateResultSites `json:"result"`
	ResultCode int                 `json:"result_code"`
}

type activateResultSites struct {
	Sites map[string]activateResultSite `json:"sites"`
}

type activateResultSite struct {
	TimeUpdated      float64                `json:"_time_updated"`
	StatusDetails    string                 `json:"_status_details"`
	Phase            string                 `json:"_phase"`
	StatusText       string                 `json:"_status_text"`
	PID              int                    `json:"_pid"`
	State            string                 `json:"_state"`
	TimeEnded        float64                `json:"_time_ended"`
	ExpectedDuration float64                `json:"_expected_duration"`
	TimeStarted      float64                `json:"_time_started"`
	SiteID           string                 `json:"_site_id"`
	Warnings         activateResultWarnings `json:"_warnings"`
}

type activateResultWarnings struct {
	CACertificate []string `json:"ca-certificates"`
	CheckMK       []string `json:"check_mk"`
}

// CheckMKException holds exceptions thrown by CheckMK
type CheckMKException struct {
	Result     string `json:"result"`
	ResultCode int    `json:"result_code"`
}

// Attributes describes the needed/optional tags for the Check_MK webAPI
type Attributes struct {
	Alias          string `json:"alias"`
	TagAgent       string `json:"tag_agent"`
	TagCriticality string `json:"tag_criticality"`
	Ipaddress      string `json:"ipaddress"`
}

// Host holds the data for a single host
type Host struct {
	Attributes `json:"attributes"`
	Hostname   string `json:"hostname"`
	Folder     string `json:"folder"`
}
