package session

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/Jimbo8702/lets_get_one/util"
	"github.com/alexedwards/scs/v2"
)

type Session struct {
	CookieLifetime string
	CookiePersist  string
	CookieName 	   string
	CookieDomain   string
	SessionType    string
	CookieSecure   string
}

func New(con *util.Config) *scs.SessionManager {
	sess := &Session{}
	return sess.init()
}

func(c *Session) init() *scs.SessionManager {
	//defaults to false
	var persist, secure bool

	// how long should sessions last? 
	minutes, err := strconv.Atoi(c.CookieLifetime)
	if err != nil {
		minutes = 60
	}

	// should cookies persist?
	if strings.ToLower(c.CookiePersist) == "true" {
		persist = true
	} 

	// must cookies be secure? 
	if strings.ToLower(c.CookieSecure) == "true" {
		secure = true
	} 

	// create session
	session := scs.New()
	session.Lifetime = time.Duration(minutes) * time.Minute
	session.Cookie.Persist = persist
	session.Cookie.Secure = secure
	session.Cookie.Name = c.CookieName
	session.Cookie.Domain = c.CookieDomain
	session.Cookie.SameSite = http.SameSiteLaxMode

	// which session store?
	switch strings.ToLower(c.SessionType) {
	case "redis":
		//
	case "postgres", "postgresql":
		//
	default:
		//cookie
	}

	return session
}