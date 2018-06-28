// Copyright 2017, 2018 Yubing Hou. All rights reserved.
// Use of this source code is governed by GPL license
// that can be found in the LICENSE file

package account

import (
	"encoding/json"
	"github.com/DancesportSoftware/das/businesslogic/reference"
	mock_reference "github.com/DancesportSoftware/das/mock/businesslogic/reference"
	"github.com/DancesportSoftware/das/viewmodel"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAccountGenderHandler_GetAccountGenderHandler(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockedGenderRepo := mock_reference.NewMockIGenderRepository(mockCtrl)

	server := GenderServer{
		IGenderRepository: mockedGenderRepo,
	}

	req, _ := http.NewRequest(http.MethodGet, "/api/account/gender", nil)
	w := httptest.NewRecorder()

	// test with zero param
	mockedGenderRepo.EXPECT().GetAllGenders().Return([]referencebll.Gender{
		referencebll.Gender{ID: 1, Name: "Female"},
		referencebll.Gender{ID: 2, Name: "Male"},
	}, nil)
	server.GetAccountGenderHandler(w, req)
	genders := make([]viewmodel.Gender, 0)
	err := json.Unmarshal([]byte(w.Body.String()), &genders)
	assert.Nil(t, err, "should return a list of genders")
	w.Flush()

	// test with bad param should not result in error
	query := req.URL.Query()
	query.Add("badparam", "indeed")
	req.URL.RawQuery = query.Encode()
	// log.Printf("GET %s\n", req.URL.String())
	mockedGenderRepo.EXPECT().GetAllGenders().Return([]referencebll.Gender{
		referencebll.Gender{ID: 1, Name: "Female"},
		referencebll.Gender{ID: 2, Name: "Male"},
	}, nil)
	server.GetAccountGenderHandler(w, req)
	// log.Println(w.Body)
	assert.EqualValues(t, http.StatusOK, w.Code, "should not receive HTTP 400 when sending a bad request")
}
