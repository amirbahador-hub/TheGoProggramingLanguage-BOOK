package main

import (
	"awesomeProject/Git/TheGoProggramingLanguage-BOOK-/ch4/4.5.json/exc4.10/model"
	"log"
	"os"
	"text/template"
	"time"
)

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

func daysLeft(t time.Time) string {
	now := time.Now()
	day := now.AddDate(0, 0, -1)
	month := now.AddDate(0, -1, 0)
	year := now.AddDate(-1, 0, 0)
	if t.Before(day) {
		return "less than the month old"
	} else if t.Before(month) {
		return "less than year old"
	} else if t.Before(year) {
		return "more than year old"
	}
	return ""
}

const templ = `{{.TotalCount}} issues:
{{range .Items}}----------------------------------------
Number: {{.Number}}
User:   {{.User.Login}}
Title:  {{.Title | printf "%.64s"}}
Age:    {{.CreatedAt | daysAgo}} days
Time:    {{.CreatedAt | daysLeft}} days
{{end}}`

func main() {
	report := template.Must(template.New("issueList").Funcs(template.FuncMap{"daysAgo": daysAgo, "daysLeft": daysLeft}).Parse(templ))
	result, err := model.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}
