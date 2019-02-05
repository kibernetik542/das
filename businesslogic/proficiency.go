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

package businesslogic

import (
	"time"
)

type Proficiency struct {
	ID              int
	Name            string
	Description     string
	DivisionID      int
	CreateUserID    *int
	DateTimeCreated time.Time
	UpdateUserID    *int
	DateTImeUpdated time.Time
}

type SearchProficiencyCriteria struct {
	ProficiencyID int    `schema:"id"`
	DivisionID    int    `schema:"division"`
	Name          string `schema:"name"`
}

type IProficiencyRepository interface {
	SearchProficiency(criteria SearchProficiencyCriteria) ([]Proficiency, error)
	CreateProficiency(proficiency *Proficiency) error
	UpdateProficiency(proficiency Proficiency) error
	DeleteProficiency(proficiency Proficiency) error
}
