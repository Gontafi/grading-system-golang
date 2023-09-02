package services

import (
	"encoding/json"
	"errors"
	"github.com/grading-system-golang/internal/models"
	"github.com/redis/go-redis/v9"
	"time"
)

const (
	WeekDuration  = 7 * 24 * time.Hour
	MonthDuration = 30 * 24 * time.Hour
	YearDuration  = 365 * 24 * time.Hour
)

func (s *ServiceV1) GetTopRatingFromCache(period string, limit int) ([]models.Rating, error) {

	var periodTime time.Duration

	switch period {
	case "week":
		periodTime = WeekDuration
	case "month":
		periodTime = MonthDuration
	case "year":
		periodTime = YearDuration
	default:
		periodTime = 0
	}

	if periodTime == 0 {
		return []models.Rating{}, errors.New("invalid period time")
	}

	cacheKey := s.getCacheKey("topRating", period, limit)

	data, err := s.rdb.Get(s.ctx, cacheKey).Bytes()
	if err == nil {
		var ratings []models.Rating
		if err := json.Unmarshal(data, &ratings); err != nil {
			return nil, err
		}
		return ratings, nil
	} else if err != redis.Nil {
		return nil, err
	}

	ratings, err := s.repository.GetTopRating(periodTime, limit)
	if err != nil {
		return nil, err
	}

	ratingsJSON, err := json.Marshal(ratings)
	if err != nil {
		return nil, err
	}
	if err := s.rdb.Set(s.ctx, cacheKey, ratingsJSON, s.expiry).Err(); err != nil {
		return nil, err
	}

	return ratings, nil
}

func (s *ServiceV1) GetTopRatingByLessonFromCache(
	lessonID int,
	period time.Duration,
	limit int) ([]models.Rating, error) {

	cacheKey := s.getCacheKey("topRatingByLesson", lessonID, period, limit)

	data, err := s.rdb.Get(s.ctx, cacheKey).Bytes()
	if err == nil {
		var ratings []models.Rating
		if err := json.Unmarshal(data, &ratings); err != nil {
			return nil, err
		}
		return ratings, nil
	} else if err != redis.Nil {
		return nil, err
	}

	ratings, err := s.repository.GetTopRatingByLesson(lessonID, period, limit)
	if err != nil {
		return nil, err
	}

	ratingsJSON, err := json.Marshal(ratings)
	if err != nil {
		return nil, err
	}
	if err := s.rdb.Set(s.ctx, cacheKey, ratingsJSON, s.expiry).Err(); err != nil {
		return nil, err
	}

	return ratings, nil
}
