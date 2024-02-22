package internal

type LicenceProperties struct {
	ID              string
	ByDID_LicenceID string
	LicenceID       string
	Key             string
	Value           string
	Date            string
}

type LicenceInfo struct {
	ByDID_LicenceID string
	LicenceID       string
	ByD_ID          string
	Name            string
	SN              string
	Subject         string
	User            string
	Users           string
	Issuer          string
	From            string
	To              string
	Date            string
	Plugin          string
}

// TableName this method returns the table name of my datenbank "LicenceInfo" for the LicenceInfo structure.
func (LicenceInfo) TableName() string {
	return "LicenceInfo"
}

// TableName this method returns the table name of my datenbank "LicenceProperties" for the LicenceProperties structure.
func (LicenceProperties) TableName() string {
	return "LicenceProperties"
}
