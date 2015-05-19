package wizard

import (
	"html/template"
	"log"
	"net/http"
	//"os"

	"github.com/comforme/comforme/common"
	"github.com/comforme/comforme/databaseActions"
	"github.com/comforme/comforme/requireLogin"
	"github.com/comforme/comforme/templates"
)

const invalidLink = "Invalid link."

var communitiesTemplate *template.Template
var messageTemplate *template.Template
var registerTemplate *template.Template

func init() {
	// Message page template
	messageTemplate = template.Must(template.New("siteLayout").Parse(templates.SiteLayout))
	template.Must(messageTemplate.New("nav").Parse(templates.NavlessBar))
	template.Must(messageTemplate.New("wizardContent").Parse(""))
	template.Must(messageTemplate.New("content").Parse(wizardTemplateText))

	// Community selection page template
	communitiesTemplate = template.Must(template.New("siteLayout").Parse(templates.SiteLayout))
	template.Must(communitiesTemplate.New("nav").Parse(templates.NavlessBar))
	template.Must(communitiesTemplate.New("communitiesContent").Parse(templates.Communities))
	template.Must(communitiesTemplate.New("wizardContent").Parse(communitiesTemplateText))
	template.Must(communitiesTemplate.New("content").Parse(wizardTemplateText))
}

func WizardHandler(res http.ResponseWriter, req *http.Request) {

	_, err := req.Cookie("sessionid")
	if err == nil { // Found cookie
		requireLogin.RequireLogin(introWizardHandler)(res, req)
		return
	}

	data := map[string]interface{}{}

	data["errorMsg"] = invalidLink

	common.ExecTemplate(messageTemplate, res, data)
}

func introWizardHandler(res http.ResponseWriter, req *http.Request) {

	data := map[string]interface{}{}

	cookie, err := req.Cookie("sessionid")
	if err != nil {
		log.Println("Failed to retrieve sessionid:", err)
		common.Logout(res, req)
		return
	}
	sessionid := cookie.Value

	data["communitiesCols"], err = databaseActions.GetCommunityColumns(sessionid)
	if err != nil {
		log.Println("Error listing communities:", err)
		common.Logout(res, req)
		return
	}

	common.ExecTemplate(communitiesTemplate, res, data)
}

const wizardTemplateText = `
	<div class="content">
		<div class="row">
			<div class="columns communities-settings">
				<h1><i class="fi-widget"></i> Settings Wizard</h1>
                {{if .successMsg}}<div class="alert-box success">{{.successMsg}}</div>{{end}}
                {{if .errorMsg}}<div class="alert-box alert">{{.errorMsg}}</div>{{end}}
				<section>
{{ template "wizardContent" .}}
				</section>
			</div>
		</div>
	</div>
	<script src="/static/js/settings_js"></script>
`

const communitiesTemplateText = `{ template "communitiesContent" .}}
					<form action="/">
						<button type="submit">Finish</button>
					</form>`
