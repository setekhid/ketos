package main

import (
	"github.com/setekhid/ketos/cmd/chroot"
	"github.com/setekhid/ketos/cmd/commit"
	"github.com/setekhid/ketos/cmd/pull"
	"github.com/setekhid/ketos/cmd/push"
	log "github.com/sirupsen/logrus"
)

func main() {

	Command.AddCommand(
		pull.Command,
		chroot.Command,
		commit.Command,
		push.Command,
	)

	err := Command.Execute()
	if err != nil {
		log.Fatalln(err)
	}
}
