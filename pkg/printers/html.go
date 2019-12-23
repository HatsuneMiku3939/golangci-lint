package printers

import (
	"context"
	"fmt"
	"strings"

	"github.com/golangci/golangci-lint/pkg/logutils"
	"github.com/golangci/golangci-lint/pkg/result"
)

type HTML struct{}

func NewHTML() *HTML {
	return &HTML{}
}

func (p *HTML) Print(ctx context.Context, issues []result.Issue) error {
	fmt.Fprint(logutils.StdOut, `
<HTML>
	<HEAD>
		<TITLE>fesfsefsefse</TITLE>
	</HEAD>

	<BODY>
	<table style="width:100%">
	<tr>
		<th>Linter</th>
		<th>Text</th>
		<th>SourceLines</th>
	</tr>
`)
	for _, issue := range issues {
		fmt.Fprint(logutils.StdOut, fmt.Sprintf(`
	<tr>
		<td>%s</td>
		<td>%s</td>
		<td>%s</td>
	</tr>
`, issue.FromLinter, issue.Text, fmt.Sprintf("%s:%d<br>%s", issue.Pos.Filename, issue.Pos.Line, strings.Join(issue.SourceLines, "<br>"))))
	}

	fmt.Fprint(logutils.StdOut, `
	</table>
	</BODY>
</HTML>
`)
	return nil
}
