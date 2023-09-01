package services

import (
	"encoding/json"
	"fmt"
	"github.com/grading-system-golang/internal/models"
	"github.com/redis/go-redis/v9"
	"strconv"
	"strings"
	"time"
)

const (
	WeekDuration  = 7 * 24 * time.Hour
	MonthDuration = 30 * 24 * time.Hour
	YearDuration  = 365 * 24 * time.Hour
)

func (s *ServiceV1) GetTopRatingFromCache(period time.Duration, limit int) ([]models.Rating, error) {
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

	ratings, err := s.repository.GetTopRating(period, limit)
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

func (s *ServiceV1) getCacheKey(baseKey string, args ...interface{}) string {
	keyParts := []string{baseKey}
	for _, arg := range args {
		keyParts = append(keyParts, argToString(arg))
	}
	return strings.Join(keyParts, ":")
}

func argToString(arg interface{}) string {
	switch v := arg.(type) {
	case time.Duration:
		return v.String()
	case int:
		return strconv.Itoa(v)
	default:
		return fmt.Sprintf("%v", v)
	}
}
