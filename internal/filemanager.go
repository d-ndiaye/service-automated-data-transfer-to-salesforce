package internal

import (
	"github.com/google/uuid"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var propertiesReg = regexp.MustCompile("Properties: (.+)")
var snReg = regexp.MustCompile("SN: (.+)")
var subjectReg = regexp.MustCompile("Subject: (.+)")
var userReg = regexp.MustCompile("User: (.+)")
var IssuerReg = regexp.MustCompile("Issuer: (.+)")
var fromReg = regexp.MustCompile("From: (.+)")
var usersReg = regexp.MustCompile("Users: (.+)")
var toReg = regexp.MustCompile("To: (.+)")
var pluginReg = regexp.MustCompile("plugin: (\\d+)")

type FileManager struct {
	repository Repository
}

func NewFileManager(repository Repository) FileManager {
	return FileManager{
		repository: repository,
	}
}

// ReadInfosPemFile this method reads a .pem file
// return LicenceInfo, []LicenceProperties, error
func (fm FileManager) ReadInfosPemFile(folderPath string, licenceID string) (LicenceInfo, []LicenceProperties, error) {
	log.Println("Reading .the pem file :", folderPath)
	content, err := os.ReadFile(folderPath)
	if err != nil {
		log.Println("Error reading .pem file :", err)
		return LicenceInfo{}, []LicenceProperties{}, err
	}
	contentString := string(content)
	extractedInfo, extractedLicenceProperties, _ := fm.extractInfo(contentString, licenceID)
	return extractedInfo, extractedLicenceProperties, nil
}

// ExtractInfo this method extracts information from the .pem file
// return LicenceInfo, []LicenceProperties, error
func (fm FileManager) extractInfo(contentString string, licenceID string) (LicenceInfo, []LicenceProperties, error) {
	extractedProperties := propertiesReg.FindStringSubmatch(contentString)[1]
	extractedLicenceProperties := fm.extractKeyAndValue(extractedProperties, licenceID)
	extractedSN := snReg.FindStringSubmatch(contentString)[1]
	extractedSubject := subjectReg.FindStringSubmatch(contentString)[1]
	extractedUsers := usersReg.FindStringSubmatch(contentString)[1]
	extractedIssuer := IssuerReg.FindStringSubmatch(contentString)[1]
	extractedUser := userReg.FindStringSubmatch(contentString)[1]
	extractedFrom := fromReg.FindStringSubmatch(contentString)[1]
	extractedTo := toReg.FindStringSubmatch(contentString)[1]
	extractedPlugin := pluginReg.FindStringSubmatch(contentString)[1]

	i, _ := strconv.Atoi(extractedPlugin)
	pluginName, _ := mapPluginName(i)
	log.Println("map PluginName", pluginName)

	timeNow := time.Now()
	timeNowFormatted := timeNow.Format("02.01.2006")
	licenceID = fm.ExtractFileName(licenceID)
	byD_ID := fm.repository.GetByDID(licenceID)
	byDID_LicenceID := byD_ID + "_" + licenceID
	name := fm.repository.GetCustomerName(byD_ID)

	extractedInfo := LicenceInfo{
		ByDID_LicenceID: byDID_LicenceID,
		LicenceID:       licenceID,
		ByD_ID:          byD_ID,
		Name:            name,
		SN:              extractedSN,
		Subject:         extractedSubject,
		User:            extractedUser,
		Users:           extractedUsers,
		Issuer:          extractedIssuer,
		From:            extractedFrom,
		To:              extractedTo,
		Plugin:          pluginName,
		Date:            timeNowFormatted,
	}
	return extractedInfo, extractedLicenceProperties, nil
}

// extractKeyAndValue this method extracts the key and value from the LicenseProperties list and
// prepares the list of []LicenseProperties to be returned( id, ByDID_LicenceID, LicenceID, key, value, date)
// return []LicenceProperties
func (fm FileManager) extractKeyAndValue(extractedProperties string, licenceID string) []LicenceProperties {
	licenceProperties := []LicenceProperties{}
	properties := strings.Split(extractedProperties, "; ")
	timeNow := time.Now()
	for _, prop := range properties {
		parts := strings.Split(prop, ": ")
		licenceID := fm.ExtractFileName(licenceID)
		byD_ID := fm.repository.GetByDID(licenceID)
		byDID_LicenceID := byD_ID + "_" + licenceID
		timeNowFormatted := timeNow.Format("02.01.2006")
		id := uuid.New().String()
		if len(parts) == 2 {
			licenceProperties = append(licenceProperties, LicenceProperties{
				ID:              id,
				ByDID_LicenceID: byDID_LicenceID,
				LicenceID:       licenceID,
				Key:             parts[0],
				Value:           parts[1],
				Date:            timeNowFormatted,
			})
		}
	}
	return licenceProperties
}

// ExtractFileName this method extracts the file name without endings .pem
// return filename
func (fm FileManager) ExtractFileName(folderPath string) string {
	fileNameMitEndungen := filepath.Base(folderPath)
	fileName := strings.TrimSuffix(fileNameMitEndungen, filepath.Ext(fileNameMitEndungen))
	return fileName
}
