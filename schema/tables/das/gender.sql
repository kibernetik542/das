CREATE TABLE IF NOT EXISTS DAS.GENDER (
  ID SERIAL NOT NULL PRIMARY KEY,
  NAME VARCHAR(16) NOT NULL UNIQUE,
  ABBREVIATION CHARACTER(1) NOT NULL UNIQUE,
  DESCRIPTION TEXT,
  DATETIME_CREATED TIMESTAMP NOT NULL DEFAULT NOW(),
  DATETIME_UPDATED TIMESTAMP NOT NULL DEFAULT NOW()
);
INSERT INTO DAS.GENDER (NAME, ABBREVIATION, DESCRIPTION) VALUES ('Female','F','Biologically female');
INSERT INTO DAS.GENDER (NAME, ABBREVIATION, DESCRIPTION) VALUES ('Male','M','Biologically male');
INSERT INTO DAS.GENDER (NAME, ABBREVIATION, DESCRIPTION) VALUES ('Unknown', 'U', 'New Account' );