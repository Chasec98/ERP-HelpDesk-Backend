CREATE DATABASE `hd-app`;

USE `hd-app`;

CREATE TABLE `Users` (
  `ID` int NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `FirstName` varchar(255) NULL,
  `LastName` varchar(255) NULL,
  `Email` varchar(255) NULL,
  `PhoneNumber` varchar(255) NULL,
  `Username` varchar(255) NULL,
  `Password` varchar(255) NULL,
  `Active` tinyint(1) NOT NULL
);

CREATE TABLE `Tickets` (
  `ID` int NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `AssignedToID` int NULL,
  `CreatedByID` int NOT NULL,
  `Subject` tinytext NOT NULL,
  `Body` text NOT NULL,
  `Solution` text NULL,
  `CreatedDate` datetime NOT NULL,
  `ClosedDate` datetime NULL,
  FOREIGN KEY (`CreatedByID`) REFERENCES `Users` (`ID`),
  FOREIGN KEY (`AssignedToID`) REFERENCES `Users` (`ID`),
  FULLTEXT KEY `Solution` (`Solution`),
  FULLTEXT KEY `Body` (`Body`),
  FULLTEXT KEY `Subject` (`Subject`)
);

CREATE TABLE `Sessions` (
  `ID` int NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `SessionID` varchar(255) NOT NULL,
  `UserID` int NOT NULL,
  `Expires` datetime NOT NULL,
  FOREIGN KEY (`UserID`) REFERENCES `Users` (`ID`)
);

CREATE TABLE `Roles` (
  `ID` int NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `Name` varchar(255) NOT NULL
);

CREATE TABLE `Permissions` (
  `ID` int NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `Name` varchar(255) NOT NULL,
  `Description` text NULL
);

CREATE TABLE `RolesPermissions` (
  `ID` int NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `RoleID` int NOT NULL,
  `PermissionID` int NOT NULL,
  FOREIGN KEY (`RoleID`) REFERENCES `Roles` (`ID`) ON DELETE CASCADE,
  FOREIGN KEY (`PermissionID`) REFERENCES `Permissions` (`ID`) ON DELETE CASCADE
);

CREATE TABLE `UserRoles` (
  `ID` int NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `RoleID` int NOT NULL,
  `UserID` int NOT NULL,
  FOREIGN KEY (`RoleID`) REFERENCES `Roles` (`ID`) ON DELETE CASCADE,
  FOREIGN KEY (`UserID`) REFERENCES `Users` (`ID`)
);

INSERT INTO `Roles` (`Name`)
VALUES ('Non-User');

INSERT INTO `Permissions` (`Name`)
VALUES ('Read Tickets'),
('Create Tickets'),
('Edit Tickets'),
('Read Users'),
('Create Users'),
('Edit Users'),
('Read Roles'),
('Create Roles'),
('Edit Roles');