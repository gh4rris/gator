package main

import (
	"context"
	"fmt"
	"time"

	"github.com/gh4rris/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <feed_name> <url>", cmd.Name)
	}
	feedParams := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      cmd.Args[0],
		Url:       cmd.Args[1],
		UserID:    user.ID,
	}
	feed, err := s.db.CreateFeed(context.Background(), feedParams)
	if err != nil {
		return fmt.Errorf("couldn't add feed to database: %w", err)
	}
	feedFollowParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}
	feed_follow, err := s.db.CreateFeedFollow(context.Background(), feedFollowParams)
	if err != nil {
		return fmt.Errorf("couldn't create feed follow record: %w", err)
	}
	fmt.Printf("Username: %s\n", feed_follow.UserName)
	fmt.Printf("Feed Name: %s\n", feed_follow.FeedName)
	fmt.Printf("Feed ID: %v\n", feed.ID)
	fmt.Printf("Created at: %v\n", feed.CreatedAt)
	fmt.Printf("Updated at: %v\n", feed.UpdatedAt)
	fmt.Printf("URL: %s\n", feed.Url)
	return nil
}

func handlerFeeds(s *state, cmd command) error {
	if len(cmd.Args) > 0 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't retrieve feeds: %w", err)
	}
	if len(feeds) == 0 {
		fmt.Println("No leads found")
		return nil
	}
	for _, feed := range feeds {
		user, err := s.db.GetUserByID(context.Background(), feed.UserID)
		if err != nil {
			return fmt.Errorf("couldn't retrieve user data for %s: %w", feed.Name, err)
		}
		fmt.Printf("Name: %s\n", feed.Name)
		fmt.Printf("URL: %s\n", feed.Url)
		fmt.Printf("User: %s\n\n", user.Name)
	}
	return nil
}
