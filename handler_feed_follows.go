package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Kota117/blog_aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("Usage: %s <url>", cmd.Name)
	}

	url := cmd.Args[0]
	feed, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return fmt.Errorf("Couldn't get feed: %w", err)
	}

	now := time.Now().UTC()
	feed_follow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("%v couldn't follow feed %v: %w", user.Name, feed.Name, err)
	}

	fmt.Printf("feed: %v, user: %v\n", feed_follow.FeedName, feed_follow.UserName)
	return nil
}

func handlerFollowing(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("Usage: %s", cmd.Name)
	}

	feed_follows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("Couldn't get feeds being followed by %v: %w", user.Name, err)
	}

	for _, feed_follow := range feed_follows {
		fmt.Println(feed_follow.FeedName)
	}
	return nil
}
