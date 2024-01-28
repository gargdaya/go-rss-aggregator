package main

import (
	"database/sql"
	"time"

	"github.com/gargdaya/rssagg/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID       `json:"id"`
	Name      string          `json:"name"`
	Email     sql.NullString  `json:"email"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	ApiKey    string          `json:"api_key"`
}

func databaseUserToUser(user database.User) User {
	return User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		ApiKey:    user.ApiKey,
	}
}

type Feed struct {
	ID        uuid.UUID      `json:"id"`
	Name      string         `json:"name"`
	Url       string         `json:"url"`
	UserID    uuid.UUID  `json:"user_id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

func databaseFeedToFeed(feed database.Feed) Feed {
	return Feed{
		ID:        feed.ID,
		Name:      feed.Name,
		Url:       feed.Url,
		UserID:    feed.UserID,
		CreatedAt: feed.CreatedAt,
		UpdatedAt: feed.UpdatedAt,
	}
}

func databaseFeedsToFeeds(feeds []database.Feed) []Feed {
	result := []Feed{}
	for _, feed := range feeds {
		result = append(result, databaseFeedToFeed(feed))
	}
	return result
}

type FeedFollow struct {
	ID        uuid.UUID      `json:"id"`
	FeedID    uuid.UUID      `json:"feed_id"`
	UserID    uuid.UUID      `json:"user_id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

func databaseFeedFollowToFeedFollow(feedFollow database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID:        feedFollow.ID,
		FeedID:    feedFollow.FeedID,
		UserID:    feedFollow.UserID,
		CreatedAt: feedFollow.CreatedAt,
		UpdatedAt: feedFollow.UpdatedAt,
	}
}

func databaseFeedFollowsToFeedFollows(feedFollows []database.FeedFollow) []FeedFollow {
	result := []FeedFollow{}

	for _, feedFollow := range feedFollows {
		result = append(result, databaseFeedFollowToFeedFollow(feedFollow))
	}
	return result
}