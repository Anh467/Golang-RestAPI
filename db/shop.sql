USE [master]
GO
/****** Object:  Database [Shop]    Script Date: 11/2/2023 12:35:45 PM ******/
CREATE DATABASE [Shop]
 CONTAINMENT = NONE
 ON  PRIMARY 
( NAME = N'Shop', FILENAME = N'C:\Program Files\Microsoft SQL Server\MSSQL16.SQLEXPRESS\MSSQL\DATA\Shop.mdf' , SIZE = 8192KB , MAXSIZE = UNLIMITED, FILEGROWTH = 65536KB )
 LOG ON 
( NAME = N'Shop_log', FILENAME = N'C:\Program Files\Microsoft SQL Server\MSSQL16.SQLEXPRESS\MSSQL\DATA\Shop_log.ldf' , SIZE = 8192KB , MAXSIZE = 2048GB , FILEGROWTH = 65536KB )
 WITH CATALOG_COLLATION = DATABASE_DEFAULT, LEDGER = OFF
GO
ALTER DATABASE [Shop] SET COMPATIBILITY_LEVEL = 160
GO
IF (1 = FULLTEXTSERVICEPROPERTY('IsFullTextInstalled'))
begin
EXEC [Shop].[dbo].[sp_fulltext_database] @action = 'enable'
end
GO
ALTER DATABASE [Shop] SET ANSI_NULL_DEFAULT OFF 
GO
ALTER DATABASE [Shop] SET ANSI_NULLS OFF 
GO
ALTER DATABASE [Shop] SET ANSI_PADDING OFF 
GO
ALTER DATABASE [Shop] SET ANSI_WARNINGS OFF 
GO
ALTER DATABASE [Shop] SET ARITHABORT OFF 
GO
ALTER DATABASE [Shop] SET AUTO_CLOSE ON 
GO
ALTER DATABASE [Shop] SET AUTO_SHRINK OFF 
GO
ALTER DATABASE [Shop] SET AUTO_UPDATE_STATISTICS ON 
GO
ALTER DATABASE [Shop] SET CURSOR_CLOSE_ON_COMMIT OFF 
GO
ALTER DATABASE [Shop] SET CURSOR_DEFAULT  GLOBAL 
GO
ALTER DATABASE [Shop] SET CONCAT_NULL_YIELDS_NULL OFF 
GO
ALTER DATABASE [Shop] SET NUMERIC_ROUNDABORT OFF 
GO
ALTER DATABASE [Shop] SET QUOTED_IDENTIFIER OFF 
GO
ALTER DATABASE [Shop] SET RECURSIVE_TRIGGERS OFF 
GO
ALTER DATABASE [Shop] SET  ENABLE_BROKER 
GO
ALTER DATABASE [Shop] SET AUTO_UPDATE_STATISTICS_ASYNC OFF 
GO
ALTER DATABASE [Shop] SET DATE_CORRELATION_OPTIMIZATION OFF 
GO
ALTER DATABASE [Shop] SET TRUSTWORTHY OFF 
GO
ALTER DATABASE [Shop] SET ALLOW_SNAPSHOT_ISOLATION OFF 
GO
ALTER DATABASE [Shop] SET PARAMETERIZATION SIMPLE 
GO
ALTER DATABASE [Shop] SET READ_COMMITTED_SNAPSHOT OFF 
GO
ALTER DATABASE [Shop] SET HONOR_BROKER_PRIORITY OFF 
GO
ALTER DATABASE [Shop] SET RECOVERY SIMPLE 
GO
ALTER DATABASE [Shop] SET  MULTI_USER 
GO
ALTER DATABASE [Shop] SET PAGE_VERIFY CHECKSUM  
GO
ALTER DATABASE [Shop] SET DB_CHAINING OFF 
GO
ALTER DATABASE [Shop] SET FILESTREAM( NON_TRANSACTED_ACCESS = OFF ) 
GO
ALTER DATABASE [Shop] SET TARGET_RECOVERY_TIME = 60 SECONDS 
GO
ALTER DATABASE [Shop] SET DELAYED_DURABILITY = DISABLED 
GO
ALTER DATABASE [Shop] SET ACCELERATED_DATABASE_RECOVERY = OFF  
GO
ALTER DATABASE [Shop] SET QUERY_STORE = ON
GO
ALTER DATABASE [Shop] SET QUERY_STORE (OPERATION_MODE = READ_WRITE, CLEANUP_POLICY = (STALE_QUERY_THRESHOLD_DAYS = 30), DATA_FLUSH_INTERVAL_SECONDS = 900, INTERVAL_LENGTH_MINUTES = 60, MAX_STORAGE_SIZE_MB = 1000, QUERY_CAPTURE_MODE = AUTO, SIZE_BASED_CLEANUP_MODE = AUTO, MAX_PLANS_PER_QUERY = 200, WAIT_STATS_CAPTURE_MODE = ON)
GO
USE [Shop]
GO
/****** Object:  Table [dbo].[Cart]    Script Date: 11/2/2023 12:35:45 PM ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[Cart](
	[UserID] [int] NULL,
	[ProductID] [int] NULL,
	[Quantity] [int] NULL
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[Categories]    Script Date: 11/2/2023 12:35:45 PM ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[Categories](
	[CategoryID] [int] IDENTITY(1,1) NOT NULL,
	[Name] [nvarchar](50) NULL,
PRIMARY KEY CLUSTERED 
(
	[CategoryID] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[OrderDetails]    Script Date: 11/2/2023 12:35:45 PM ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[OrderDetails](
	[OrderDetailID] [int] IDENTITY(1,1) NOT NULL,
	[OrderID] [int] NULL,
	[ProductID] [int] NULL,
	[Quantity] [int] NULL,
	[Price] [decimal](10, 2) NULL,
PRIMARY KEY CLUSTERED 
(
	[OrderDetailID] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[Orders]    Script Date: 11/2/2023 12:35:45 PM ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[Orders](
	[OrderID] [int] IDENTITY(1,1) NOT NULL,
	[UserID] [int] NULL,
	[OrderDate] [datetime] NULL,
	[Status] [nvarchar](50) NULL,
	[Address] [nvarchar](200) NULL,
PRIMARY KEY CLUSTERED 
(
	[OrderID] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[Products]    Script Date: 11/2/2023 12:35:45 PM ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[Products](
	[ProductID] [int] IDENTITY(1,1) NOT NULL,
	[Name] [nvarchar](max) NULL,
	[Description] [nvarchar](max) NULL,
	[Price] [float] NULL,
	[CategoryID] [int] NULL,
PRIMARY KEY CLUSTERED 
(
	[ProductID] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY] TEXTIMAGE_ON [PRIMARY]
GO
/****** Object:  Table [dbo].[Reviews]    Script Date: 11/2/2023 12:35:45 PM ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[Reviews](
	[ReviewID] [int] IDENTITY(1,1) NOT NULL,
	[ProductID] [int] NULL,
	[UserID] [int] NULL,
	[Rating] [int] NULL,
	[Comment] [nvarchar](max) NULL,
PRIMARY KEY CLUSTERED 
(
	[ReviewID] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY] TEXTIMAGE_ON [PRIMARY]
GO
/****** Object:  Table [dbo].[Users]    Script Date: 11/2/2023 12:35:45 PM ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[Users](
	[UserID] [int] IDENTITY(1,1) NOT NULL,
	[Email] [nvarchar](100) NULL,
	[Password] [nvarchar](100) NULL,
	[Role] [nvarchar](20) NULL,
	[FullName] [nvarchar](100) NULL,
	[Address] [nvarchar](200) NULL,
PRIMARY KEY CLUSTERED 
(
	[UserID] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
INSERT [dbo].[Cart] ([UserID], [ProductID], [Quantity]) VALUES (1, 1, 5)
INSERT [dbo].[Cart] ([UserID], [ProductID], [Quantity]) VALUES (1, 2, 12)
INSERT [dbo].[Cart] ([UserID], [ProductID], [Quantity]) VALUES (1, 3, 3)
INSERT [dbo].[Cart] ([UserID], [ProductID], [Quantity]) VALUES (1, 4, 11)
INSERT [dbo].[Cart] ([UserID], [ProductID], [Quantity]) VALUES (1, 5, 17)
GO
SET IDENTITY_INSERT [dbo].[Categories] ON 

INSERT [dbo].[Categories] ([CategoryID], [Name]) VALUES (1, N'Quần jeans Hàng hiệu loại 1')
INSERT [dbo].[Categories] ([CategoryID], [Name]) VALUES (2, N'Quần jeans Hàng hiệu loại 1')
INSERT [dbo].[Categories] ([CategoryID], [Name]) VALUES (3, N'Áo khoác')
INSERT [dbo].[Categories] ([CategoryID], [Name]) VALUES (4, N'Quần thể thao')
INSERT [dbo].[Categories] ([CategoryID], [Name]) VALUES (5, N'Áo len')
SET IDENTITY_INSERT [dbo].[Categories] OFF
GO
SET IDENTITY_INSERT [dbo].[OrderDetails] ON 

INSERT [dbo].[OrderDetails] ([OrderDetailID], [OrderID], [ProductID], [Quantity], [Price]) VALUES (1, 1, 1, 2, CAST(29.99 AS Decimal(10, 2)))
INSERT [dbo].[OrderDetails] ([OrderDetailID], [OrderID], [ProductID], [Quantity], [Price]) VALUES (2, 2, 2, 1, CAST(39.99 AS Decimal(10, 2)))
INSERT [dbo].[OrderDetails] ([OrderDetailID], [OrderID], [ProductID], [Quantity], [Price]) VALUES (3, 3, 3, 3, CAST(49.99 AS Decimal(10, 2)))
INSERT [dbo].[OrderDetails] ([OrderDetailID], [OrderID], [ProductID], [Quantity], [Price]) VALUES (4, 4, 4, 2, CAST(19.99 AS Decimal(10, 2)))
INSERT [dbo].[OrderDetails] ([OrderDetailID], [OrderID], [ProductID], [Quantity], [Price]) VALUES (5, 5, 5, 4, CAST(34.99 AS Decimal(10, 2)))
SET IDENTITY_INSERT [dbo].[OrderDetails] OFF
GO
SET IDENTITY_INSERT [dbo].[Orders] ON 

INSERT [dbo].[Orders] ([OrderID], [UserID], [OrderDate], [Status], [Address]) VALUES (1, 1, CAST(N'2023-10-27T10:11:57.473' AS DateTime), N'delivering', N'44 Hoàng hoa thám ')
INSERT [dbo].[Orders] ([OrderID], [UserID], [OrderDate], [Status], [Address]) VALUES (2, 1, CAST(N'2023-10-27T10:11:57.473' AS DateTime), N'canceled', NULL)
INSERT [dbo].[Orders] ([OrderID], [UserID], [OrderDate], [Status], [Address]) VALUES (3, 1, CAST(N'2023-10-27T10:11:57.473' AS DateTime), N'canceled', NULL)
INSERT [dbo].[Orders] ([OrderID], [UserID], [OrderDate], [Status], [Address]) VALUES (4, 1, CAST(N'2023-10-27T10:11:57.473' AS DateTime), N'Đã đặt hàng', NULL)
INSERT [dbo].[Orders] ([OrderID], [UserID], [OrderDate], [Status], [Address]) VALUES (5, 1, CAST(N'2023-10-27T10:11:57.473' AS DateTime), N'canceled', NULL)
INSERT [dbo].[Orders] ([OrderID], [UserID], [OrderDate], [Status], [Address]) VALUES (8, 1, NULL, N'confirm', NULL)
INSERT [dbo].[Orders] ([OrderID], [UserID], [OrderDate], [Status], [Address]) VALUES (9, 1, NULL, N'confirm', NULL)
INSERT [dbo].[Orders] ([OrderID], [UserID], [OrderDate], [Status], [Address]) VALUES (10, 1, NULL, N'confirm', NULL)
INSERT [dbo].[Orders] ([OrderID], [UserID], [OrderDate], [Status], [Address]) VALUES (11, 1, NULL, N'confirm', NULL)
INSERT [dbo].[Orders] ([OrderID], [UserID], [OrderDate], [Status], [Address]) VALUES (13, 1, CAST(N'2006-01-02T15:04:05.000' AS DateTime), N'confirm', NULL)
INSERT [dbo].[Orders] ([OrderID], [UserID], [OrderDate], [Status], [Address]) VALUES (14, 1, CAST(N'2006-01-02T15:04:05.000' AS DateTime), N'confirm', NULL)
INSERT [dbo].[Orders] ([OrderID], [UserID], [OrderDate], [Status], [Address]) VALUES (15, 1, CAST(N'2006-01-02T15:04:05.000' AS DateTime), N'confirm', NULL)
INSERT [dbo].[Orders] ([OrderID], [UserID], [OrderDate], [Status], [Address]) VALUES (16, 1, CAST(N'2006-01-02T15:04:05.000' AS DateTime), N'confirm', NULL)
INSERT [dbo].[Orders] ([OrderID], [UserID], [OrderDate], [Status], [Address]) VALUES (17, 1, CAST(N'2006-01-02T15:04:05.000' AS DateTime), N'delivering', NULL)
INSERT [dbo].[Orders] ([OrderID], [UserID], [OrderDate], [Status], [Address]) VALUES (18, 1, CAST(N'2006-01-02T15:04:05.000' AS DateTime), N'delivering', NULL)
INSERT [dbo].[Orders] ([OrderID], [UserID], [OrderDate], [Status], [Address]) VALUES (19, 1, CAST(N'2006-01-02T15:04:05.000' AS DateTime), N'confirm', NULL)
INSERT [dbo].[Orders] ([OrderID], [UserID], [OrderDate], [Status], [Address]) VALUES (20, 1, CAST(N'2006-01-02T15:04:05.000' AS DateTime), N'confirm', NULL)
SET IDENTITY_INSERT [dbo].[Orders] OFF
GO
SET IDENTITY_INSERT [dbo].[Products] ON 

INSERT [dbo].[Products] ([ProductID], [Name], [Description], [Price], [CategoryID]) VALUES (1, N'Áo loong heheheh', N'Áo khoác nam dáng dài dàidàidàidàidàidàidàidàidàidàidàidàidàidàimàu đenđenmmmmmđenmmmmm', 109.99, 2)
INSERT [dbo].[Products] ([ProductID], [Name], [Description], [Price], [CategoryID]) VALUES (2, N'Quần jeans đen', N'Quần jeans nam màu đen', 39.99, 2)
INSERT [dbo].[Products] ([ProductID], [Name], [Description], [Price], [CategoryID]) VALUES (3, N'Áo khoác nỉ đen', N'Áo khoác nam dáng dài màu đen', 49.99, 3)
INSERT [dbo].[Products] ([ProductID], [Name], [Description], [Price], [CategoryID]) VALUES (4, N'Quần thể thao xanh', N'Quần thể thao nam màu xanh', 19.99, 4)
INSERT [dbo].[Products] ([ProductID], [Name], [Description], [Price], [CategoryID]) VALUES (5, N'Áo len cổ lọ xám', N'Áo len nam cổ lọ màu xám', 34.99, 5)
INSERT [dbo].[Products] ([ProductID], [Name], [Description], [Price], [CategoryID]) VALUES (6, N'Áo khoác nỉ đen', N'Áo khoác nam dáng dài màu đen', 49.99, 3)
INSERT [dbo].[Products] ([ProductID], [Name], [Description], [Price], [CategoryID]) VALUES (8, N'Áo khoác nỉ đenđennnnnnnđennnnnnn', N'Áo khoác nam dáng dài màu đenđenmmmmmđenmmmmm', 99.99, 4)
INSERT [dbo].[Products] ([ProductID], [Name], [Description], [Price], [CategoryID]) VALUES (9, N'Áo khoác nỉ đenđennnnnnnđennnnnnn', N'Áo khoác nam dáng dài màu đenđenmmmmmđenmmmmm', 99.99, 4)
INSERT [dbo].[Products] ([ProductID], [Name], [Description], [Price], [CategoryID]) VALUES (10, N'Áo khoác nỉ đenđennnnnnnđennnnnnn', N'Áo khoác nam dáng dài màu đenđenmmmmmđenmmmmm', 99.99, 4)
INSERT [dbo].[Products] ([ProductID], [Name], [Description], [Price], [CategoryID]) VALUES (11, N'Áo khoác nỉ đenđennnnnnnđennnnnnn', N'Áo khoác nam dáng dài màu đenđenmmmmmđenmmmmm', 99.99, 4)
INSERT [dbo].[Products] ([ProductID], [Name], [Description], [Price], [CategoryID]) VALUES (13, N'Áo loong cuuu nnnnnnn', N'Áo khoác nam dáng dài màu đenđenmmmmmđenmmmmm', 99.99, 4)
INSERT [dbo].[Products] ([ProductID], [Name], [Description], [Price], [CategoryID]) VALUES (15, N'Áo loong cuuu nnnnnnn', N'Áo khoác nam dáng dài màu đenđenmmmmmđenmmmmm', 99.99, 4)
INSERT [dbo].[Products] ([ProductID], [Name], [Description], [Price], [CategoryID]) VALUES (16, N'Áo loong cuuu nnnnnnn', N'Áo khoác nam dáng dài màu đenđenmmmmmđenmmmmm', 99.99, 4)
INSERT [dbo].[Products] ([ProductID], [Name], [Description], [Price], [CategoryID]) VALUES (17, N'Áo loong cuuu nnnnnnn', N'Áo khoác nam dáng dài màu đenđenmmmmmđenmmmmm', 99.99, 4)
SET IDENTITY_INSERT [dbo].[Products] OFF
GO
SET IDENTITY_INSERT [dbo].[Users] ON 

INSERT [dbo].[Users] ([UserID], [Email], [Password], [Role], [FullName], [Address]) VALUES (1, N'john@example.com', N'yenhi123', N'user', N'Nguyễn Ngọc Yến Nhi', N'40 á đù  Võ Chí Công')
INSERT [dbo].[Users] ([UserID], [Email], [Password], [Role], [FullName], [Address]) VALUES (2, N'jane@example.com', N'hashed_password', N'user', N'Trần Văn Hùng', NULL)
INSERT [dbo].[Users] ([UserID], [Email], [Password], [Role], [FullName], [Address]) VALUES (3, N'admin@example.com', N'hashed_password', N'admin', N'Lê Thị Thu', NULL)
INSERT [dbo].[Users] ([UserID], [Email], [Password], [Role], [FullName], [Address]) VALUES (4, N'sarah@example.com', N'hashed_password', N'user', N'Phạm Minh Đức', NULL)
INSERT [dbo].[Users] ([UserID], [Email], [Password], [Role], [FullName], [Address]) VALUES (5, N'david@example.com', N'hashed_password', N'user', N'Hoàng Thị Nga', NULL)
INSERT [dbo].[Users] ([UserID], [Email], [Password], [Role], [FullName], [Address]) VALUES (26, N'abc@123gmail.com', N'hashed_passwordpasswordddpassworddd', N'user', N'32r433r3', NULL)
INSERT [dbo].[Users] ([UserID], [Email], [Password], [Role], [FullName], [Address]) VALUES (30, N'abc@123gm123gm1123gm1ail.com', N'hashed_passwordpasswordddpassworddd', N'user', N'32r433r3', NULL)
INSERT [dbo].[Users] ([UserID], [Email], [Password], [Role], [FullName], [Address]) VALUES (31, N'abab1ab1c@123gmail.com', N'hashed_passwordpasswordddpassworddd', N'user', N'32r433r3', NULL)
SET IDENTITY_INSERT [dbo].[Users] OFF
GO
/****** Object:  Index [unique_UserID_ProductID]    Script Date: 11/2/2023 12:35:45 PM ******/
ALTER TABLE [dbo].[Cart] ADD  CONSTRAINT [unique_UserID_ProductID] UNIQUE NONCLUSTERED 
(
	[UserID] ASC,
	[ProductID] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, SORT_IN_TEMPDB = OFF, IGNORE_DUP_KEY = OFF, ONLINE = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
GO
SET ANSI_PADDING ON
GO
/****** Object:  Index [UC_Email]    Script Date: 11/2/2023 12:35:45 PM ******/
ALTER TABLE [dbo].[Users] ADD  CONSTRAINT [UC_Email] UNIQUE NONCLUSTERED 
(
	[Email] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, SORT_IN_TEMPDB = OFF, IGNORE_DUP_KEY = OFF, ONLINE = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
GO
ALTER TABLE [dbo].[Cart]  WITH CHECK ADD FOREIGN KEY([ProductID])
REFERENCES [dbo].[Products] ([ProductID])
GO
ALTER TABLE [dbo].[Cart]  WITH CHECK ADD FOREIGN KEY([UserID])
REFERENCES [dbo].[Users] ([UserID])
GO
ALTER TABLE [dbo].[OrderDetails]  WITH CHECK ADD FOREIGN KEY([OrderID])
REFERENCES [dbo].[Orders] ([OrderID])
GO
ALTER TABLE [dbo].[OrderDetails]  WITH CHECK ADD FOREIGN KEY([ProductID])
REFERENCES [dbo].[Products] ([ProductID])
GO
ALTER TABLE [dbo].[Orders]  WITH CHECK ADD FOREIGN KEY([UserID])
REFERENCES [dbo].[Users] ([UserID])
GO
ALTER TABLE [dbo].[Products]  WITH CHECK ADD FOREIGN KEY([CategoryID])
REFERENCES [dbo].[Categories] ([CategoryID])
GO
ALTER TABLE [dbo].[Reviews]  WITH CHECK ADD FOREIGN KEY([ProductID])
REFERENCES [dbo].[Products] ([ProductID])
GO
ALTER TABLE [dbo].[Reviews]  WITH CHECK ADD FOREIGN KEY([UserID])
REFERENCES [dbo].[Users] ([UserID])
GO
USE [master]
GO
ALTER DATABASE [Shop] SET  READ_WRITE 
GO
