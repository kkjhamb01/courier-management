package api

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/kkjhamb01/courier-management/common/config"
)

type redirectHandler struct {
	config   config.UaaData
	template *template.Template
}

func (h *redirectHandler) render(oauth string, code string, userType string) string {
	var tpl bytes.Buffer
	var data = make(map[string]string)
	data["code"] = code
	data["oauth"] = oauth
	data["type"] = userType
	err := h.template.Execute(&tpl, data)
	if err != nil {
		log.Println("HTTP Server cannot decode to deeplink template")
		return ""
	}
	return tpl.String()
}

func (h *redirectHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	code := request.URL.Query().Get("code")
	log.Printf("HTTP Server New Request code = %v, path = %v\n", code, request.URL.Path)
	if code != "" {
		var oauth string
		if strings.Contains(request.URL.Path, "/google") {
			oauth = "google"
		} else if strings.Contains(request.URL.Path, "/facebook") {
			oauth = "facebook"
		} else {
			return
		}

		var userType string
		if strings.Contains(request.URL.Path, "/driver") {
			userType = "courier"
		} else if strings.Contains(request.URL.Path, "/passenger") {
			userType = "client"
		} else {
			return
		}

		response.Write([]byte(h.render(oauth, code, userType)))
	}
}

func StartWebServer() {
	config := config.Uaa()

	templateStr := config.DeepLinkTemplate
	templateData, err := ioutil.ReadFile(templateStr)
	if err != nil {
		log.Fatalf("HTTP Server cannot find deep link template %v", err)
	}
	templateStr = string(templateData)

	template, err := template.New("deeplink").Parse(templateStr)
	if err != nil {
		log.Fatalf("HTTP Server invalid deep link template %v", err)
	}

	redirectHandler := &http.Server{
		Addr: ":8087",
		Handler: &redirectHandler{
			config:   config,
			template: template,
		},
	}
	go func() {
		log.Println("HTTP Server Start redirect link handler")
		log.Fatal(redirectHandler.ListenAndServe())
	}()

}
