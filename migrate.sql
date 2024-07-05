-- Create table Rooms
CREATE TABLE Rooms (
    RoomID INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    CategoryID INT,
    EmployeeID INT,
    Capacity INT,
    Description TEXT,
    FOREIGN KEY (CategoryID) REFERENCES Categories (CategoryID),
    FOREIGN KEY (EmployeeID) REFERENCES Employees (EmployeeID)
);

-- Create table Guests
CREATE TABLE Guests (
    GuestID INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    Lastname VARCHAR(50),
    Firstname VARCHAR(50),
    Patronymic VARCHAR(50),
    PassportNumber VARCHAR(20),
    BirthDate DATE,
    CityOfOrigin VARCHAR(100)
);

-- Create table Categories
CREATE TABLE Categories (
    CategoryID INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    Name VARCHAR(100),
    RoomCount INT,
    PricePerDay DECIMAL(10, 2)
);

-- Create table OccupiedRooms
CREATE TABLE OccupiedRooms (
    ID INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    RoomID INT,
    GuestID INT,
    ServiceID INT,
    CheckInDate DATE,
    CheckOutDate DATE,
    AccommodationCost DECIMAL(10, 2),
    FOREIGN KEY (RoomID) REFERENCES Rooms (RoomID),
    FOREIGN KEY (GuestID) REFERENCES Guests (GuestID),
    FOREIGN KEY (ServiceID) REFERENCES AdditionalServices (ServiceID)
);

-- Create table AdditionalServices
CREATE TABLE AdditionalServices (
    ServiceID INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    ServiceName VARCHAR(100),
    Price DECIMAL(10, 2)
);

-- Create table PaymentMethods
CREATE TABLE PaymentMethods (
    ID INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    GuestID INT,
    PaymentForm VARCHAR(50),
    PaidAmount DECIMAL(10, 2),
    FOREIGN KEY (GuestID) REFERENCES Guests (GuestID)
);

-- Create table Employees
CREATE TABLE Employees (
    EmployeeID INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    Lastname VARCHAR(50),
    Firstname VARCHAR(50),
    Patronymic VARCHAR(50),
    BirthDate DATE,
    PositionID INT,
    FOREIGN KEY (PositionID) REFERENCES Positions (PositionID)
);

-- Create table Positions
CREATE TABLE Positions (
    PositionID INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    Title VARCHAR(100),
    Salary DECIMAL(10, 2)
);

-- Create table Children
CREATE TABLE Children (
    ChildID INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    GuestID INT,
    Lastname VARCHAR(50),
    Firstname VARCHAR(50),
    Patronymic VARCHAR(50),
    BirthDate DATE,
    FOREIGN KEY (GuestID) REFERENCES Guests (GuestID)
);