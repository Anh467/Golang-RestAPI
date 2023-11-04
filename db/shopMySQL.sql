

/****** Object: Database `Shop` Script Date: 11/2/2023 12:35:45 PM ******/
CREATE DATABASE IF NOT EXISTS `Shop`;

USE `Shop`;

/****** Object: Table `Cart` Script Date: 11/2/2023 12:35:45 PM ******/
CREATE TABLE `Cart` (
    `UserID` INT NULL,
    `ProductID` INT NULL,
    `Quantity` INT NULL
);

/****** Object: Table `Categories` Script Date: 11/2/2023 12:35:45 PM ******/
CREATE TABLE `Categories` (
    `CategoryID` INT AUTO_INCREMENT NOT NULL,
    `Name` VARCHAR(50) NULL,
    PRIMARY KEY (`CategoryID`)
);

/****** Object: Table `OrderDetails` Script Date: 11/2/2023 12:35:45 PM ******/
CREATE TABLE `OrderDetails` (
    `OrderDetailID` INT AUTO_INCREMENT NOT NULL,
    `OrderID` INT NULL,
    `ProductID` INT NULL,
    `Quantity` INT NULL,
    `Price` DECIMAL(10, 2) NULL,
    PRIMARY KEY (`OrderDetailID`)
);

/****** Object: Table `Orders` Script Date: 11/2/2023 12:35:45 PM ******/
CREATE TABLE `Orders` (
    `OrderID` INT AUTO_INCREMENT NOT NULL,
    `UserID` INT NULL,
    `OrderDate` DATETIME NULL,
    `Status` VARCHAR(50) NULL,
    `Address` VARCHAR(200) NULL,
    PRIMARY KEY (`OrderID`)
);

/****** Object: Table `Products` Script Date: 11/2/2023 12:35:45 PM ******/
CREATE TABLE `Products` (
    `ProductID` INT AUTO_INCREMENT NOT NULL,
    `Name` VARCHAR(255) NULL,
    `Description` TEXT NULL,
    `Price` DOUBLE NULL,
    `CategoryID` INT NULL,
    PRIMARY KEY (`ProductID`)
);

/****** Object: Table `Reviews` Script Date: 11/2/2023 12:35:45 PM ******/
CREATE TABLE `Reviews` (
    `ReviewID` INT AUTO_INCREMENT NOT NULL,
    `ProductID` INT NULL,
    `UserID` INT NULL,
    `Rating` INT NULL,
    `Comment` TEXT NULL,
    PRIMARY KEY (`ReviewID`)
);

/****** Object: Table `Users` Script Date: 11/2/2023 12:35:45 PM ******/
CREATE TABLE `Users` (
    `UserID` INT AUTO_INCREMENT NOT NULL,
    `Email` VARCHAR(100) NULL,
    `Password` VARCHAR(100) NULL,
    `Role` VARCHAR(20) NULL,
    `FullName` VARCHAR(100) NULL,
    `Address` VARCHAR(200) NULL,
    PRIMARY KEY (`UserID`)
);

INSERT INTO `Cart` (`UserID`, `ProductID`, `Quantity`) VALUES (1, 1, 5);
INSERT INTO `Cart` (`UserID`, `ProductID`, `Quantity`) VALUES (1, 2, 12);
INSERT INTO `Cart` (`UserID`, `ProductID`, `Quantity`) VALUES (1, 3, 3);
INSERT INTO `Cart` (`UserID`, `ProductID`, `Quantity`) VALUES (1, 4, 11);
INSERT INTO `Cart` (`UserID`, `ProductID`, `Quantity`) VALUES (1, 5, 17);

SET IDENTITY_INSERT `Categories` ON;

INSERT INTO `Categories` (`CategoryID`, `Name`) VALUES (1, N'Quần jeans Hàng hiệu loại 1');
INSERT INTO `Categories` (`CategoryID`, `Name`) VALUES (2, N'Quần jeans Hàng hiệu loại 1');
INSERT INTO `Categories` (`CategoryID`, `Name`) VALUES (3, N'Áo khoác');
INSERT INTO `Categories` (`CategoryID`, `Name`) VALUES (4, N'Quần thể thao');
INSERT INTO `Categories` (`CategoryID`, `Name`) VALUES (5, N'Áo len');

