package internal

import "fmt"

const (
	Invoice                  = 1 << iota // same as (1 << 0) or 0b00000001
	HenrichsenAktenplan                  // 2
	ADS                                  // 4
	Contract                             // 8
	OutlookWeb                           // 16
	OutlookSync                          // 32
	OfficeML                             // 64
	MobileClient                         // 128
	ContractMandantenModul               // 256
	ContractControllingModul             // 512
	ReportingModul                       // 1024
	HRRecords                            // 2048
	ContractSAPKonnektor                 // 4096
	Gadgets                              // 8192
	Inbox                                // 16384
	Unknown1                             // 32768
	Unknown2                             // 65536
	Unknown3                             // 131072
	Unknown4                             // 262144
	Unknown5                             // 524288
	CONTRACT4A                           // 1048576
	Unknown6                             // 2097152
	CONTRACT4B                           // 4194304
)

var PluginMap = map[int]string{
	Invoice:                  "Invoice",
	HenrichsenAktenplan:      "Henrichsen Aktenplan / Keine allgemeine Info !!!",
	ADS:                      "ADS",
	Contract:                 "Contract",
	OutlookWeb:               "Outlook Web",
	OutlookSync:              "Outlook sync",
	OfficeML:                 "OfficeML",
	MobileClient:             "MobileClient",
	ContractMandantenModul:   "Contract Mandanten Modul",
	ContractControllingModul: "Contract Controlling Modul",
	ReportingModul:           "Reporting-Modul",
	HRRecords:                "HR Records",
	ContractSAPKonnektor:     "Contract SAP Konnektor ( nur für Otris interessant )",
	Gadgets:                  "Gadgets",
	Inbox:                    "Inbox",
	Unknown1:                 "Unknown1",
	Unknown2:                 "Unknown2",
	Unknown3:                 "Unknown3",
	Unknown4:                 "Unknown4",
	Unknown5:                 "Unknown5",
	CONTRACT4A:               "CONTRACT 4 Bestandteil A E-AKTE",
	Unknown6:                 "Unknown6",
	CONTRACT4B:               "CONTRACT 4 Bestandteil B",
}

// mapPluginName this method evaluates the value given on plugin and returns its given meaning
// return string, error
func mapPluginName(flags int) (string, error) {
	for bit, name := range PluginMap {
		if flags&bit != 0 {
			return name, nil
		}
	}
	return "", fmt.Errorf("is not a valid plugin")
}
