package reference

import (
	"encoding/json"
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/controller/util"
	"github.com/DancesportSoftware/das/viewmodel"
	"net/http"
)

type StudioServer struct {
	businesslogic.IStudioRepository
}

// GET /api/reference/studio
func (server StudioServer) SearchStudioHandler(w http.ResponseWriter, r *http.Request) {
	criteria := new(businesslogic.SearchStudioCriteria)

	if parseErr := util.ParseRequestData(r, criteria); parseErr != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, util.HTTP400InvalidRequestData, parseErr.Error())
		return
	}

	studios, err := server.SearchStudio(*criteria)
	if err != nil {
		util.RespondJsonResult(w, http.StatusInternalServerError, util.HTTP500ErrorRetrievingData, err.Error())
		return
	}
	data := make([]viewmodel.Studio, 0)
	for _, each := range studios {
		data = append(data, viewmodel.StudioDataModelToViewModel(each))
	}

	output, _ := json.Marshal(data)
	w.Write(output)
}

// POST /api/reference/studio
func (server StudioServer) CreateStudioHandler(w http.ResponseWriter, r *http.Request) {
	util.RespondJsonResult(w, http.StatusNotImplemented, "not implemented", nil)
}

// PUT /api/reference/studio
func (server StudioServer) UpdateStudioHandler(w http.ResponseWriter, r *http.Request) {
	util.RespondJsonResult(w, http.StatusNotImplemented, "not implemented", nil)
}

// DELETE /api/reference/studio
func (server StudioServer) DeleteStudioHandler(w http.ResponseWriter, r *http.Request) {
	util.RespondJsonResult(w, http.StatusNotImplemented, "not implemented", nil)
}
