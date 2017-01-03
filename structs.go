package cmkapi

//#-------------------------------------------------------------------------------------------------------------------------------------------
type StructGetHostResult struct {
        Result struct {
                Attributes struct {
                        NetworkScan struct {
                                ScanInterval int `json:"scan_interval"`
                                ExcludeRanges []interface{} `json:"exclude_ranges"`
                                IPRanges []interface{} `json:"ip_ranges"`
                                RunAs string `json:"run_as"`
                        } `json:"network_scan"`
                        TagAgent string `json:"tag_agent"`
                        SnmpCommunity string `json:"snmp_community"`
                        Ipv6Address string `json:"ipv6address"`
                        Alias string `json:"alias"`
                        TagCriticality string `json:"tag_criticality"`
                        Site string `json:"site"`
                        TagAddressFamily string `json:"tag_address_family"`
                        Contactgroups []interface{} `json:"contactgroups"`
                        NetworkScanResult struct {
                                Start interface{} `json:"start"`
                                State interface{} `json:"state"`
                                End interface{} `json:"end"`
                                Output string `json:"output"`
                        } `json:"network_scan_result"`
                        Parents []interface{} `json:"parents"`
                        Ipaddress string `json:"ipaddress"`
                        TagNetworking string `json:"tag_networking"`
                } `json:"attributes"`
                Hostname string `json:"hostname"`
                Path string `json:"path"`
        } `json:"result"`
        ResultCode int `json:"result_code"`
}

//#-------------------------------------------------------------------------------------------------------------------------------------------
type StructPutResult struct {
        Result string `json:"result"`
        ResultCode int `json:"result_code"`
}
//#-------------------------------------------------------------------------------------------------------------------------------------------
type Attributes struct {
        Alias string `json:"alias"`
        TagAgent string `json:"tag_agent"`
        TagCriticality string `json:"tag_criticality"`
        Ipaddress string `json:"ipaddress"`
}
//#-------------------------------------------------------------------------------------------------------------------------------------------
type Host struct {
        Attributes `json:"attributes"`
        Hostname string `json:"hostname"`
        Folder string `json:"folder"`
}
//#-------------------------------------------------------------------------------------------------------------------------------------------
