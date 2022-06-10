package routes

import (
	"io"
	"net/http"
)

type Route struct {
	Method string
	HttpResponse http.ResponseWriter
}

func (route *Route) render() {
	switch route.Method {
		case "GET":
			route.getHome()
		case "POST":
			route.postHome()
		case "DELETE":
			route.deleteHome()
		default:
			route.unhandledMethod()
	}
}

func (route *Route) getHome() {
	io.WriteString(route.HttpResponse, GetHomePage()) 
}

func (route *Route) postHome() {
	io.WriteString(route.HttpResponse, "We are creating our home") 
}

func (route *Route) deleteHome() {
	io.WriteString(route.HttpResponse, "We are delete our home") 
}

func (route *Route) unhandledMethod() {
	io.WriteString(route.HttpResponse, "This method is unhandled") 
}

func HandleHome(w http.ResponseWriter, req *http.Request) {
	route := Route{req.Method, w}

	route.render()
}
