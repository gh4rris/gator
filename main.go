package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gh4rris/gator/internal/config"
	"github.com/gh4rris/gator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatalf("couldn't connect to database: %v", err)
	}
	defer db.Close()
	dbQueries := database.New(db)
	programState := &state{
		db:  dbQueries,
		cfg: &cfg,
	}
	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)
	cmds.register("agg", handlerAgg)
	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}
	command := command{
		Name: os.Args[1],
		Args: os.Args[2:],
	}
	err = cmds.run(programState, command)
	if err != nil {
		log.Fatal(err)
	}
}
