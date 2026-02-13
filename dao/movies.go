package dao

import (
	"errors"
	"time"

	models "github.com/micheldevs/florobot/models"
	"gorm.io/gorm"
)

func InsertMovie(movie *models.Movie) error {
	if result := DB.Create(&movie); result.Error != nil {
		return result.Error
	}

	return nil
}

func GetMoviesFrom(fromDate time.Time, toDate time.Time) ([]models.Movie, error) {
	var movies []models.Movie
	if result := DB.Order("premiere_date asc").Where("premiere_date >= ?", fromDate).Where("premiere_date <= ?", toDate).Find(&movies); result.Error != nil {
		return movies, result.Error
	}
	return movies, nil
}

func GetMovieByExtId(extId string) (models.Movie, error) {
	var movie models.Movie
	if result := DB.Where("ext_id = ?", extId).First(&movie); result.Error != nil {
		return movie, result.Error
	}
	return movie, nil
}

func ExistMovie(extId string) (bool, error) {
	var movie models.Movie
	if result := DB.Where("ext_id = ?", extId).First(&movie); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, nil
	}
	return true, nil
}
