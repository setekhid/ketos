package rootpath

import (
	"os"
	"strings"
)

var (
	KetosChrootToImg  bool
	KetosChrootDir    string
	KetosChrootImgTag string
)

func init0() {

	toImageEnv := strings.ToLower(os.Getenv("KETOS_CHROOT_TOIMG"))
	KetosChrootToImg = toImageEnv == "true" || toImageEnv == "t" ||
		toImageEnv == "yes" || toImageEnv == "y" ||
		toImageEnv == "1"

	KetosChrootDir = os.Getenv("KETOS_CHROOT_DIR")
	if len(KetosChrootDir) <= 0 {
		KetosChrootDir = "/_ketos"
	}

	KetosChrootImgTag = os.Getenv("KETOS_CHROOT_IMGTAG")
	if len(KetosChrootImgTag) <= 0 {
		KetosChrootImgTag = "latest"
	}
}
