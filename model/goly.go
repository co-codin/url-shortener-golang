package model

func GetAllGolies() ([]Goly, error) {
	var golies []Goly

	tx := db.Find(&golies)

	if tx.Error != nil {
		return []Goly{}, tx.Error
	}

	return golies, nil
} 

func GetGoly(id uint64) (Goly, error) {
	var goly Goly

	tx := db.Where("id = ?", id).First(&goly)

	if tx.Error != nil {
		return Goly{}, tx.Error
	}

	return goly, nil
}