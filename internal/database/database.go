package database

import (
	"context"
	"database/sql"

	"github.com/YerzatCode/sql_/internal/config"
	_ "github.com/mattn/go-sqlite3"
)

func DatabaseInit(dbc *config.Database) *sql.DB {
	db, err := sql.Open(dbc.DB, dbc.Path)
	if err != nil {
		panic(err)
	}

	CreateTable(db)

	// defer db.Close()

	return db
}

func CreateTable(db *sql.DB) {
	stmt, err := db.PrepareContext(context.Background(), `CREATE TABLE IF NOT EXISTS Rooms (
        RoomID INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        CategoryID INT,
        EmployeeID INT,
        Capacity INT,
        Description TEXT,
        FOREIGN KEY (CategoryID) REFERENCES Categories (CategoryID),
        FOREIGN KEY (EmployeeID) REFERENCES Employees (EmployeeID)
    );`)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		panic(err)
	}

	stmt, err = db.PrepareContext(context.Background(), `CREATE TABLE IF NOT EXISTS Guests (
        GuestID INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        Lastname VARCHAR(50),
        Firstname VARCHAR(50),
        Patronymic VARCHAR(50),
        PassportNumber VARCHAR(20),
        BirthDate DATE,
        CityOfOrigin VARCHAR(100)
    );`)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		panic(err)
	}

	// Create table Categories
	stmt, err = db.PrepareContext(context.Background(), `CREATE TABLE IF NOT EXISTS Categories (
        CategoryID INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        Name VARCHAR(100),
        RoomCount INT,
        PricePerDay DECIMAL(10, 2)
    );`)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		panic(err)
	}

	// Create table OccupiedRooms
	stmt, err = db.PrepareContext(context.Background(), `CREATE TABLE IF NOT EXISTS OccupiedRooms (
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
    );`)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		panic(err)
	}

	// Create table AdditionalServices
	stmt, err = db.PrepareContext(context.Background(), `CREATE TABLE IF NOT EXISTS AdditionalServices (
        ServiceID INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        ServiceName VARCHAR(100),
        Price DECIMAL(10, 2)
    );`)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		panic(err)
	}

	// Create table PaymentMethods
	stmt, err = db.PrepareContext(context.Background(), `CREATE TABLE IF NOT EXISTS PaymentMethods (
        ID INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        GuestID INT,
        PaymentForm VARCHAR(50),
        PaidAmount DECIMAL(10, 2),
        FOREIGN KEY (GuestID) REFERENCES Guests (GuestID)
    );`)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		panic(err)
	}

	// Create table Employees
	stmt, err = db.PrepareContext(context.Background(), `CREATE TABLE IF NOT EXISTS Employees (
        EmployeeID INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        Lastname VARCHAR(50),
        Firstname VARCHAR(50),
        Patronymic VARCHAR(50),
        BirthDate DATE,
        PositionID INT,
        FOREIGN KEY (PositionID) REFERENCES Positions (PositionID)
    );`)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		panic(err)
	}

	// Create table Positions
	stmt, err = db.PrepareContext(context.Background(), `CREATE TABLE IF NOT EXISTS Positions (
        PositionID INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        Title VARCHAR(100),
        Salary DECIMAL(10, 2)
    );`)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		panic(err)
	}

	// Create table Children
	stmt, err = db.PrepareContext(context.Background(), `CREATE TABLE IF NOT EXISTS Children (
        ChildID INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        GuestID INT,
        Lastname VARCHAR(50),
        Firstname VARCHAR(50),
        Patronymic VARCHAR(50),
        BirthDate DATE,
        FOREIGN KEY (GuestID) REFERENCES Guests (GuestID)
    );`)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		panic(err)
	}
}
