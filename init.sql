CREATE TABLE schema_version (
    version INTEGER NOT NULL PRIMARY KEY
);
INSERT INTO schema_version VALUES (1);

CREATE TABLE sport (
    id         INTEGER  NOT NULL PRIMARY KEY,
    name_en    STRING   NOT NULL,
    name_en_us STRING   NULL,
    name_en_gb STRING   NULL,
    name_es    STRING   NULL,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE UNIQUE INDEX sport_ak1 ON sport (name_en);
INSERT INTO sport (id, name_en, name_en_us, name_en_gb, name_es) 
    VALUES (1, 'football', 'soccer', 'football', 'fútbol');

CREATE TABLE league (
    id         INTEGER  NOT NULL PRIMARY KEY,
    sport_id   INTEGER  NOT NULL,
    name       STRING   NOT NULL,    
    short_name STRING   NOT NULL,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY(sport_id) REFERENCES sport(id)
);
CREATE UNIQUE INDEX league_ak1 ON league(sport_id, name);
CREATE UNIQUE INDEX league_ak2 ON league(sport_id, short_name);
INSERT INTO league (id, sport_id, name, short_name) 
    VALUES (1, 1, 'Major League Soccer', 'MLS');

CREATE TABLE league_season (
    id          INTEGER  NOT NULL PRIMARY KEY,
    league_id   INTEGER  NOT NULL,
    season_name STRING   NOT NULL,
    start_date  DATE     NOT NULL,
    end_date    DATE     NOT NULL,
    updated_at  DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,    
    FOREIGN KEY(league_id) REFERENCES league(id)
);
CREATE UNIQUE INDEX league_season_ak1 ON league_season (league_id, start_date);
INSERT INTO league_season (id, league_id, season_name, start_date, end_date) 
    VALUES (1, 1, '2013', '2013-03-02', '2013-10-23'),
           (2, 1, '2014', '2014-03-08', '2014-10-26');

CREATE TABLE league_season_group (
    id                INTEGER   NOT NULL PRIMARY KEY,
    league_season_id  INTEGER   NOT NULL,
    season_group_type STRING    NOT NULL,
    season_group_name STRING    NOT NULL,
    updated_at        DATETIME  NOT NULL DEFAULT CURRENT_TIMESTAMP,    
    FOREIGN KEY(league_season_id) REFERENCES league_season(id)
);
INSERT INTO league_season_group (id, league_season_id, season_group_type, season_group_name) 
    VALUES (1, 1, 'conference', 'eastern'),
           (2, 1, 'conference', 'western'),
           (3, 2, 'conference', 'eastern'),
           (4, 2, 'conference', 'western');

CREATE TABLE league_team (
    id         INTEGER  NOT NULL PRIMARY KEY,
    league_id  INTEGER  NOT NULL,
    name       STRING   NOT NULL,
    short_name STRING   NOT NULL,        
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,    
    FOREIGN KEY(league_id) REFERENCES league(id)
);
CREATE UNIQUE INDEX league_team_ak1 ON league_team (league_id, name);
CREATE UNIQUE INDEX league_team_ak2 ON league_team (league_id, short_name);
INSERT INTO league_team (id, league_id, name, short_name) 
    VALUES (1, 1, 'Chivas USA', 'Chivas'),
           (2, 1, 'Columbus Crew', 'Crew'),
           (3, 1, 'FC Dallas', 'FC Dallas'),
           (4, 1, 'LA Galaxy', 'Galaxy'),
           (5, 1, 'Houston Dynamo', 'Dynamo'),
           (6, 1, 'Montreal Impact', 'Impact'),
           (7, 1, 'Colorado Rapids', 'Rapids'),
           (8, 1, 'New York Red Bull', 'NYRB'),
           (9, 1, 'New England Revolution', 'Revs'),
           (10, 1, 'Real Salt Lake', 'RSL'),
           (11, 1, 'San Jose Earthquakes', 'Quakes'),
           (12, 1, 'Seattle Sounders FC', 'Sounders'),
           (13, 1, 'Sporting Kansas City', 'Sporting KC'),
           (14, 1, 'Toronto FC', 'TFC'),
           (15, 1, 'Portland Timbers', 'Timbers'),
           (16, 1, 'Philadelphia Union', 'Union'),
           (17, 1, 'Vancouver Whitecaps', 'Whitecaps');

CREATE TABLE league_team_group (    
    id                     INTEGER  NOT NULL PRIMARY KEY,
    league_team_id         INTEGER  NOT NULL,
    league_season_group_id INTEGER  NOT NULL,
    updated_at             DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,    
    FOREIGN KEY(league_team_id) REFERENCES league_team(id),
    FOREIGN KEY(league_season_group_id) REFERENCES league_season_group(id)   
);
-- MLS 2013
INSERT INTO league_team_group (id, league_team_id, league_season_group_id) 
    VALUES (1, 1, 2),
           (2, 2, 1),
           (3, 3, 2),
           (4, 4, 2),
           (5, 5, 1),
           (6, 6, 1),
           (7, 7, 2),
           (8, 8, 1),
           (9, 9, 1),
           (10, 10, 2),
           (11, 11, 2),
           (12, 12, 2),
           (13, 13, 1),
           (14, 14, 1),
           (15, 15, 2),
           (16, 16, 1),
           (17, 17, 2);
-- MLS 2014
INSERT INTO league_team_group (id, league_team_id, league_season_group_id) 
    VALUES (18, 1, 4),
           (19, 2, 3),
           (20, 3, 4),
           (21, 4, 4),
           (22, 5, 3),
           (23, 6, 3),
           (24, 7, 4),
           (25, 8, 3),
           (26, 9, 3),
           (27, 10, 4),
           (28, 11, 4),
           (29, 12, 4),
           (30, 13, 3),
           (31, 14, 3),
           (32, 15, 4),
           (33, 16, 3),
           (34, 17, 4);
