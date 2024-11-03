package main

import (
	"context"
	"fmt"
	"time"

	"github.com/gh4rris/gator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}
	feed, err := s.db.GetFeedByURL(context.Background(), cmd.Args[0])
	if err != nil {
		return fmt.Errorf("couldn't find feed in database: %w", err)
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
	fmt.Printf("Added to followed: %s\n", feed_follow.FeedName)
	fmt.Printf("User: %s\n", feed_follow.UserName)
	return nil
}

func handlerFollowing(s *state, cmd command, user database.User) error {
	if len(cmd.Args) > 0 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}
	follows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("couldn't get feed follows: %w", err)
	}
	if len(follows) == 0 {
		fmt.Printf("No feed follows found for %s\n", s.cfg.CurrentUserName)
		return nil
	}
	fmt.Printf("Name: %s\n", s.cfg.CurrentUserName)
	fmt.Println("Feeds: ")
	for _, follow := range follows {
		fmt.Printf("* %s\n", follow.FeedName)
	}
	return nil
}

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}
	unfollowParams := database.UnfollowParams{
		UserID: user.ID,
		Url:    cmd.Args[0],
	}
	err := s.db.Unfollow(context.Background(), unfollowParams)
	if err != nil {
		return fmt.Errorf("couldn't unfollow feed: %w", err)
	}
	fmt.Println("Unfollowed feed successfully")
	return nil
}
