package sass

import (
	"bytes"
	"path/filepath"
	"sync"

	"github.com/omeid/slurp"
	"github.com/yourheropaul/gosass"
)

func Compile(c *slurp.C) slurp.Stage {

	return func(in <-chan slurp.File, out chan<- slurp.File) {

		var wg sync.WaitGroup
		defer wg.Wait()

		for file := range in {

			// Skip underscored files like Ruby sass
			if string(filepath.Base(file.Path)[0]) == "_" {
				continue
			}

			wg.Add(1)
			go func(file slurp.File) {

				defer wg.Done()

				ctx := gosass.FileContext{
					Options: gosass.Options{
						OutputStyle:  gosass.COMPRESSED_STYLE,
						IncludePaths: make([]string, 0),
					},
					InputPath:    file.Path,
					OutputString: "",
					ErrorStatus:  0,
					ErrorMessage: "",
				}

				gosass.CompileFile(&ctx)

				if ctx.ErrorStatus != 0 {
					c.Error("Sass error: ", ctx.ErrorMessage)
					return
				}

				buf := bytes.NewBufferString(ctx.OutputString)

				file.Reader = buf
				file.FileInfo.SetSize(int64(buf.Len()))

				out <- file

			}(file)
		}
	}
}
