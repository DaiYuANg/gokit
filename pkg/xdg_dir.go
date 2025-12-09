package pkg

import (
	"os"
	"path"

	"github.com/adrg/xdg"
)

func WithUserHome(dir string) string {
	return path.Join(xdg.Home, dir)
}

func WithTemp(dir string) string {
	return path.Join(os.TempDir(), dir)
}
