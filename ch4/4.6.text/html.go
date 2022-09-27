package main

import (
	"awesomeProject/Git/TheGoProggramingLanguage-BOOK-/ch4/4.6.text/model"
	"html/template"
	"log"
	"os"
	"time"
)

func DaysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

const temp2 = `
<h1>{{.TotalCount}} issues</h1>
<table>
<tr style='text-align: left'>
<th>#</th>
<th>State</th>
<th>User</th>
<th>Title</th>
</tr>
{{range .Items}}
<tr>
<td><a href='{{.HTMLURL}}'>{{.Number}}</td>
<td>{{.State}}</td>
<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
<td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`

//	enter this command and see result in result.html : go run html.go repo:golang/go is:open json decoder > result.html
func main() {
	report := template.Must(template.New("issueList").Funcs(template.FuncMap{"daysAgo": DaysAgo}).Parse(temp2))
	result, err := model.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}
