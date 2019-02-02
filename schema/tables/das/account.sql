-- create das.user_account table
CREATE TABLE IF NOT EXISTS DAS.ACCOUNT(
  ID SERIAL NOT NULL PRIMARY KEY,
  UID TEXT NOT NULL UNIQUE,
  ACCOUNT_STATUS_ID INTEGER NOT NULL REFERENCES DAS.ACCOUNT_STATUS(ID),
  USER_GENDER_ID INTEGER NOT NULL REFERENCES DAS.GENDER(ID),
  LAST_NAME VARCHAR(64) NOT NULL,
  MIDDLE_NAMES TEXT,
  FIRST_NAME VARCHAR(64) NOT NULL,
  DATE_OF_BIRTH DATE,
  EMAIL TEXT NOT NULL UNIQUE,
  PHONE TEXT NOT NULL,
  DATETIME_CREATED TIMESTAMP NOT NULL DEFAULT NOW(),
  DATETIME_UPDATED TIMESTAMP NOT NULL DEFAULT NOW(),
  TOS_ACCEPTED BOOLEAN NOT NULL DEFAULT FALSE,
  PP_ACCEPTED BOOLEAN NOT NULL DEFAULT FALSE,
  BY_GUARDIAN BOOLEAN NOT NULL DEFAULT FALSE,
  GUARDIAN_SIGNATURE TEXT
);

-- create indexes for faster search
CREATE INDEX ON DAS.ACCOUNT (LAST_NAME);
CREATE INDEX ON DAS.ACCOUNT (FIRST_NAME);
CREATE INDEX ON DAS.ACCOUNT (EMAIL);
CREATE INDEX ON DAS.ACCOUNT (PHONE);
CREATE INDEX ON DAS.ACCOUNT (ACCOUNT_STATUS_ID);
CREATE INDEX ON DAS.ACCOUNT (DATE_OF_BIRTH);

-- create audit
CREATE TABLE IF NOT EXISTS DAS.ACCOUNT_L (
  ID SERIAL NOT NULL PRIMARY KEY,
  ACCOUNT_ID INTEGER NOT NULL,
  ACCOUNT_STATUS INTEGER,
  PHONE TEXT,
  EMAIL TEXT,
  PASSWORD_HASH BYTEA,
  TOS_ACCEPTED BOOLEAN,
  PP_ACCEPTED BOOLEAN,
  CHANGED_ON TIMESTAMP NOT NULL
);

-- create account phone change trigger
CREATE OR REPLACE FUNCTION LOG_ACCOUNT_PHONE_CHANGES ()
  RETURNS TRIGGER AS
$BODY$
  BEGIN
    IF NEW.PHONE != OLD.PHONE THEN
      INSERT INTO DAS.ACCOUNT_L(ACCOUNT_ID, PHONE, CHANGED_ON)
        VALUES (OLD.ID, OLD.PHONE, NOW());
    END IF;
    RETURN NEW;

  END
$BODY$
  LANGUAGE plpgsql VOLATILE COST 100;

CREATE TRIGGER ACCOUNT_PHONE_CHANGES
  BEFORE UPDATE
  ON DAS.ACCOUNT
  FOR EACH ROW
  EXECUTE PROCEDURE LOG_ACCOUNT_PHONE_CHANGES();