USE [Homework2]
GO

SET ANSI_NULLS ON
GO

SET QUOTED_IDENTIFIER ON
GO

CREATE PROCEDURE [dbo].[SP_CustomerManagement]
	@Process NVARCHAR(30) = NULL,
	@Id INT = -1,
	@FirstName NVARCHAR(80) = NULL,
	@LastName NVARCHAR(80) = NULL,
	@Address NVARCHAR(160) = NULL,
	@Birthday INT = NULL

AS
BEGIN
	IF @Process = 'CREATE' 
		BEGIN
			INSERT INTO TB_Customer(FirstName,LastName, Address, Birthday)
			VALUES (@FirstName,@LastName, @Address, CONVERT(date,@Birthday))
		END

	ELSE IF @Process = 'UPDATE'
		BEGIN
			UPDATE TB_Customer
			SET FirstName = @FirstName,
				LastName = @LastName,
				Address = @Address,
				Birthday = CONVERT(date,@Birthday)
			WHERE Id = @Id
		END

	ELSE IF @Process = 'DELETE'
		BEGIN
			UPDATE TB_Customer
			SET Delflag = 1
			WHERE Id = @Id
		END

	ELSE IF @Process = 'FIND'
		BEGIN
			SELECT  c.Id,
					c.FirstName,
					c.LastName,
					c.Address,
					c.Birthday
			FROM TB_Customer c
			WHERE Id = @Id
			AND Delflag = 0
		END
	ELSE 
		BEGIN
			SELECT  c.Id,
					c.FirstName,
					c.LastName,
					c.Address,
					c.Birthday
			FROM TB_Customer c
			WHERE Delflag = 0
		END
END
GO


