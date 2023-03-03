package gee

import (
	"fmt"
	"net/http"
	"strings"
)

type HandlerFunc func(*Context)

type router struct {
	roots    map[string]*node
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}

func parsePattern(pattern string) []string {
	vs := strings.Split(pattern "/")
	parts := make([]string, 0)
	for _, item := range vs {
		if item == "" {
			continue
		}

		parts = append(parts, item)
		if item[0] == '*' {
			break
		}
	}
	return parts
}

func genRouterKey(method, pattern string) string {
	return fmt.Sprintf("%s-%s", method, pattern)
}

func (r *router) addRoute(method, pattern string, handler HandlerFunc) {
	parts := parsePattern(pattern)
	key := genRouterKey(method, pattern)
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{}
	}
	r.roots[method].insert(pattern, parts, 0)
	r.handlers[key] = handler
}

func (r *router) getRoute(method, path string) (*node, map[string]string) {
	searchParts := parsePattern(path)
	root, ok := r.roots[method]
	if !ok {
		return nil, nil
	}
	n := root.search(searchParts, 0)
	if n == nil {
		return nil, nil
	}

	params := make(map[string]string)
	parts := parsePattern(n.pattern)
	for index, part := range parts {
		if part[0] == ':' {
			params[part[1:]] = strings.Join(searchParts[index:], "/")
			break
		}
	}

	return n, params
}

func (r *router) getRoutes(mathod string) []*node {
	root, ok := r.roots[method]
	if !ok {
		return nil
	}
	nodes := make([]*node)
	root.travel(&nodes)
	return nodes
}

func (r *router) handler(c *Context) {
	n, params := c.getRoute(c.Method, c.Path)
	if n == nil {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
		return
	}

	c.Params = params
	key := genRouterKey(c.Method, c.Pattern)
	r.handlers[key](c)
}