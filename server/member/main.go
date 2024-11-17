package main

import (
	"github.com/ei-sugimoto/microtodo/server/member/cmd"
	"github.com/ei-sugimoto/microtodo/server/member/infra"
)

func main() {
	infra.Migrate()

	cmd.Serve()
}
