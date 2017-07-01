package html5

import (
	"bytes"
	"context"
	"html/template"

	"github.com/bytesparadise/libasciidoc/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

var blockImageTmpl *template.Template
var blockImageMacroTmpl *template.Template

// initializes the templates
func init() {
	blockImageTmpl = newTemplate("block image", `
<div{{if .ID }} id="{{.ID.Value}}"{{ end }} class="imageblock">
<div class="content">
{{if .Link}}<a class="image" href="{{.Link.Path}}">{{end}}<img src="{{.Macro.Path}}" alt="{{.Macro.Alt}}"{{if .Macro.Width}} width="{{.Macro.Width}}"{{end}}{{if .Macro.Height}} height="{{.Macro.Height}}"{{end}}>{{if .Link}}</a>{{end}}
</div>
{{if .Title}}<div class="title">{{.Title.Content}}</div>{{end}}
</div>`)
	blockImageMacroTmpl = newTemplate("block image", "")
}

// <div id="img-foobar" class="imageblock">
// <div class="content">
// <a class="image" href="http://foo.bar"><img src="images/foo.png" alt="the foo.png image" width="600" height="400"></a>
// </div>
// <div class="title">Figure 1. A title to foobar</div>
// </div>
func renderBlockImage(ctx context.Context, block types.BlockImage) ([]byte, error) {
	result := bytes.NewBuffer(make([]byte, 0))
	err := blockImageTmpl.Execute(result, block)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to render block image")
	}
	log.Debugf("rendered block image: %s", result.Bytes())
	return result.Bytes(), nil
}