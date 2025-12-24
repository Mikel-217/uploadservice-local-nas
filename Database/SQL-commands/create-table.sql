CREATE TABLE Users (
    UserID int NOT NULL AUTO_INCREMENT,
    UserName varchar(255),
    UserPW varchar(255),

    PRIMARY KEY (UserID)
);


CREATE TABLE UserDirectorys (
    DirID int NOT NULL AUTO_INCREMENT,
    UserID int NOT NULL,
    DirName varchar(255),
    DirPath varchar(255),

    PRIMARY KEY (DirID),
    FOREIGN KEY (UserID) REFERENCES Users(UserID)
);


CREATE TABLE ActiveAccessTokens (
    TokenID int NOT NULL AUTO_INCREMENT,
    ActiveToken varchar(255),
    ExpirationDate TIMESTAMP,

    PRIMARY KEY (TokenID)
);


CREATE TABLE UserFiles (
    FileID int NOT NULL AUTO_INCREMENT,
    FileName varchar(200),
    FilePath varchar(255),
    DirID int NOT NULL,
    UserID int NOT NULL,

    PRIMARY KEY (FileID),
    FOREIGN KEY (DirID) REFERENCES UserDirectorys(DirID),
    FOREIGN KEY (UserID) REFERENCES Users(UserID)
);
