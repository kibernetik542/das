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

package account

import (
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/controller/util"
	"github.com/DancesportSoftware/das/controller/util/authentication"
	"github.com/DancesportSoftware/das/viewmodel"
	"log"
	"net/http"
)

// AccountServer provides a virtual server that handles requests that are related to Account
type AccountServer struct {
	businesslogic.IAccountRepository
	businesslogic.IAccountRoleRepository
	businesslogic.IOrganizerProvisionRepository
	businesslogic.IOrganizerProvisionHistoryRepository
	businesslogic.IUserPreferenceRepository
}

// RegisterAccountHandler handle the request
// 	POST /api/account/register
// Accepted JSON parameters:
//	{
//		"email": "awesomeuser@email.com",
//		"phone": 1234567890,
//		"password": !@#$1234,
//		"firstname": "Awesome",
//		"lastname": "User"
//	}
func (server AccountServer) RegisterAccountHandler(w http.ResponseWriter, r *http.Request) {
	createAccount := new(viewmodel.CreateAccount)

	if err := util.ParseRequestBodyData(r, createAccount); err != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if err := createAccount.Validate(); err != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	account := createAccount.ToAccountModel()

	strategy := businesslogic.CreateAccountStrategy{
		AccountRepo:          server.IAccountRepository,
		RoleRepository:       server.IAccountRoleRepository,
		PreferenceRepository: server.IUserPreferenceRepository,
	}

	if err := strategy.CreateAccount(account, createAccount.Password); err != nil {
		util.RespondJsonResult(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	util.RespondJsonResult(w, http.StatusOK, "success", nil)
}

// AccountAuthenticationHandler handles the request:
// 	POST /api/v1.0account/authenticate
// Accepted JSON parameters:
// 	{
//		"username": "user@email.com",
//		"password": "password"
//	}
// Sample returned response:
//	{
//		"status": 200,
//		"message": "authorized",
//		"data": {
//			"token": "some.jwt.token",
//		}
//	}
func (server AccountServer) AccountAuthenticationHandler(w http.ResponseWriter, r *http.Request) {
	loginDTO := new(viewmodel.Login)
	err := util.ParseRequestBodyData(r, loginDTO)
	if loginDTO.Email == "" {
		util.RespondJsonResult(w, http.StatusBadRequest, "Email is required", nil)
		return
	}
	if loginDTO.Password == "" {
		util.RespondJsonResult(w, http.StatusBadRequest, "Password is required", nil)
		return
	}
	if err != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, "invalid credential", nil)
		return
	} else if len(loginDTO.Email) < 4 || len(loginDTO.Password) < 8 {
		util.RespondJsonResult(w, http.StatusBadRequest, "invalid credential", nil)
		return
	}

	err = businesslogic.AuthenticateUser(loginDTO.Email, loginDTO.Password, server.IAccountRepository)

	if err != nil {
		util.RespondJsonResult(w, http.StatusUnauthorized, err.Error(), nil)
		return
	}
	account := businesslogic.GetAccountByEmail(loginDTO.Email, server.IAccountRepository)

	// user jwt authentication
	authString := authentication.GenerateAuthenticationToken(account)
	if err != nil {
		log.Printf("[error] generating client credential: %s\n", err.Error())
		util.RespondJsonResult(w, http.StatusUnauthorized, "error in generating client credential", nil)
		return
	}
	response := struct {
		Token string `json:"token"`
	}{Token: authString}
	util.RespondJsonResult(w, http.StatusOK, "authorized", response)
	return
}
