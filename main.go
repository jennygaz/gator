package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/jennygaz/gator/internal/config"
	"github.com/jennygaz/gator/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DBURL)
	dbQueries := database.New(db)

	appState := State{
		cfg: &cfg,
		db:  dbQueries,
	}

	cmds := Commands{
		registeredCommands: make(map[string]func(*State, Command) error),
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", resetHandler)
	cmds.register("users", usersHandler)
	cmds.register("agg", handlerAggregator)
	cmds.register("addfeed", handlerAddFeed)
	cmds.register("feeds", handlerGetFeeds)
	cmds.register("follow", followHandler)
	cmds.register("following", handlerListFeedFollows)

	if len(os.Args) < 2 {
		log.Fatalf("Usage: cli <command> [args...]")
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	err = cmds.run(&appState, Command{Name: cmdName, Args: cmdArgs})
	if err != nil {
		log.Fatal(err)
	}
}
