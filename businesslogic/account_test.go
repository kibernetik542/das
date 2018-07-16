// Dancesport Application System (DAS)
// Copyright (C) 2017, 2018 Yubing Hou
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package businesslogic_test

import (
	"errors"
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/businesslogic/reference"
	"github.com/DancesportSoftware/das/mock/businesslogic"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var testAthleteAccount = businesslogic.Account{
	FirstName:             "First Name",
	LastName:              "Last Name",
	UserGenderID:          referencebll.GENDER_MALE,
	DateOfBirth:           time.Date(2017, time.January, 1, 1, 1, 1, 1, time.UTC),
	ToSAccepted:           true,
	PrivacyPolicyAccepted: true,
	Email:      "test@test.com",
	Phone:      "1232234442",
	Signature:  "I am a parent",
	ByGuardian: true,
}

var testOrganizerAccount = businesslogic.Account{
	FirstName:             "Mighty",
	LastName:              "Meerkat",
	UserGenderID:          referencebll.GENDER_FEMALE,
	DateOfBirth:           time.Date(1997, time.May, 22, 1, 1, 1, 1, time.UTC),
	ToSAccepted:           true,
	PrivacyPolicyAccepted: true,
	Email: "mighty.meerkat@email.com",
	Phone: "3321231232",
}

func TestGetAccountByEmail(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockedAccountRepo := mock_businesslogic.NewMockIAccountRepository(mockCtrl)
	mockedAccountRepo.EXPECT().SearchAccount(businesslogic.SearchAccountCriteria{
		Email: "test@email.com",
	}).Return(nil, errors.New("should not return an account"))
	mockedAccountRepo.EXPECT().SearchAccount(businesslogic.SearchAccountCriteria{
		Email: "newuser@email.com",
	}).Return([]businesslogic.Account{
		{ID: 1, Email: "newuser@email.com"},
	}, nil)

	result := businesslogic.GetAccountByEmail("test@email.com", mockedAccountRepo)
	assert.Equal(t, 0, result.ID)
	assert.Equal(t, "", result.Email)

	result = businesslogic.GetAccountByEmail("newuser@email.com", mockedAccountRepo)
	assert.NotNil(t, result)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "newuser@email.com", result.Email)

}

func TestGetAccountByID(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockedAccountRepo := mock_businesslogic.NewMockIAccountRepository(mockCtrl)
	mockedAccountRepo.EXPECT().SearchAccount(businesslogic.SearchAccountCriteria{
		ID: 1,
	}).Return(nil, errors.New("should not return an account"))
	mockedAccountRepo.EXPECT().SearchAccount(businesslogic.SearchAccountCriteria{
		ID: 2,
	}).Return([]businesslogic.Account{
		{
			ID: 2, Email: "newuser@email.com",
		},
	}, nil)

	result := businesslogic.GetAccountByID(1, mockedAccountRepo)
	assert.Equal(t, 0, result.ID)
	assert.Equal(t, "", result.Email)

	result = businesslogic.GetAccountByID(2, mockedAccountRepo)
	assert.NotNil(t, result)
	assert.Equal(t, 2, result.ID)
	assert.Equal(t, "newuser@email.com", result.Email)
}

func TestGetAccountByUUID(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockedAccountRepo := mock_businesslogic.NewMockIAccountRepository(mockCtrl)
	mockedAccountRepo.EXPECT().SearchAccount(businesslogic.SearchAccountCriteria{
		UUID: "abc",
	}).Return(nil, errors.New("should not return an account"))
	mockedAccountRepo.EXPECT().SearchAccount(businesslogic.SearchAccountCriteria{
		UUID: "123",
	}).Return([]businesslogic.Account{
		{
			ID: 2, Email: "newuser@email.com",
		},
	}, nil)

	result := businesslogic.GetAccountByUUID("abc", mockedAccountRepo)
	assert.Equal(t, 0, result.ID)
	assert.Equal(t, "", result.Email)

	result = businesslogic.GetAccountByUUID("123", mockedAccountRepo)
	assert.NotNil(t, result)
	assert.Equal(t, 2, result.ID)
	assert.Equal(t, "newuser@email.com", result.Email)
}

func TestCreateAccountStrategy_CreateAccount(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockedAccountRepo := mock_businesslogic.NewMockIAccountRepository(mockCtrl)
	mockedAccountRepo.EXPECT().SearchAccount(businesslogic.SearchAccountCriteria{
		Email: "test@test.com",
	}).Return([]businesslogic.Account{}, errors.New("account does not exist"))
	mockedAccountRepo.EXPECT().CreateAccount(gomock.Any()).Return(nil)

	strategy := businesslogic.CreateAccountStrategy{
		AccountRepo: mockedAccountRepo,
	}

	err := strategy.CreateAccount(testAthleteAccount, "testpassword")
	assert.Nil(t, err, "should not throw an error when creating account of non-organizer")
}

func TestCreateOrganizerAccountStrategy_CreateAccount(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockedAccountRepo := mock_businesslogic.NewMockIAccountRepository(mockCtrl)
	mockedProvisionRepo := mock_businesslogic.NewMockIOrganizerProvisionRepository(mockCtrl)
	mockedHistoryRepo := mock_businesslogic.NewMockIOrganizerProvisionHistoryRepository(mockCtrl)

	mockedAccountRepo.EXPECT().SearchAccount(businesslogic.SearchAccountCriteria{
		Email: "mighty.meerkat@email.com",
	}).Return([]businesslogic.Account{}, errors.New("account does not exist"))
	mockedAccountRepo.EXPECT().CreateAccount(gomock.Any()).Return(nil)
	mockedProvisionRepo.EXPECT().CreateOrganizerProvision(gomock.Any()).Return(nil)
	mockedHistoryRepo.EXPECT().CreateOrganizerProvisionHistory(gomock.Any()).Return(nil)

	strategy := businesslogic.CreateOrganizerAccountStrategy{
		AccountRepo:   mockedAccountRepo,
		ProvisionRepo: mockedProvisionRepo,
		HistoryRepo:   mockedHistoryRepo,
	}

	err := strategy.CreateAccount(testOrganizerAccount, "testpassword")
	assert.Nil(t, err, "should create organizer account with CreateOrganizerAccountStrategy")
}
