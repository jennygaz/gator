package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jennygaz/gator/internal/database"
)

func handlerAggregator(s *State, cmd Command) error {
	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("couldn't fetch feed: %w", err)
	}
	fmt.Printf("Feed: %+v\n", feed)
	return nil
}

func handlerAddFeed(s *State, cmd Command) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: addfeed <feed_name> <feed_url>")
	}

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("error while getting current user: %w", err)
	}

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      cmd.Args[0],
		Url:       cmd.Args[1],
	})
	if err != nil {
		return fmt.Errorf("error while creating feed: %+v", err)
	}

	fmt.Printf("Name: %v\nURL: %v\n", feed.Name, feed.Url)
	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't create feed follow: %v", err)
	}

	fmt.Println("Feed to follow created successfully:")
	fmt.Printf("* Username: %v\n* Feed Name: %v\n", feedFollow.UserName, feedFollow.FeedName)
	fmt.Println()
	return nil
}

func handlerGetFeeds(s *State, cmd Command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("unable to get all feeds: %w", err)
	}

	for _, feed := range feeds {
		fmt.Printf("Name: %v\nURL: %v\nUser: %v\n\n", feed.FeedName, feed.FeedUrl, feed.UserName)
	}
	return nil
}
