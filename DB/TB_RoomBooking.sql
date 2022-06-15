USE [Homework2]
GO

/****** Object:  Table [dbo].[TB_Customer]    Script Date: 23-Nov-21 3:48:11 PM ******/
SET ANSI_NULLS ON
GO

SET QUOTED_IDENTIFIER ON
GO

CREATE TABLE [dbo].[TB_RoomBooking](
	[Id] [int] IDENTITY(1,1) NOT NULL,
	[ReserveDate] [date] NULL,
	[ReserveStartTime] [time] NULL,
	[ReserveEndTime] [time] NULL,
	[RoomNo][varchar](10) NULL,
	[Delflag] [int] NULL,
 CONSTRAINT [PK_TB_RoomBooking] PRIMARY KEY CLUSTERED 
(
	[Id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
) ON [PRIMARY]
GO

ALTER TABLE [dbo].[TB_RoomBooking] ADD  CONSTRAINT [DF_TB_RoomBooking_Delflag]  DEFAULT ((0)) FOR [Delflag]
GO


