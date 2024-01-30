// myarchive.go
package myarchive

import (
	"github.com/grafana/xk6"
	"github.com/grafana/xk6/output"
	"github.com/mholt/archiver"
	"io/ioutil"
	"path/filepath"
)

const modPath = "k6/x/myarchive"

func init() {
	xk6.RegisterExtension(&xk6.Extension{
		Name: modPath,
		Setup: func(ctx *xk6.ExtensionContext) {
			output.RegisterExtension(ctx, &MyArchiveExtension{})
		},
	})
}

// MyArchiveExtension represents the extension
type MyArchiveExtension struct{}

// MyArchive creates a zip file from the specified data files
func (e *MyArchiveExtension) MyArchive(zipFileName string, dataFileNames []string) error {
	zipPath := filepath.Join(ctx.ScratchDirectory(), zipFileName)
	zipOutput := archiver.NewZip()
	zipOutput.Create(zipPath)

	for _, dataFileName := range dataFileNames {
		data, err := ioutil.ReadFile(dataFileName)
		if err != nil {
			return err
		}

		zipOutput.Write(data, nil)
	}

	return zipOutput.Close()
}
