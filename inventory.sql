/*
This file contains a script of Transact SQL (T-SQL) command to interact with a database 'Inventory'.
Requirements:
- SQL Server 2022 is installed and running
- database 'Inventory' already exists.
Details:
- Checks if the database 'Inventory' exists. If not, exit with an error message.
- Sets the default database to 'Inventory'.
- Creates a 'categories' table and related 'products' table if they do not already exist.
- Remove all rows from the 'products' and 'categories' tables
- Populates the 'Categories' table with sample data.
- Populates the 'Products' table with sample data.
- Create a stored procedure to get all categories.
*/

-- Check if the database 'Inventory' exists 
IF NOT EXISTS (SELECT name FROM master.dbo.sysdatabases WHERE name = N'Inventory')
BEGIN
    PRINT 'The database [Inventory] does not exist. Please create the database and run this script again.'
    RETURN
END

-- Set the default database to 'Inventory'
USE Inventory
GO

-- Create a 'Categories' table if it does not already exist
IF NOT EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[Categories]') AND type in (N'U'))
BEGIN
    CREATE TABLE [dbo].[Categories]
    (
        [CategoryId] INT PRIMARY KEY,
        [CategoryName] NVARCHAR(50) NOT NULL
    )
END

-- Create a 'Products' table if it does not already exist
IF NOT EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[Products]') AND type in (N'U'))
BEGIN
    CREATE TABLE [dbo].[Products]
    (
        [ProductId] INT PRIMARY KEY,
        [ProductName] NVARCHAR(50) NOT NULL,
        [UnitPrice] DECIMAL(10, 2) NOT NULL,
        [CategoryId] INT FOREIGN KEY REFERENCES Categories(CategoryId),
        -- Add a created date column
        [CreatedDate] DATETIME DEFAULT GETDATE(),
        -- Add a modified date column
        [ModifiedDate] DATETIME DEFAULT GETDATE()
    )
END

-- Remove all rows from the 'Products' and 'Categories' tables
TRUNCATE TABLE [dbo].[Products]
TRUNCATE TABLE [dbo].[Categories]

-- Populate the 'Categories' table with sample data
INSERT INTO [dbo].[Categories] ([CategoryId], [CategoryName])
VALUES
(1, 'Electronics'),
(2, 'Clothing'),
(3, 'Books')

-- Populate the 'Products' table with sample data
INSERT INTO [dbo].[Products] ([ProductId], [ProductName], [UnitPrice], [CategoryId])
VALUES
(1, 'Laptop', 999.99, 1),
(2, 'T-shirt', 19.99, 2),
(3, 'Book', 14.99, 3)


-- Create a stored procedure to get all categories
IF OBJECT_ID('GetAllCategories') IS NOT NULL
BEGIN
    DROP PROCEDURE GetAllCategories
END

CREATE PROCEDURE GetAllCategories
AS
BEGIN
    SELECT * FROM [dbo].[Categories]
END


-- Display a success message
PRINT 'The inventory database has been initialized with sample data.'
