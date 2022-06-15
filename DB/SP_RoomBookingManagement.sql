USE [HOMEWORK2]
GO

/****** Object:  StoredProcedure [dbo].[SP_RoomBookingManagement]    Script Date: 15-06-22 3:36:58 PM ******/
SET ANSI_NULLS ON
GO

SET QUOTED_IDENTIFIER ON
GO

CREATE PROCEDURE [dbo].[SP_RoomBookingManagement]
	@Process NVARCHAR(30) = NULL,
	@Id INT = -1,
	@ReserveDate date = NULL,
	@ReserveStartTime time = NULL,
	@ReserveEndTime time = NULL,
	@RoomNo varchar (10) = NULL
AS
BEGIN
	IF @Process = 'CREATE' 
		BEGIN
			INSERT INTO TB_RoomBooking(ReserveDate, ReserveStartTime,ReserveEndTime,RoomNo)
			VALUES (@ReserveDate, @ReserveStartTime,@ReserveEndTime,@RoomNo)
		END
	ELSE IF @Process = 'UPDATE'
		BEGIN
			UPDATE TB_RoomBooking
			SET ReserveDate = @ReserveDate,
				ReserveStartTime = @ReserveStartTime,
				ReserveEndTime = @ReserveEndTime,
				RoomNo = @RoomNo
			WHERE Id = @Id
		END
	ELSE IF @Process = 'DELETE'
		BEGIN
			UPDATE TB_RoomBooking
			SET Delflag = 1
			WHERE Id = @Id
		END
	ELSE IF @Process = 'FIND'
		BEGIN
			SELECT  r.Id,
					r.ReserveDate,
					r.ReserveStartTime,
					r.ReserveEndTime,
					r.RoomNo
			FROM TB_RoomBooking r
			WHERE Id = @Id
			AND Delflag = 0
		END
	ELSE 
		BEGIN
			SELECT  r.Id,
					r.ReserveDate,
					r.ReserveStartTime,
					r.ReserveEndTime,
					r.RoomNo
			FROM TB_RoomBooking r
			WHERE Delflag = 0
		END
END
GO


