package internal

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"service-automatisierte-daten-in-salesforce/pkg/config"
)

type Repository struct {
	db *gorm.DB
}

type IRepository interface {
	InsertLicenceInfo(licenceInfo LicenceInfo) (LicenceInfo, error)
	InsertLicenceProperties(licenceProperties []LicenceProperties) (LicenceProperties, error)
	GetByDID(licenceID string) string
	GetCustomerName(bydID string) string
}

func NewRepository(db *gorm.DB) Repository {
	return Repository{
		db: db,
	}
}

func InitDb(confMsql config.Mysql) (*gorm.DB, error) {
	Db, error := connectDB(confMsql)
	return Db, error
}

// This method creates the connection to the MySQL databank
// and create LicenceInfo and LicenceProperties tables if it doesn't exist
func connectDB(confMsql config.Mysql) (*gorm.DB, error) {
	dsn := confMsql.Username + ":" + confMsql.Password + "@tcp" + "(" + confMsql.Host + ":" + confMsql.Port + ")/" + confMsql.Dbname + "?" + "parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(LicenceInfo{})
	err = db.AutoMigrate(LicenceProperties{})
	if err != nil {
		log.Println("Error create table: ", err)
		return nil, err
	}

	return db, nil
}

// InsertLicenceInfo this method inserts LicenseInfo data into the LicenseInfo table
// return LicenceInfo, error
func (r Repository) InsertLicenceInfo(licenceInfo LicenceInfo) (LicenceInfo, error) {
	r.db.Create(licenceInfo)
	log.Println("licenceInfo insered")
	return LicenceInfo{}, nil
}

// InsertLicenceProperties this method inserts LicenceProperties data into the LicenceProperties table
// return LicenceProperties, error
func (r Repository) InsertLicenceProperties(licenceProperties []LicenceProperties) (LicenceProperties, error) {
	for _, prop := range licenceProperties {
		r.db.Create(prop)
	}
	return LicenceProperties{}, nil
}

// GetByDID this method searches in the table tmp_customer_license for the customer with the same licenseID
// then retrieve the customer's ByD_ID
// return bydID
func (r Repository) GetByDID(licenceID string) string {
	var bydID string
	err := r.db.Raw("SELECT ByD_ID FROM tmp_customer_license WHERE License = ?", licenceID).Scan(&bydID).Error
	if err != nil {
		return "error getBy_ID"
	}
	return bydID
}

// GetCustomerName this method searches in the table customer for the customer with the same bydID
// then retrieve the company name, which is the customer's name
// return clientName
func (r Repository) GetCustomerName(bydID string) string {
	var customerName string
	err := r.db.Raw("SELECT Company FROM customer WHERE Number = ?", bydID).Scan(&customerName).Error
	if err != nil {
		return "error GetClientName"
	}
	return customerName
}
