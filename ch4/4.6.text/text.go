package main

import (
	"awesomeProject/Git/TheGoProggramingLanguage-BOOK-/ch4/4.6.text/model"
	"html/template"
	"log"
	"os"
	"time"
)

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

const templ = `{{.TotalCount}} issues:
{{range .Items}}----------------------------------------
Number: {{.Number}}
User:   {{.User.Login}}
Title:  {{.Title | printf "%.64s"}}
Age:    {{.CreatedAt | daysAgo}} days
{{end}}`

//	enter this command and see result in result.txt : go run text.go repo:golang/go is:open json decoder > result.txt
func main() {
	report := template.Must(template.New("issueList").Funcs(template.FuncMap{"daysAgo": daysAgo}).Parse(templ))
	result, err := model.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}
