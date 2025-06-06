package model

// Perform database migration

func migration() {
	// Auto migration mode
	_ = DB.AutoMigrate(&User{})
	_ = DB.AutoMigrate(&Video{})
}
