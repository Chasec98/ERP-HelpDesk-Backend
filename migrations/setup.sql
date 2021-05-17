CREATE DATABASE `tickets`;

USE `tickets`;

CREATE TABLE `Tickets` (
  `ID` int NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `AssignedToID` int NULL,
  `CreatedByID` int NOT NULL,
  `Subject` tinytext NOT NULL,
  `Body` text NOT NULL,
  `Solution` text NULL,
  `CreatedDate` datetime NOT NULL,
  `ClosedDate` datetime NULL,
  FULLTEXT KEY `Solution` (`Solution`),
  FULLTEXT KEY `Body` (`Body`),
  FULLTEXT KEY `Subject` (`Subject`)
);

CREATE TABLE `Comments` (
  `ID` int NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `Text` longtext NOT NULL,
  `CreatedByID` int NOT NULL,
  `TicketID` int NOT NULL,
  FOREIGN KEY (`TicketID`) REFERENCES `Tickets` (`ID`) ON DELETE CASCADE,
  FULLTEXT KEY `Text` (`Text`)
);

CREATE DATABASE `users`;

USE `users`;

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
  `RoleID` int NOT NULL,
  `PermissionID` int NOT NULL,
  PRIMARY KEY (`RoleID`, `PermissionID`),
  FOREIGN KEY (`RoleID`) REFERENCES `Roles` (`ID`) ON DELETE CASCADE,
  FOREIGN KEY (`PermissionID`) REFERENCES `Permissions` (`ID`) ON DELETE CASCADE
);

CREATE TABLE `UserRoles` (
  `RoleID` int NOT NULL,
  `UserID` int NOT NULL,
  PRIMARY KEY (`RoleID`, `UserID`),
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