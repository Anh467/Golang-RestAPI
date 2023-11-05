INSERT INTO Users (Email, Password, Role, FullName, Address) VALUES
(N'john@example.com', N'yenhi123', N'user', N'Nguyễn Ngọc Yến Nhi', N'40 á đù Võ Chí Công'),
(N'jane@example.com', N'hashed_password', N'user', N'Trần Văn Hùng', NULL),
(N'admin@example.com', N'hashed_password', N'admin', N'Lê Thị Thu', NULL),
(N'sarah@example.com', N'hashed_password', N'user', N'Phạm Minh Đức', NULL),
(N'david@example.com', N'hashed_password', N'user', N'Hoàng Thị Nga', NULL),
(N'abc@123gmail.com', N'hashed_passwordpasswordddpassworddd', N'user', N'32r433r3', NULL),
(N'abc@123gm123gm1123gm1ail.com', N'hashed_passwordpasswordddpassworddd', N'user', N'32r433r3', NULL),
(N'abab1ab1c@123gmail.com', N'hashed_passwordpasswordddpassworddd', N'user', N'32r433r3', NULL);


INSERT INTO shop.Products (Name, Description, Price, CategoryID) VALUES
(N'Áo loong heheheh', N'Áo khoác nam dáng dài dàidàidàidàidàidàidàidàidàidàidàidàidàidàimàu đenđenmmmmmđenmmmmm', 109.99, 2),
(N'Quần jeans đen', N'Quần jeans nam màu đen', 39.99, 2),
(N'Áo khoác nỉ đen', N'Áo khoác nam dáng dài màu đen', 49.99, 3),
(N'Quần thể thao xanh', N'Quần thể thao nam màu xanh', 19.99, 4),
(N'Áo len cổ lọ xám', N'Áo len nam cổ lọ màu xám', 34.99, 5),
(N'Áo khoác nỉ đen', N'Áo khoác nam dáng dài màu đen', 49.99, 3),
(N'Áo khoác nỉ đenđennnnnnnđennnnnnn', N'Áo khoác nam dáng dài màu đenđenmmmmmđenmmmmm', 99.99, 4),
(N'Áo khoác nỉ đenđennnnnnnđennnnnnn', N'Áo khoác nam dáng dài màu đenđenmmmmmđenmmmmm', 99.99, 4),
(N'Áo khoác nỉ đenđennnnnnnđennnnnnn', N'Áo khoác nam dáng dài màu đenđenmmmmmđenmmmmm', 99.99, 4),
(N'Áo khoác nỉ đenđennnnnnnđennnnnnn', N'Áo khoác nam dáng dài màu đenđenmmmmmđenmmmmm', 99.99, 4),
(N'Áo loong cuuu nnnnnnn', N'Áo khoác nam dáng dài màu đenđenmmmmmđenmmmmm', 99.99, 4),
(N'Áo loong cuuu nnnnnnn', N'Áo khoác nam dáng dài màu đenđenmmmmmđenmmmmm', 99.99, 4),
(N'Áo loong cuuu nnnnnnn', N'Áo khoác nam dáng dài màu đenđenmmmmmđenmmmmm', 99.99, 4),
(N'Áo loong cuuu nnnnnnn', N'Áo khoác nam dáng dài màu đenđenmmmmmđenmmmmm', 99.99, 4);



INSERT INTO shop.Categories (Name) VALUES 
(N'Quần jeans Hàng hiệu loại 1'),
(N'Quần jeans Hàng hiệu loại 1'),
(N'Áo khoác'),
(N'Quần thể thao'),
 (N'Áo len');


INSERT INTO shop.Orders (UserID, OrderDate, Status, Address) VALUES
(1, CAST(N'2023-10-27T10:11:57.473' AS DateTime), N'delivering', N'44 Hoàng hoa thám'),
(1, CAST(N'2023-10-27T10:11:57.473' AS DateTime), N'canceled', NULL),
(1, CAST(N'2023-10-27T10:11:57.473' AS DateTime), N'canceled', NULL),
(1, CAST(N'2023-10-27T10:11:57.473' AS DateTime), N'Đã đặt hàng', NULL),
(1, CAST(N'2023-10-27T10:11:57.473' AS DateTime), N'canceled', NULL),
(1, NULL, N'confirm', NULL),
(1, NULL, N'confirm', NULL),
(1, NULL, N'confirm', NULL),
(1, NULL, N'confirm', NULL),
(1, CAST(N'2006-01-02T15:04:05.000' AS DateTime), N'confirm', NULL),
(1, CAST(N'2006-01-02T15:04:05.000' AS DateTime), N'confirm', NULL),
(1, CAST(N'2006-01-02T15:04:05.000' AS DateTime), N'confirm', NULL),
(1, CAST(N'2006-01-02T15:04:05.000' AS DateTime), N'confirm', NULL),
(1, CAST(N'2006-01-02T15:04:05.000' AS DateTime), N'delivering', NULL),
(1, CAST(N'2006-01-02T15:04:05.000' AS DateTime), N'delivering', NULL),
(1, CAST(N'2006-01-02T15:04:05.000' AS DateTime), N'confirm', NULL),
(1, CAST(N'2006-01-02T15:04:05.000' AS DateTime), N'confirm', NULL);

INSERT INTO shop.OrderDetails (OrderID, ProductID, Quantity, Price) VALUES
(1, 1, 2, CAST(29.99 AS Decimal(10, 2))),
(2, 2, 1, CAST(39.99 AS Decimal(10, 2))),
(3, 3, 3, CAST(49.99 AS Decimal(10, 2))),
(4, 4, 2, CAST(19.99 AS Decimal(10, 2))),
(5, 5, 4, CAST(34.99 AS Decimal(10, 2)));

