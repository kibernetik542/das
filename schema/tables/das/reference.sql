-- Create table das.country
-- Indexed Columns: COUNTRY_ID, COUNTRY_NAME, COUNTRY_ABBR
CREATE TABLE IF NOT EXISTS DAS.COUNTRY (
   ID SERIAL NOT NULL PRIMARY KEY,
   NAME TEXT NOT NULL UNIQUE, -- indexed
   ABBREVIATION VARCHAR(8) NOT NULL UNIQUE, -- indexed
   CREATE_USER_ID INTEGER REFERENCES DAS.ACCOUNT (ID),
   DATETIME_CREATED TIMESTAMP NOT NULL DEFAULT NOW(),
   UPDATE_USER_ID INTEGER REFERENCES DAS.ACCOUNT (ID),
   DATETIME_UPDATED TIMESTAMP NOT NULL DEFAULT NOW()
);


INSERT INTO DAS.COUNTRY(ABBREVIATION, NAME) VALUES('USA','United States');
INSERT INTO DAS.COUNTRY(ABBREVIATION, NAME) VALUES('CAN','Canada');
INSERT INTO DAS.COUNTRY(ABBREVIATION, NAME) VALUES('ALB','Albania');
INSERT INTO DAS.COUNTRY(ABBREVIATION, NAME) VALUES('AND','Andorra');
INSERT INTO DAS.COUNTRY(ABBREVIATION, NAME) VALUES('ARG','Argentina');
INSERT INTO DAS.COUNTRY(ABBREVIATION, NAME) VALUES('ARM','Armenia');
INSERT INTO DAS.COUNTRY(ABBREVIATION, NAME) VALUES('AUS','Australia');
INSERT INTO DAS.COUNTRY(ABBREVIATION, NAME) VALUES('AUT','Austria');
INSERT INTO DAS.COUNTRY(ABBREVIATION, NAME) VALUES('AZE','Azerbaijan');
INSERT INTO DAS.COUNTRY(ABBREVIATION, NAME) VALUES('BAH','Bahamas');
INSERT INTO DAS.COUNTRY(ABBREVIATION, NAME) VALUES('BAR','Barbados');
INSERT INTO DAS.COUNTRY(ABBREVIATION, NAME) VALUES('BLR','Belarus');
INSERT INTO DAS.COUNTRY(ABBREVIATION, NAME) VALUES('BEL','Belgium');
INSERT INTO DAS.COUNTRY(ABBREVIATION, NAME) VALUES('BIZ','Belize');
INSERT INTO DAS.COUNTRY(ABBREVIATION, NAME) VALUES('BIH','Bosnia and Herzegovina');
INSERT INTO DAS.COUNTRY(ABBREVIATION, NAME) VALUES('BRA','Brazil');
INSERT INTO DAS.COUNTRY(ABBREVIATION, NAME) VALUES('BUL','Bulgaria');
INSERT INTO DAS.COUNTRY(ABBREVIATION, NAME) VALUES('CHI','Chile');
INSERT INTO DAS.COUNTRY(ABBREVIATION, NAME) VALUES('CRO','Croatia');
INSERT INTO DAS.COUNTRY(ABBREVIATION, NAME) VALUES('CYP','Cyprus');
INSERT INTO DAS.COUNTRY(ABBREVIATION, NAME) VALUES('CZE','Czech Republic');
INSERT INTO DAS.COUNTRY(ABBREVIATION, NAME) VALUES('DEN','Denmark');
INSERT INTO DAS.COUNTRY(ABBREVIATION, NAME) VALUES('EST','Estonia');
INSERT INTO DAS.COUNTRY(ABBREVIATION, NAME) VALUES('FIN','Finland');
INSERT INTO DAS.COUNTRY(ABBREVIATION, NAME) VALUES('FRA','France');
INSERT INTO DAS.COUNTRY(ABBREVIATION, NAME) VALUES('GEO','Georgia');
INSERT INTO DAS.COUNTRY(ABBREVIATION, NAME) VALUES('GER','Germany');
INSERT INTO DAS.COUNTRY(ABBREVIATION, NAME) VALUES('GBR','Great Britain');
INSERT INTO DAS.COUNTRY(ABBREVIATION, NAME) VALUES('GRE','Greece');
INSERT INTO DAS.COUNTRY(ABBREVIATION, NAME) VALUES('HUN','Hungary');
INSERT INTO DAS.COUNTRY(ABBREVIATION, NAME) VALUES('ISL','Iceland');
INSERT INTO DAS.COUNTRY(ABBREVIATION, NAME) VALUES('IND','India');
INSERT INTO DAS.COUNTRY(ABBREVIATION, NAME) VALUES('IRL','Ireland');
INSERT INTO DAS.COUNTRY(ABBREVIATION, NAME) VALUES('ISR','Israel');
INSERT INTO DAS.COUNTRY(ABBREVIATION, NAME) VALUES('ITA','Italy');
INSERT INTO DAS.COUNTRY(ABBREVIATION, NAME) VALUES('JPN','Japan');
INSERT INTO DAS.COUNTRY(ABBREVIATION, NAME) VALUES('LAT','Latvia');
INSERT INTO DAS.COUNTRY(ABBREVIATION, NAME) VALUES('LIE','Liechtenstein');
INSERT INTO DAS.COUNTRY(ABBREVIATION, NAME) VALUES('LTU','Lithuania');
INSERT INTO DAS.COUNTRY(ABBREVIATION, NAME) VALUES('LUX','Luxembourg');
INSERT INTO DAS.COUNTRY(ABBREVIATION, NAME) VALUES('MKD','Macedonia');
INSERT INTO DAS.COUNTRY(ABBREVIATION, NAME) VALUES('MAS','Malaysia');
INSERT INTO DAS.COUNTRY(ABBREVIATION, NAME) VALUES('MLT','Malta');
INSERT INTO DAS.COUNTRY(ABBREVIATION, NAME) VALUES('MEX','Mexico');
INSERT INTO DAS.COUNTRY(ABBREVIATION, NAME) VALUES('MDA','Moldova');
INSERT INTO DAS.COUNTRY(ABBREVIATION, NAME) VALUES('MON','Monaco');
INSERT INTO DAS.COUNTRY(ABBREVIATION, NAME) VALUES('MNE','Montenegro');
INSERT INTO DAS.COUNTRY(ABBREVIATION, NAME) VALUES('NED','Netherlands');
INSERT INTO DAS.COUNTRY(ABBREVIATION, NAME) VALUES('NOR','Norway');
INSERT INTO DAS.COUNTRY(ABBREVIATION, NAME) VALUES('NZL','New Zealand');
INSERT INTO DAS.COUNTRY(ABBREVIATION, NAME) VALUES('POL','Poland');
INSERT INTO DAS.COUNTRY(ABBREVIATION, NAME) VALUES('POR','Portugal');
INSERT INTO DAS.COUNTRY(ABBREVIATION, NAME) VALUES('ROU','Romania');
INSERT INTO DAS.COUNTRY(ABBREVIATION, NAME) VALUES('RUS','Russia');
INSERT INTO DAS.COUNTRY(ABBREVIATION, NAME) VALUES('SRB','Serbia');
INSERT INTO DAS.COUNTRY(ABBREVIATION, NAME) VALUES('SIN','Singapore');
INSERT INTO DAS.COUNTRY(ABBREVIATION, NAME) VALUES('SVK','Slovakia');
INSERT INTO DAS.COUNTRY(ABBREVIATION, NAME) VALUES('SLO','Slovenia');
INSERT INTO DAS.COUNTRY(ABBREVIATION, NAME) VALUES('RSA','South Africa');
INSERT INTO DAS.COUNTRY(ABBREVIATION, NAME) VALUES('KOR','South Korea');
INSERT INTO DAS.COUNTRY(ABBREVIATION, NAME) VALUES('ESP','Spain');
INSERT INTO DAS.COUNTRY(ABBREVIATION, NAME) VALUES('SWE','Sweden');
INSERT INTO DAS.COUNTRY(ABBREVIATION, NAME) VALUES('SUI','Switzerland');
INSERT INTO DAS.COUNTRY(ABBREVIATION, NAME) VALUES('TWN','Taiwan');
INSERT INTO DAS.COUNTRY(ABBREVIATION, NAME) VALUES('TUR','Turkey');
INSERT INTO DAS.COUNTRY(ABBREVIATION, NAME) VALUES('UKR','Ukraine');
INSERT INTO DAS.COUNTRY(ABBREVIATION, NAME) VALUES('VIE','Vietnam');

