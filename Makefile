all: run
run:
	go run server.go actors.go auth.go config.go database.go directors.go movies.go movies_actions.go router.go users.go utils.go