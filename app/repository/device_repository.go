package repository

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/fernandamelov/device-management-api/app/models"
	_ "github.com/lib/pq"
)

var db *sql.DB

func InitializeDatabase() {
	var err error
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))

	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	// Criar a tabela se nao existir
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS devices (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		brand VARCHAR(255) NOT NULL,
		creation_time TIMESTAMP
	);`

	_, err = db.Exec(createTableQuery)
	if err != nil {
		log.Fatal(err)
	}
}

func AddDevice(device models.Device) (int64, error) {
	var id int64
	err := db.QueryRow("INSERT INTO devices (name, brand, creation_time) VALUES ($1, $2, $3) RETURNING id", device.Name, device.Brand, time.Now()).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func GetDevice(id int64) (models.Device, error) {
	var device models.Device
	err := db.QueryRow("SELECT id, name, brand FROM devices WHERE id = $1", id).Scan(&device.ID, &device.Name, &device.Brand)
	return device, err
}

func ListDevices() ([]models.Device, error) {
	rows, err := db.Query("SELECT id, name, brand FROM devices")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var devices []models.Device
	for rows.Next() {
		var device models.Device
		if err := rows.Scan(&device.ID, &device.Name, &device.Brand); err != nil {
			return nil, err
		}
		devices = append(devices, device)
	}
	return devices, nil
}

func UpdateDevice(id int64, device models.Device) error {
	_, err := db.Exec("UPDATE devices SET name = $1, brand = $2, creation_time = $3 WHERE id = $4", device.Name, device.Brand, time.Now(), id)
	return err
}

func PartialUpdateDevice(id int64, device models.Device) error {
	query := "UPDATE devices SET "
	params := []interface{}{}
	i := 1
	if device.Name != "" {
		query += fmt.Sprintf("name = $%d, ", i)
		params = append(params, device.Name)
		i++
	}
	if device.Brand != "" {
		query += fmt.Sprintf("brand = $%d, ", i)
		params = append(params, device.Brand)
		i++
	}

	query = query[:len(query)-2] // removes the last comma and space

	query += fmt.Sprintf(" WHERE id = $%d", i)
	params = append(params, id)
	_, err := db.Exec(query, params...)
	return err
}

func DeleteDevice(id int64) error {
	_, err := db.Exec("DELETE FROM devices WHERE id = $1", id)
	return err
}

func SearchDevicesByBrand(brand string) ([]models.Device, error) {
	rows, err := db.Query("SELECT id, name, brand FROM devices WHERE brand = $1", brand)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var devices []models.Device
	for rows.Next() {
		var device models.Device
		if err := rows.Scan(&device.ID, &device.Name, &device.Brand); err != nil {
			return nil, err
		}
		devices = append(devices, device)
	}
	return devices, nil
}
