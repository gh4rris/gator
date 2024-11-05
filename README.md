Gator is a multi-user RSS blog aggregator application. It is a CLI
tool and intended for local use.

In order to run the program, you will need both Postgres and Go
installed.

To install Gator, run the Go command:

go install github.com/gh4rris/gator

Once the program is installed you will need to set up a configuration
file in your home directory: ~/.gatorconfig.json

Copy and paste the following configuration into your json file and
replace the values with your database connection string:

{
"db_url":"postgres://postgres:username@localhost:5432/database?sslmode=disable",
"current_user_name":""
}

Commands for use:

- gator register <name>: register a new user
- gator login <name>: login to a already registered user
- gator users: list users in database
- gator addfeed <feed_name> <url>: add a new feed to the database that can
  be aggregated, current user will automatically follow
- gator feeds: list all the feeds and the user who added them
- gator follow <url>: user follow new feed
- gator unfollow <url>: user unfollow feed
- gator following: list current users following feeds
- gator agg <duration>: aggregates through feeds and posts them so they can
  be browsed. (duration format: 1m/30s/1m30s)
- gator browse: browse through posts in feeds current user is following
- gator reset: reset the database
