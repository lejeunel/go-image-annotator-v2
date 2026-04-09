package web

import (
	s "github.com/lejeunel/go-image-annotator-v2/shared"
	html "github.com/lejeunel/go-image-annotator-v2/shared/html"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
	"net/http"
)

func MakeHomePage() Node {
	return html.MakeTitledPage("Home", Div(Text("This is a sentence with a "), html.MakeTextLink("#", "link")),
		html.Scripts(html.ScriptIncludes{}),
		html.NavBarActivatedItems{Home: true}, s.RepoURL)
}

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	MakeHomePage().Render(w)

}
