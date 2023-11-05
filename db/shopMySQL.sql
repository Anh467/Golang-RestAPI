-- Tạo Cơ sở dữ liệu
CREATE DATABASE Shop;

-- Sử dụng Cơ sở dữ liệu
USE Shop;

-- Tạo bảng Users
CREATE TABLE Users (
    UserID INT AUTO_INCREMENT,
    Email VARCHAR(100),
    Password VARCHAR(100),
    Role VARCHAR(20),
    FullName VARCHAR(100),
    Address VARCHAR(200),
    PRIMARY KEY (UserID),
    UNIQUE (Email)
);

-- Tạo bảng Cart
CREATE TABLE Cart (
    UserID INT,
    ProductID INT,
    Quantity INT,
    PRIMARY KEY (UserID, ProductID)
);

-- Tạo bảng Categories
CREATE TABLE Categories (
    CategoryID INT AUTO_INCREMENT,
    Name VARCHAR(50),
    PRIMARY KEY (CategoryID)
);
-- Tạo bảng Products
CREATE TABLE Products (
    ProductID INT AUTO_INCREMENT,
    Name TEXT,
    Description TEXT,
    Price FLOAT,
    CategoryID INT,
    PRIMARY KEY (ProductID),
    FOREIGN KEY (CategoryID) REFERENCES Categories(CategoryID)
);

CREATE TABLE Orders (
    OrderID INT AUTO_INCREMENT,
    UserID INT,
    OrderDate DATETIME,
    Status VARCHAR(50),
    Address VARCHAR(200),
    PRIMARY KEY (OrderID),
    FOREIGN KEY (UserID) REFERENCES Users(UserID)
);

-- Tạo bảng OrderDetails
CREATE TABLE OrderDetails (
    OrderDetailID INT AUTO_INCREMENT,
    OrderID INT,
    ProductID INT,
    Quantity INT,
    Price DECIMAL(10, 2),
    PRIMARY KEY (OrderDetailID),
    FOREIGN KEY (OrderID) REFERENCES Orders(OrderID),
    FOREIGN KEY (ProductID) REFERENCES Products(ProductID)
);

-- Tạo bảng Orders


-- Tạo bảng Reviews
CREATE TABLE Reviews (
    ReviewID INT AUTO_INCREMENT,
    ProductID INT,
    UserID INT,
    Rating INT,
    Comment TEXT,
    PRIMARY KEY (ReviewID),
    FOREIGN KEY (ProductID) REFERENCES Products(ProductID),
    FOREIGN KEY (UserID) REFERENCES Users(UserID)
);

