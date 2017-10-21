package chroot

import (
	"github.com/pkg/errors"
	"github.com/setekhid/ketos/pkg/ketos/metadata"
	"io"
	"os/exec"
	"path/filepath"
)

type Engine interface {
	Run(root string, imageTag string, cmd []string,
		stdin io.Reader, stdout, stderr io.Writer) error
}

func NewEngineByName(name string) (Engine, error) {

	switch name {

	case "ld.so":
		return newLDEngine(), nil

	default:
		return nil, errors.New("didn't recognize engine by name")
	}
}

func SeekKetosRoot(from string) (string, error) {

	ketosFolder, err := metadata.SeekKetosFolder(from)
	if err != nil {
		return "", errors.Wrap(err, "get ketos folder")
	}

	return filepath.Dir(ketosFolder), nil
}

type ldEngine struct {
	ldPreload string
}

func (e ldEngine) Run(root string, imageTag string, cmd []string,
	stdin io.Reader, stdout, stderr io.Writer) error {

	exe := exec.Command(cmd[0], cmd[1:]...)
	exe.Env = append(exe.Env,
		"LD_PRELOAD="+e.ldPreload,
		"KETOS_CHROOT_WD=TRUE",
		"KETOS_CHROOT_ROOT="+root,
		"KETOS_CHROOT_IMGTAG="+imageTag,
	)
	exe.Stdin = stdin
	exe.Stdout = stdout
	exe.Stderr = stderr

	return exe.Run()
}

func newLDEngine() Engine {
	return ldEngine{
		ldPreload: "/usr/local/lib/libketos-chroot.so",
	}
}
