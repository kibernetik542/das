package partnershipdal

import (
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/Masterminds/squirrel"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
	"time"
)

var blacklistRepository = PostgresPartnershipRequestBlacklistRepository{
	Database:   nil,
	SqlBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

var blacklist = businesslogic.PartnershipRequestBlacklistEntry{
	BlockedUser: businesslogic.Account{ID: 1},
}

func TestPostgresPartnershipRequestBlacklistRepository_CreatePartnershipRequestBlacklist(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	blacklistRepository.Database = db

	rows := sqlmock.NewRows(
		[]string{
			"ID", "REPORTER_ID", "BLOCKED_USER_ID", "BLACKLIST_REASON_ID", "DETAIL", "WHITELISTED_IND",
			"CREATED_USER_ID", "DATETIME_CREATED", "UPDATE_USER_ID", "DATETIME_UPDATED",
		},
	).AddRow(
		1, 1, 2, 3, "HARASSMENT", false, 1, time.Now(), 1, time.Now(),
	).AddRow(
		1, 1, 3, 2, "SPAM", false, 1, time.Now(), 1, time.Now(),
	)

	mock.ExpectQuery(`SELECT ID, REPORTER_ID, BLOCKED_USER_ID, BLACKLIST_REASON_ID, DETAIL, 
		WHITELISTED_IND, CREATE_USER_ID, DATETIME_CREATED, UPDATE_USER_ID, DATETIME_UPDATED FROM 
		DAS.PARTNERSHIP_REQUEST_BLACKLIST`).WillReturnRows(rows)
	results, err := blacklistRepository.SearchPartnershipRequestBlacklist(businesslogic.SearchPartnershipRequestBlacklistCriteria{})

	assert.Nil(t, err, "should not throw when empty parameter is provided")
	assert.NotNil(t, results, "should at least return an empty array")
}
