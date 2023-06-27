-- Create datavase structure


-- Create Tables

-- Users are the ones able to use the system
CREATE TABLE AppUser (
    usrUsername VARCHAR(45) NOT NULL PRIMARY KEY UNIQUE,
    usrPassword VARCHAR(255) NOT NULL
);

CREATE TABLE Member (
    memId SERIAL NOT NULL PRIMARY KEY UNIQUE,
    memName VARCHAR(45) NOT NULL,
    rankId INT NOT NULL
);

CREATE TABLE Rank (
    rankId SERIAL NOT NULL PRIMARY KEY UNIQUE,
    rankName VARCHAR(45) NOT NULL,
    rankColor SMALLINT NOT NULL
);

CREATE TABLE Role (
    roleId SERIAL NOT NULL PRIMARY KEY UNIQUE,
    roleName VARCHAR(45) NOT NULL
);

CREATE TABLE MemberRole (
    memRoleId SERIAL NOT NULL PRIMARY KEY UNIQUE,
    memId INT NOT NULL,
    roleId INT NOT NULL
);

CREATE TABLE History (
    hisId SERIAL NOT NULL PRIMARY KEY UNIQUE,
    hisDate TIMESTAMP NOT NULL DEFAULT NOW(),
    hisDescription VARCHAR(255) NOT NULL,
    memId INT NOT NULL,
    rankId INT NULL,
    roleId INT NULL
);

-- Create Foreign Keys
ALTER TABLE Member ADD CONSTRAINT FK_Member_Rank FOREIGN KEY (rankId) REFERENCES Rank (rankId);
ALTER TABLE MemberRole ADD CONSTRAINT FK_MemberRole_User FOREIGN KEY (memId) REFERENCES Member (memId);
ALTER TABLE MemberRole ADD CONSTRAINT FK_MemberRole_Role FOREIGN KEY (roleId) REFERENCES Role (roleId);
ALTER TABLE History ADD CONSTRAINT FK_History_User FOREIGN KEY (memId) REFERENCES Member (memId);
ALTER TABLE History ADD CONSTRAINT FK_History_Rank FOREIGN KEY (rankId) REFERENCES Rank (rankId);
ALTER TABLE History ADD CONSTRAINT FK_History_Role FOREIGN KEY (roleId) REFERENCES Role (roleId);
