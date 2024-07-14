package function

import(
	"database/sql"
	en "enigma-laundry/entity"
	"fmt"
)

//Service function
func InsertService(db *sql.DB, service en.Service) {
	if err := service.Validate(); err != nil {
		fmt.Println("Validation error:", err)
		return
	}

	query := "INSERT INTO service (nama, satuan, harga) VALUES ($1, $2, $3)"
	_, err := db.Exec(query, service.ServiceName, service.Satuan, service.Price)
	if err != nil {
		fmt.Println("Insert error:", err)
	} else {
		fmt.Println("Service successfully inserted")
	}
}

func ViewServices(db *sql.DB) {
	rows, err := db.Query("SELECT nama, satuan, harga FROM service")
	if err != nil {
		fmt.Println("View error:", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var service en.Service
		if err := rows.Scan(&service.ServiceName, &service.Satuan, &service.Price); err != nil {
			fmt.Println("Scan error:", err)
		} else {
			fmt.Printf("Name: %s, Satuan: %s, Price: %.2f\n", service.ServiceName, service.Satuan, service.Price)
		}
	}
}

func UpdateService(db *sql.DB, name string, service en.Service) {
	if err := service.Validate(); err != nil {
		fmt.Println("Validation error:", err)
		return
	}

	query := "UPDATE service SET nama = $2, satuan = $3, harga = $4 WHERE nama = $1"
	_, err := db.Exec(query, name, service.ServiceName, service.Satuan, service.Price)
	if err != nil {
		fmt.Println("Update error:", err)
	} else {
		fmt.Println("Service successfully updated")
	}
}

func DeleteService(db *sql.DB, name string) {
	query := "DELETE FROM service WHERE nama = $1"
	_, err := db.Exec(query, name)
	if err != nil {
		fmt.Println("Delete error:", err)
	} else {
		fmt.Println("Service successfully deleted")
	}
}