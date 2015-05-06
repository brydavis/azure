if exists (SELECT 1 FROM sys.tables WHERE object_id = OBJECT_ID('Persons')) BEGIN; DROP TABLE Persons; END;

 -- GO
create table Persons ( PersonsID integer NOT NULL identity(1, 1), Id varchar(40) NULL, FirstName varchar(255) NULL, LastName varchar(255) NULL, Birthdate datetime NULL, SSN varchar(255) NULL, Street varchar(255) NULL, City varchar(255) NULL, Zip varchar(10) NULL, State varchar(50) NULL, Race varchar(255) NULL, Latino varchar(255) NULL, Gender varchar(255) NULL, Veteran varchar(255) NULL, MaritialStatus varchar(255) NULL, DisabilityStatus varchar(255) NULL, EducationLevel varchar(255) NULL, EmploymentStatus varchar(255) NULL, PRIMARY KEY (PersonsID));
