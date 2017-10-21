package rootpath

import (
	"os"
	"strings"
)

var (
	KetosChrootWD     = IsChrootWorkingDir()
	KetosChrootRoot   = ChrootDirectory()
	KetosChrootImgTag = ChrootImageTag()
)

func IsChrootWorkingDir() bool {

	isOverlay := false
	overlayEnv := strings.ToLower(os.Getenv("KETOS_CHROOT_WD"))
	if overlayEnv == "true" || overlayEnv == "t" ||
		overlayEnv == "yes" || overlayEnv == "y" ||
		overlayEnv == "1" {

		isOverlay = true
	}

	return isOverlay
}

func ChrootDirectory() string {

	dir := os.Getenv("KETOS_CHROOT_ROOT")
	if len(dir) <= 0 {
		dir = "/"
	}

	return dir
}

func ChrootImageTag() string {

	tag := os.Getenv("KETOS_CHROOT_IMGTAG")
	if len(tag) <= 0 {
		tag = "latest"
	}

	return tag
}