SET IDENTITY_INSERT `Categories` OFF;

SET IDENTITY_INSERT `OrderDetails` ON;

INSERT INTO `OrderDetails` (`OrderDetailID`, `OrderID`, `ProductID`, `Quantity`, `Price`) VALUES (1, 1, 1, 2, CAST(29.99 AS DECIMAL(10, 2)));
INSERT INTO `OrderDetails` (`OrderDetailID`, `OrderID`, `ProductID`, `Quantity`, `Price`) VALUES (2, 2, 2, 1, CAST(39.99 AS DECIMAL(10, 2)));
INSERT INTO `OrderDetails` (`OrderDetailID`, `OrderID`, `ProductID`, `Quantity`, `Price`) VALUES (3, 3, 3, 3, CAST(49.99 AS DECIMAL(10, 2)));
INSERT INTO `OrderDetails` (`OrderDetailID`, `OrderID`, `ProductID`, `Quantity`, `Price`) VALUES (4, 4, 4, 2, CAST(19.99 AS DECIMAL(10, 2)));
INSERT INTO `OrderDetails` (`OrderDetailID`, `OrderID`, `ProductID`, `Quantity`, `Price`) VALUES (5, 5, 5, 4, CAST(34.99 AS DECIMAL(10, 2)));

SET IDENTITY_INSERT `OrderDetails` OFF;

SET IDENTITY_INSERT `Orders` ON;

INSERT INTO `Orders` (`OrderID`, `UserID`, `OrderDate`, `Status`, `Address`) VALUES (1, 1, CAST(N'2023-10-27T10:11:57.473' AS DATETIME), N'delivering', N'44 Hoàng hoa thám ');
-- Các lệnh INSERT khác ở đây...

SET IDENTITY_INSERT `Orders` OFF;

SET IDENTITY_INSERT `Products` ON;

-- INSERT INTO `Products` ở đây...

SET IDENTITY_INSERT `Products` OFF;

SET IDENTITY_INSERT `Users` ON;

-- INSERT INTO `Users` ở đây...

SET IDENTITY_INSERT `Users` OFF;

/****** Object: Index `unique_UserID_ProductID` Script Date: 11/2/2023 12:35:45 PM ******/
ALTER TABLE `Cart` ADD CONSTRAINT `unique_UserID_ProductID` UNIQUE NONCLUSTERED (`UserID` ASC, `ProductID` ASC);

/****** Object: Index `UC_Email` Script Date: 11/2/2023 12:35:45 PM ******/
ALTER TABLE `Users` ADD CONSTRAINT `UC_Email` UNIQUE NONCLUSTERED (`Email` ASC);

ALTER TABLE `Cart` WITH CHECK ADD FOREIGN KEY (`ProductID`) REFERENCES `Products` (`ProductID`);
ALTER TABLE `Cart` WITH CHECK ADD FOREIGN KEY (`UserID`) REFERENCES `Users` (`UserID`);
ALTER TABLE `OrderDetails` WITH CHECK ADD FOREIGN KEY (`OrderID`) REFERENCES `Orders` (`OrderID`);
ALTER TABLE `OrderDetails` WITH CHECK ADD FOREIGN KEY (`ProductID`) REFERENCES `Products` (`ProductID`);
ALTER TABLE `Orders` WITH CHECK ADD FOREIGN KEY (`UserID`) REFERENCES `Users` (`UserID`);
ALTER TABLE `Products` WITH CHECK ADD FOREIGN KEY (`CategoryID`) REFERENCES `Categories` (`CategoryID`);
ALTER TABLE `Reviews` WITH CHECK ADD FOREIGN KEY (`ProductID`) REFERENCES `Products` (`ProductID`);
ALTER TABLE `Reviews` WITH CHECK ADD FOREIGN KEY (`UserID`) REFERENCES `Users` (`UserID`);
