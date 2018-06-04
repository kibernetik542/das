package provision

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/dataaccess/common"
	"github.com/Masterminds/squirrel"
	"log"
)

const (
	DAS_ORGANIZER_PROVISION                  = "DAS.ORGANIZER_PROVISION"
	DAS_ORGANIZER_PROVISION_COL_ORGANIZER_ID = "ORGANIZER_ID"
	DAS_ORGANIZER_PROVISION_COL_HOSTED       = "HOSTED"
	DAS_ORGANIZER_PROVISION_COL_AVAILABLE    = "AVAILABLE"
)

type PostgresOrganizerProvisionRepository struct {
	Database   *sql.DB
	SqlBuilder squirrel.StatementBuilderType
}

func (repo PostgresOrganizerProvisionRepository) CreateOrganizerProvision(provision *businesslogic.OrganizerProvision) error {
	if repo.Database == nil {
		return errors.New("data source of PostgresOrganizerProvisionRepository is not specified")
	}
	stmt := repo.SqlBuilder.Insert("").
		Into(DAS_ORGANIZER_PROVISION).
		Columns(
			DAS_ORGANIZER_PROVISION_COL_ORGANIZER_ID,
			DAS_ORGANIZER_PROVISION_COL_HOSTED,
			DAS_ORGANIZER_PROVISION_COL_AVAILABLE,
			common.COL_CREATE_USER_ID,
			common.COL_DATETIME_CREATED,
			common.COL_UPDATE_USER_ID,
			common.COL_DATETIME_UPDATED,
		).Values(provision.OrganizerID, provision.Hosted, provision.Available, provision.CreateUserID, provision.DateTimeCreated, provision.UpdateUserID, provision.DateTimeUpdated)
	_, err := stmt.RunWith(repo.Database).Exec()
	if err != nil {
		log.Printf("[error] initializing organizer organizer: %s\n", err.Error())
		return err
	}

	//CreateOrganizerProvisionHistoryEntry(accountID, 0, "initial organizer", accountID)
	if err != nil {
		log.Printf("[error] initializing organizer organizer history: %s\n", err.Error())
		return err
	}
	return err
}

func (repo PostgresOrganizerProvisionRepository) UpdateOrganizerProvision(provision businesslogic.OrganizerProvision) error {
	if repo.Database == nil {
		return errors.New("data source of PostgresOrganizerProvisionRepository is not specified")
	}
	stmt := repo.SqlBuilder.Update("").
		Table(DAS_ORGANIZER_PROVISION).
		Set(DAS_ORGANIZER_PROVISION_COL_AVAILABLE, provision.Available).
		Set(DAS_ORGANIZER_PROVISION_COL_HOSTED, provision.Hosted).
		Set(common.COL_DATETIME_UPDATED, provision.DateTimeUpdated).
		Where(squirrel.Eq{DAS_ORGANIZER_PROVISION_COL_ORGANIZER_ID: provision.OrganizerID})
	_, err := stmt.RunWith(repo.Database).Exec()
	return err
}

func (repo PostgresOrganizerProvisionRepository) SearchOrganizerProvision(
	criteria businesslogic.SearchOrganizerProvisionCriteria) ([]businesslogic.OrganizerProvision, error) {
	if repo.Database == nil {
		return nil, errors.New("data source of PostgresOrganizerProvisionRepository is not specified")
	}

	stmt := repo.SqlBuilder.Select(fmt.Sprintf("%s, %s, %s, %s, %s, %s, %s, %s",
		common.PRIMARY_KEY,
		DAS_ORGANIZER_PROVISION_COL_ORGANIZER_ID,
		DAS_ORGANIZER_PROVISION_COL_HOSTED,
		DAS_ORGANIZER_PROVISION_COL_AVAILABLE,
		common.COL_CREATE_USER_ID,
		common.COL_DATETIME_CREATED,
		common.COL_UPDATE_USER_ID,
		common.COL_DATETIME_UPDATED)).
		From(DAS_ORGANIZER_PROVISION).Where(squirrel.Eq{DAS_ORGANIZER_PROVISION_COL_ORGANIZER_ID: criteria.OrganizerID})

	rows, err := stmt.RunWith(repo.Database).Query()

	provisions := make([]businesslogic.OrganizerProvision, 0)
	for rows.Next() {
		each := businesslogic.OrganizerProvision{}
		rows.Scan(
			&each.ID,
			&each.OrganizerID,
			&each.Hosted,
			&each.Available,
			&each.CreateUserID,
			&each.DateTimeCreated,
			&each.UpdateUserID,
			&each.DateTimeUpdated,
		)
		provisions = append(provisions, each)
	}

	return provisions, err
}

func (repo PostgresOrganizerProvisionRepository) DeleteOrganizerProvision(provision businesslogic.OrganizerProvision) error {
	if repo.Database == nil {
		return errors.New("data source of PostgresOrganizerProvisionRepository is not specified")
	}
	return errors.New("not implemented")
}