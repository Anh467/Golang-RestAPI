-- 1 6:30
ALTER TABLE OrderDetails
DROP COLUMN Price;

-- 2 6:40
ALTER TABLE OrderDetails
DROP COLUMN OrderDetailID;

-- 3 6:41
ALTER TABLE OrderDetails
ADD CONSTRAINT constraint_unique_OrderID_ProductID UNIQUE (OrderID, ProductID);