-- Get_Country_By_Id returns a table of country where ID is cid
-- To invoke this method in CLI or Go, run query `SELECT * FROM GET_COUNTRY_BY_ID (1);`
CREATE OR REPLACE FUNCTION get_country_by_id (cid INTEGER)
  RETURNS TABLE (
                  ID INTEGER,
                  NAME TEXT,
                  ABBREVIATION VARCHAR(8),
                  CREATE_USER_ID INTEGER,
                  DATETIME_CREATED TIMESTAMP,
                  UPDATE_USER_ID INTEGER,
                  DATETIME_UPDATE TIMESTAMP
                )
AS
$$
BEGIN
  RETURN QUERY SELECT
                 DAS.COUNTRY.ID,
                 DAS.COUNTRY.NAME,
                 DAS.COUNTRY.ABBREVIATION,
                 DAS.COUNTRY.CREATE_USER_ID,
                 DAS.COUNTRY.DATETIME_CREATED,
                 DAS.COUNTRY.UPDATE_USER_ID,
                 DAS.COUNTRY.DATETIME_UPDATED
               FROM DAS.COUNTRY
               WHERE DAS.COUNTRY.ID = cid;
END;
$$
  LANGUAGE PLPGSQL;

-- Create table das.state
-- Indexed Columns:
CREATE TABLE IF NOT EXISTS DAS.STATE (
                                       ID SERIAL NOT NULL PRIMARY KEY,
                                       NAME VARCHAR(32) NOT NULL,
                                       ABBREVIATION VARCHAR (8),
                                       COUNTRY_ID INTEGER NOT NULL REFERENCES DAS.COUNTRY (ID),
                                       CREATE_USER_ID INTEGER REFERENCES DAS.ACCOUNT (ID),
                                       DATETIME_CREATED TIMESTAMP NOT NULL DEFAULT NOW(),
                                       UPDATE_USER_ID INTEGER REFERENCES DAS.ACCOUNT (ID),
                                       DATETIME_UPDATED TIMESTAMP NOT NULL DEFAULT NOW(),
                                       UNIQUE (NAME, COUNTRY_ID)
);

