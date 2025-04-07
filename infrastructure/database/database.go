package database

import (
	"fmt"
	"log"

	"github.com/LucioSchiavoni/scan-host/config"
	"github.com/LucioSchiavoni/scan-host/infrastructure/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DBUser, config.DBPass, config.DBHost, config.DBName)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error al conectar con la base de datos:", err)
	}

	err = DB.AutoMigrate(&models.Equipo{}, models.Aplicacion{}, models.EquipoApp{})
	if err != nil {
		log.Fatal("Error al migrar la tabla equipos:", err)
	}

	log.Println("✅ Conexión a la base de datos establecida")
	log.Println("✅ Tabla equipos migrada correctamente")
}
