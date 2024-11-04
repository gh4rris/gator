package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gh4rris/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	if len(cmd.Args) > 1 {
		return fmt.Errorf("usage: %s <limit>", cmd.Name)
	}
	limit := 2
	if len(cmd.Args) == 1 {
		if converted, err := strconv.Atoi(cmd.Args[0]); err == nil {
			limit = converted
		} else {
			return fmt.Errorf("invalid limit: %w", err)
		}
	}
	getPostsParams := database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	}
	posts, err := s.db.GetPostsForUser(context.Background(), getPostsParams)
	if err != nil {
		return fmt.Errorf("couln't get users posts: %w", err)
	}
	for _, post := range posts {
		fmt.Printf("Feed: %s\n", post.FeedName)
		fmt.Printf("Published: %s\n", post.PublishedAt.Time.Format("Mon Jan 1"))
		fmt.Printf("Title: %s\n", post.Title)
		fmt.Printf("Description: %s\n", post.Description.String)
		fmt.Printf("Url: %s\n", post.Url)
		fmt.Println("===========================")
	}
	return nil
}