CREATE INDEX ON DAS.STATE (COUNTRY_ID);
CREATE INDEX ON DAS.STATE (NAME);

INSERT INTO DAS.STATE (NAME, ABBREVIATION, COUNTRY_ID) VALUES('Alabama', 'AL', (SELECT C.ID FROM DAS.COUNTRY C WHERE C.ABBREVIATION = 'USA'));
INSERT INTO DAS.STATE (NAME, ABBREVIATION, COUNTRY_ID) VALUES('Alaska', 'AK', (SELECT C.ID FROM DAS.COUNTRY C WHERE C.ABBREVIATION = 'USA'));
INSERT INTO DAS.STATE (NAME, ABBREVIATION, COUNTRY_ID) VALUES('Arizona', 'AZ', (SELECT C.ID FROM DAS.COUNTRY C WHERE C.ABBREVIATION = 'USA'));
INSERT INTO DAS.STATE (NAME, ABBREVIATION, COUNTRY_ID) VALUES('Arkansas', 'AR', (SELECT C.ID FROM DAS.COUNTRY C WHERE C.ABBREVIATION = 'USA'));
INSERT INTO DAS.STATE (NAME, ABBREVIATION, COUNTRY_ID) VALUES('California', 'CA', (SELECT C.ID FROM DAS.COUNTRY C WHERE C.ABBREVIATION = 'USA'));
INSERT INTO DAS.STATE (NAME, ABBREVIATION, COUNTRY_ID) VALUES('Colorado', 'CO', (SELECT C.ID FROM DAS.COUNTRY C WHERE C.ABBREVIATION = 'USA'));
INSERT INTO DAS.STATE (NAME, ABBREVIATION, COUNTRY_ID) VALUES('Connecticut', 'CT', (SELECT C.ID FROM DAS.COUNTRY C WHERE C.ABBREVIATION = 'USA'));
INSERT INTO DAS.STATE (NAME, ABBREVIATION, COUNTRY_ID) VALUES('District of Columbia','DC',(SELECT C.ID FROM DAS.COUNTRY C WHERE C.ABBREVIATION = 'USA'));
INSERT INTO DAS.STATE (NAME, ABBREVIATION, COUNTRY_ID) VALUES('Delaware','DE',(SELECT C.ID FROM DAS.COUNTRY C WHERE C.ABBREVIATION = 'USA'));
INSERT INTO DAS.STATE (NAME, ABBREVIATION, COUNTRY_ID) VALUES('Florida','FL',(SELECT C.ID FROM DAS.COUNTRY C WHERE C.ABBREVIATION = 'USA'));
INSERT INTO DAS.STATE (NAME, ABBREVIATION, COUNTRY_ID) VALUES('Georgia','GA',(SELECT C.ID FROM DAS.COUNTRY C WHERE C.ABBREVIATION = 'USA'));
INSERT INTO DAS.STATE (NAME, ABBREVIATION, COUNTRY_ID) VALUES('Hawaii','HI',(SELECT C.ID FROM DAS.COUNTRY C WHERE C.ABBREVIATION = 'USA'));
INSERT INTO DAS.STATE (NAME, ABBREVIATION, COUNTRY_ID) VALUES('Idaho','ID',(SELECT C.ID FROM DAS.COUNTRY C WHERE C.ABBREVIATION = 'USA'));
INSERT INTO DAS.STATE (NAME, ABBREVIATION, COUNTRY_ID) VALUES('Illinois','IL',(SELECT C.ID FROM DAS.COUNTRY C WHERE C.ABBREVIATION = 'USA'));
INSERT INTO DAS.STATE (NAME, ABBREVIATION, COUNTRY_ID) VALUES('Indiana','IN',(SELECT C.ID FROM DAS.COUNTRY C WHERE C.ABBREVIATION = 'USA'));
INSERT INTO DAS.STATE (NAME, ABBREVIATION, COUNTRY_ID) VALUES('Iowa','IA',(SELECT C.ID FROM DAS.COUNTRY C WHERE C.ABBREVIATION = 'USA'));
INSERT INTO DAS.STATE (NAME, ABBREVIATION, COUNTRY_ID) VALUES('Kansas','KS',(SELECT C.ID FROM DAS.COUNTRY C WHERE C.ABBREVIATION = 'USA'));
INSERT INTO DAS.STATE (NAME, ABBREVIATION, COUNTRY_ID) VALUES('Kentucky','KY',(SELECT C.ID FROM DAS.COUNTRY C WHERE C.ABBREVIATION = 'USA'));
INSERT INTO DAS.STATE (NAME, ABBREVIATION, COUNTRY_ID) VALUES('Louisiana','LA',(SELECT C.ID FROM DAS.COUNTRY C WHERE C.ABBREVIATION = 'USA'));
INSERT INTO DAS.STATE (NAME, ABBREVIATION, COUNTRY_ID) VALUES('Maine','ME', (SELECT C.ID FROM DAS.COUNTRY C WHERE C.ABBREVIATION = 'USA'));
INSERT INTO DAS.STATE (NAME, ABBREVIATION, COUNTRY_ID) VALUES('Maryland','MD',(SELECT C.ID FROM DAS.COUNTRY C WHERE C.ABBREVIATION = 'USA'));
INSERT INTO DAS.STATE (NAME, ABBREVIATION, COUNTRY_ID) VALUES('Massachusetts','MA',(SELECT C.ID FROM DAS.COUNTRY C WHERE C.ABBREVIATION = 'USA'));
INSERT INTO DAS.STATE (NAME, ABBREVIATION, COUNTRY_ID) VALUES('Michigan','MI',(SELECT C.ID FROM DAS.COUNTRY C WHERE C.ABBREVIATION = 'USA'));
INSERT INTO DAS.STATE (NAME, ABBREVIATION, COUNTRY_ID) VALUES('Minnesota','MN',(SELECT C.ID FROM DAS.COUNTRY C WHERE C.ABBREVIATION = 'USA'));
INSERT INTO DAS.STATE (NAME, ABBREVIATION, COUNTRY_ID) VALUES('Mississippi','MS',(SELECT C.ID FROM DAS.COUNTRY C WHERE C.ABBREVIATION = 'USA'));
INSERT INTO DAS.STATE (NAME, ABBREVIATION, COUNTRY_ID) VALUES('Missouri','MO',(SELECT C.ID FROM DAS.COUNTRY C WHERE C.ABBREVIATION = 'USA'));
INSERT INTO DAS.STATE (NAME, ABBREVIATION, COUNTRY_ID) VALUES('Montana','MT',(SELECT C.ID FROM DAS.COUNTRY C WHERE C.ABBREVIATION = 'USA'));
INSERT INTO DAS.STATE (NAME, ABBREVIATION, COUNTRY_ID) VALUES('Nebraska','NE',(SELECT C.ID FROM DAS.COUNTRY C WHERE C.ABBREVIATION = 'USA'));
INSERT INTO DAS.STATE (NAME, ABBREVIATION, COUNTRY_ID) VALUES('Nevada','NV',(SELECT C.ID FROM DAS.COUNTRY C WHERE C.ABBREVIATION = 'USA'));
INSERT INTO DAS.STATE (NAME, ABBREVIATION, COUNTRY_ID) VALUES('New Hampshire','NH',(SELECT C.ID FROM DAS.COUNTRY C WHERE C.ABBREVIATION = 'USA'));
INSERT INTO DAS.STATE (NAME, ABBREVIATION, COUNTRY_ID) VALUES('New Jersey','NJ',(SELECT C.ID FROM DAS.COUNTRY C WHERE C.ABBREVIATION = 'USA'));
INSERT INTO DAS.STATE (NAME, ABBREVIATION, COUNTRY_ID) VALUES('New Mexico','NM',(SELECT C.ID FROM DAS.COUNTRY C WHERE C.ABBREVIATION = 'USA'));
INSERT INTO DAS.STATE (NAME, ABBREVIATION, COUNTRY_ID) VALUES('New York','NY',(SELECT C.ID FROM DAS.COUNTRY C WHERE C.ABBREVIATION = 'USA'));
INSERT INTO DAS.STATE (NAME, ABBREVIATION, COUNTRY_ID) VALUES('North Carolina','NC',(SELECT C.ID FROM DAS.COUNTRY C WHERE C.ABBREVIATION = 'USA'));
INSERT INTO DAS.STATE (NAME, ABBREVIATION, COUNTRY_ID) VALUES('North Dakota','ND',(SELECT C.ID FROM DAS.COUNTRY C WHERE C.ABBREVIATION = 'USA'));
INSERT INTO DAS.STATE (NAME, ABBREVIATION, COUNTRY_ID) VALUES('Ohio','OH',(SELECT C.ID FROM DAS.COUNTRY C WHERE C.ABBREVIATION = 'USA'));
INSERT INTO DAS.STATE (NAME, ABBREVIATION, COUNTRY_ID) VALUES('Oklahoma','OK',(SELECT C.ID FROM DAS.COUNTRY C WHERE C.ABBREVIATION = 'USA'));
INSERT INTO DAS.STATE (NAME, ABBREVIATION, COUNTRY_ID) VALUES('Oregon','OR',(SELECT C.ID FROM DAS.COUNTRY C WHERE C.ABBREVIATION = 'USA'));
INSERT INTO DAS.STATE (NAME, ABBREVIATION, COUNTRY_ID) VALUES('Pennsylvania','PA',(SELECT C.ID FROM DAS.COUNTRY C WHERE C.ABBREVIATION = 'USA'));
INSERT INTO DAS.STATE (NAME, ABBREVIATION, COUNTRY_ID) VALUES('Puerto Rico','PR',(SELECT C.ID FROM DAS.COUNTRY C WHERE C.ABBREVIATION = 'USA'));
INSERT INTO DAS.STATE (NAME, ABBREVIATION, COUNTRY_ID) VALUES('Rhode Island','RI',(SELECT C.ID FROM DAS.COUNTRY C WHERE C.ABBREVIATION = 'USA'));
INSERT INTO DAS.STATE (NAME, ABBREVIATION, COUNTRY_ID) VALUES('South Carolina','SC',(SELECT C.ID FROM DAS.COUNTRY C WHERE C.ABBREVIATION = 'USA'));
INSERT INTO DAS.STATE (NAME, ABBREVIATION, COUNTRY_ID) VALUES('South Dakota','SD',(SELECT C.ID FROM DAS.COUNTRY C WHERE C.ABBREVIATION = 'USA'));
INSERT INTO DAS.STATE (NAME, ABBREVIATION, COUNTRY_ID) VALUES('Tennessee','TN',(SELECT C.ID FROM DAS.COUNTRY C WHERE C.ABBREVIATION = 'USA'));
INSERT INTO DAS.STATE (NAME, ABBREVIATION, COUNTRY_ID) VALUES('Texas','TX',(SELECT C.ID FROM DAS.COUNTRY C WHERE C.ABBREVIATION = 'USA'));
INSERT INTO DAS.STATE (NAME, ABBREVIATION, COUNTRY_ID) VALUES('Utah','UT',(SELECT C.ID FROM DAS.COUNTRY C WHERE C.ABBREVIATION = 'USA'));
INSERT INTO DAS.STATE (NAME, ABBREVIATION, COUNTRY_ID) VALUES('Vermont','VT',(SELECT C.ID FROM DAS.COUNTRY C WHERE C.ABBREVIATION = 'USA'));
INSERT INTO DAS.STATE (NAME, ABBREVIATION, COUNTRY_ID) VALUES('Virginia','VA',(SELECT C.ID FROM DAS.COUNTRY C WHERE C.ABBREVIATION = 'USA'));
INSERT INTO DAS.STATE (NAME, ABBREVIATION, COUNTRY_ID) VALUES('Washington','WA',(SELECT C.ID FROM DAS.COUNTRY C WHERE C.ABBREVIATION = 'USA'));
INSERT INTO DAS.STATE (NAME, ABBREVIATION, COUNTRY_ID) VALUES('West Virginia','WV',(SELECT C.ID FROM DAS.COUNTRY C WHERE C.ABBREVIATION = 'USA'));
INSERT INTO DAS.STATE (NAME, ABBREVIATION, COUNTRY_ID) VALUES('Wisconsin','WI',(SELECT C.ID FROM DAS.COUNTRY C WHERE C.ABBREVIATION = 'USA'));
INSERT INTO DAS.STATE (NAME, ABBREVIATION, COUNTRY_ID) VALUES('Wyoming','WY',(SELECT C.ID FROM DAS.COUNTRY C WHERE C.ABBREVIATION = 'USA'));
INSERT INTO DAS.STATE (NAME, ABBREVIATION, COUNTRY_ID) VALUES('Alberta','AB',(SELECT C.ID FROM DAS.COUNTRY C WHERE C.ABBREVIATION = 'CAN'));
INSERT INTO DAS.STATE (NAME, ABBREVIATION, COUNTRY_ID) VALUES('British Columbia','BC',(SELECT C.ID FROM DAS.COUNTRY C WHERE C.ABBREVIATION = 'CAN'));
INSERT INTO DAS.STATE (NAME, ABBREVIATION, COUNTRY_ID) VALUES('Manitoba','MB',(SELECT C.ID FROM DAS.COUNTRY C WHERE C.ABBREVIATION = 'CAN'));
INSERT INTO DAS.STATE (NAME, ABBREVIATION, COUNTRY_ID) VALUES('New Brunswick','NB',(SELECT C.ID FROM DAS.COUNTRY C WHERE C.ABBREVIATION = 'CAN'));
INSERT INTO DAS.STATE (NAME, ABBREVIATION, COUNTRY_ID) VALUES('Newfoundland and Labrador','NL',  (SELECT C.ID FROM DAS.COUNTRY C WHERE C.ABBREVIATION = 'CAN'));
INSERT INTO DAS.STATE (NAME, ABBREVIATION, COUNTRY_ID) VALUES('Northwest Territories','NT',(SELECT C.ID FROM DAS.COUNTRY C WHERE C.ABBREVIATION = 'CAN'));
INSERT INTO DAS.STATE (NAME, ABBREVIATION, COUNTRY_ID) VALUES('Nova Scotia','NS', (SELECT C.ID FROM DAS.COUNTRY C WHERE C.ABBREVIATION = 'CAN'));
INSERT INTO DAS.STATE (NAME, ABBREVIATION, COUNTRY_ID) VALUES('Nunavut','NU',(SELECT C.ID FROM DAS.COUNTRY C WHERE C.ABBREVIATION = 'CAN'));
INSERT INTO DAS.STATE (NAME, ABBREVIATION, COUNTRY_ID) VALUES('Prince Edward Island','PE', (SELECT C.ID FROM DAS.COUNTRY C WHERE C.ABBREVIATION = 'CAN'));
INSERT INTO DAS.STATE (NAME, ABBREVIATION, COUNTRY_ID) VALUES('Ontario','ON',(SELECT C.ID FROM DAS.COUNTRY C WHERE C.ABBREVIATION = 'CAN'));
INSERT INTO DAS.STATE (NAME, ABBREVIATION, COUNTRY_ID) VALUES('Quebec','QC', (SELECT C.ID FROM DAS.COUNTRY C WHERE C.ABBREVIATION = 'CAN'));
INSERT INTO DAS.STATE (NAME, ABBREVIATION, COUNTRY_ID) VALUES('Saskatchewan','SK',(SELECT C.ID FROM DAS.COUNTRY C WHERE C.ABBREVIATION = 'CAN'));
INSERT INTO DAS.STATE (NAME, ABBREVIATION, COUNTRY_ID) VALUES('Yukon','YT',(SELECT C.ID FROM DAS.COUNTRY C WHERE C.ABBREVIATION = 'CAN'));