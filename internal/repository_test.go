package internal

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

// This method test the InsertLicenceInfo method, which makes insertions on the DB
func TestInsertLicenceInfo(t *testing.T) {
	// Create an instance of LicenceInfo to mock the data to be inserted
	licenceInfo := LicenceInfo{
		ByDID_LicenceID: "01_02",
		LicenceID:       "02",
		ByD_ID:          "01",
		Name:            "test",
		SN:              "0",
		Subject:         "test",
		User:            "test",
		Users:           "test",
		Issuer:          "test",
		From:            "25.11.2023",
		To:              "25.12.2023",
		Plugin:          "1",
		Date:            "24.11.2023",
	}
	// Create an instance of the IRepository mock
	repositoryMock := NewIRepositoryMock(t)
	repositoryMock.EXPECT().InsertLicenceInfo(mock.AnythingOfType("LicenceInfo")).Run(func(l LicenceInfo) {
		assert.Equal(t, "02", l.LicenceID)
	}).Return(licenceInfo, nil).Once()
	// Call the method to be tested: "InsertLicenceInfo"
	result, err := repositoryMock.InsertLicenceInfo(licenceInfo)
	// Making assertions
	assert.Equal(t, "02", result.LicenceID)
	assert.Equal(t, "01_02", result.ByDID_LicenceID)
	assert.Equal(t, "test", result.Name)
	assert.Nil(t, err)
}
