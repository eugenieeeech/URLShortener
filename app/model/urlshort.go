package model

func GetAllurlshorts() ([]URLshort, error) {
	var urlshorts []URLshort

	tx := db.Find(&urlshorts)
	if tx.Error != nil {
		return []URLshort{}, tx.Error
	}

	return urlshorts, nil
}

func Geturlshorts(id uint64) (URLshort, error) {
	var urlshort URLshort

	tx := db.Where("id = ?", id).First(&urlshort)

	if tx.Error != nil {
		return URLshort{}, tx.Error
	}

	return urlshort, nil
}

func CreateURLshort(urlshort URLshort) error {
	tx := db.Create(&urlshort)
	return tx.Error
}

func UpdateURLshort(urlshort URLshort) error {
	tx := db.Save(&urlshort)
	return tx.Error
}

func DeleteURLshort(id uint64) error {
	tx := db.Unscoped().Delete(&URLshort{}, id)
	return tx.Error
}

func FindByURLshortUrl(url string) (URLshort, error) {
	var urlshort URLshort
	tx := db.Where("urlshort = ?", url).First(&urlshort)
	return urlshort, tx.Error
}
