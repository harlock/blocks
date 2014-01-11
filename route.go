package blocks

import (
	"reflect"
	"log"
	"strings"
	"regexp"
)

type Route struct {
	RouteNode
	controllerT reflect.Type
	method string
	pathRegExp *regexp.Regexp
}

func newRoute(path string, controller interface{}, action string) *Route {
	r := new(Route)
	r.RouteNode.initialize()
	r.controllerT = reflect.TypeOf(controller)
	r.setPath(path)
	r.method = action

	return r
}


// Replace variables for regular expressions
// and compile the regexp with the path
func (this *Route) setPath(path string) {
	this.path = path

	// Replace variables with regexp
	replaced := routeVarsReplaceRegExp.ReplaceAllString(path, routeVarsReplacement)
	this.pathRegExp = regexp.MustCompile(replaced)
}
var routeVarsReplaceRegExp = regexp.MustCompile(`:([\w]+)`)
const routeVarsReplacement = "(?P<$1>.+)"


func (this *Route) Match(p Pather) bool {
	return this.pathRegExp.MatchString(p.Path())
}

// Find Routes
func (this *Route) Find(request Pather) (*Route, bool)  {
	// Find first in childrens
	route, found := this.findChildrens(request)
	if found {
		return route, true
	}

	// if childrens don't match, try to with the route itself
	if this.Match(request) {
		return this, true;
	}
	return nil, false
}


// Return the name lowercase and without controller
// example: HomeController => home
func (r *Route) ControllerName() string {
	return extractControllerName(r.controllerT)
}

func (r *Route) ActionName() string {
	return strings.ToLower(r.method)
}

func (this *Route) Inspect() {
	log.Println(this.Path())
	for _, r := range this.routes {
		r.Inspect()
	}
}