package dbconnect

import . "education.api/entities"

func InitialMigration() {
	connection := DbInit()
	defer CloseDatabase(connection)
	connection.AutoMigrate(&User{}, &Branch{}, &Announcement{}, &Class{}, &Lesson{}, &StudentPayment{})
}
