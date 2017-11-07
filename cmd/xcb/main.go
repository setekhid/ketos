package main

import (
	"github.com/setekhid/ketos/cmd/commit"
	"github.com/setekhid/ketos/cmd/init"
	"github.com/setekhid/ketos/cmd/oplayer"
	"github.com/setekhid/ketos/cmd/opmanifest"
	"github.com/setekhid/ketos/cmd/pull"
	"github.com/setekhid/ketos/cmd/push"
	"github.com/setekhid/ketos/cmd/tag"
	"github.com/setekhid/ketos/cmd/version"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetLevel(log.DebugLevel)
}

func main() {

	Command.AddCommand(
		version.Command,

		// upper level
		initk.Command,
		tag.Command,
		pull.Command,
		commit.Command,
		push.Command,

		// lower level
		opmanifest.CatManifest,
		opmanifest.PutManifest,
		oplayer.HasLayer,
		oplayer.PullLayer,
		oplayer.PushLayer,
	)

	err := Command.Execute()
	if err != nil {
		log.Fatalln(err)
	}
}
