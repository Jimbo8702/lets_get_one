package render

import (
	"fmt"
	"log"
	"net/http"

	"github.com/CloudyKit/jet/v6"
	"github.com/Jimbo8702/lets_get_one/util"
	"github.com/alexedwards/scs/v2"
)

type Renderer interface {
	Page(http.ResponseWriter, *http.Request, string, interface{}, interface{}) error
	checkAuth(*TemplateData,  *http.Request) *TemplateData
}

type TemplateData struct {
	CSRFToken 		string
	Port 			string
	ServerName 		string
	IsAuthenticated bool
	Secure 			bool
	IntMap 			map[string]int
	StringMap 		map[string]string
	FloatMap 		map[string]float32
	Data 			map[string]interface{}
}

type JetRenderer struct {
	RootPath 	string
	Secure 	 	bool
	Port 	 	string
	ServerName 	string
	Session 	*scs.SessionManager
	views 		*jet.Set
}

func New(sess *scs.SessionManager, config *util.Config) Renderer {
	var views = jet.NewSet(
		jet.NewOSFileSystemLoader(fmt.Sprintf("%s/views", config.RootPath)), 
		jet.InDevelopmentMode(),
	)
	
	return &JetRenderer{
		RootPath: config.RootPath,
		Port: config.Port,
		Session: sess,
		views: views,
	}
}

func (c *JetRenderer) Page(w http.ResponseWriter, r *http.Request, templateName string, variables, data interface{}) error {
	var vars jet.VarMap

	if variables == nil {
		vars = make(jet.VarMap)
	} else {
		vars = variables.(jet.VarMap)
	}

	td := &TemplateData{}
	if data != nil {
		td = data.(*TemplateData)
	}
	
	td = c.checkAuth(td, r)

	t, err := c.views.GetTemplate(fmt.Sprintf("%s.jet", templateName))
	if err != nil {
		log.Println(err)
		return err
	}
	if err = t.Execute(w, vars, td); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (c *JetRenderer) checkAuth(td *TemplateData, r *http.Request) *TemplateData {
	if c.Session.Exists(r.Context(), "userID") {
		td.IsAuthenticated = true
	}
	return td
}