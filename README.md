# sqlc-spanner-sample
sqlc with spanner sample implementation.

## Getting started
### Install tools
```
go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
```

### DB setup
- Create Spanner database
- Create tables

```sql
CREATE TABLE Users (
  ID STRING(36) NOT NULL,
  Name STRING(255)
) PRIMARY KEY(ID);

CREATE TABLE Addresses (
  ID STRING(36) NOT NULL,
  UserID STRING(36) NOT NULL,
  Address STRING(255) NOT NULL,
  CONSTRAINT FK_UserAddress FOREIGN KEY (UserID) REFERENCES Users (ID)
) PRIMARY KEY(ID);

Create Index AddressesByUserID On Addresses(UserID);
```

- Insert sample data

```sql
-- Users
INSERT INTO
  Users (ID, Name)
VALUES
  ('199f8059-558a-4c6f-aad3-526859cfa88e', 'User A'),
  ('7d586bac-1c9e-4c3d-aefe-53c8649352f0', 'User B'),
  ('a94c740a-3791-492d-88be-b05cf0b85252', 'User C');

-- Addresses
INSERT INTO
  Addresses (ID, UserID, Address)
VALUES
  ('3b9a8a97-5975-46f9-81ff-2b4b7b74d996', '199f8059-558a-4c6f-aad3-526859cfa88e', 'Address A'), -- User A
  ('bf640863-0636-4ba4-aa77-b4b20bfff7c0', '7d586bac-1c9e-4c3d-aefe-53c8649352f0', 'Address B'), -- User B
  ('2c002eae-e0fd-40bf-b09f-c8f582302500', '199f8059-558a-4c6f-aad3-526859cfa88e', 'Address C'), -- User A
  ('d3be1e3f-0ad4-4d7e-bf43-fe61e890cd32', 'a94c740a-3791-492d-88be-b05cf0b85252', 'Address D'), -- User C
  ('ab02a7fa-7edc-4b1d-8ef6-ae82b69f9d00', '7d586bac-1c9e-4c3d-aefe-53c8649352f0', 'Address E'), -- User B
  ('0ea07838-34dd-4fa4-af2f-06babaab75a0', '199f8059-558a-4c6f-aad3-526859cfa88e', 'Address F'); -- User A
```

### Gcloud setup
```
gcloud login
```
