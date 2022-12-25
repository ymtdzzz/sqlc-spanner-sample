CREATE TABLE Users (
  ID VARCHAR(36) NOT NULL PRIMARY KEY,
  Name VARCHAR(255) NOT NULL
);

CREATE TABLE Addresses (
  ID VARCHAR(36) NOT NULL PRIMARY KEY,
  UserID VARCHAR(36) NOT NULL,
  Address VARCHAR(255) NOT NULL
);

CREATE TABLE Counters (
  ID VARCHAR(36) NOT NULL PRIMARY KEY,
  Count BIGINT NOT NULL
);
