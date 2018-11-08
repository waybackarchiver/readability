package readability

import (
	"bytes"

	"golang.org/x/net/html"
)

// getElementsByTagName returns a collection of HTML elements with the given
// tag name. If tag name is an asterisk, a list of all the available HTML nodes
// will be returned instead.
//
// See: https://developer.mozilla.org/en-US/docs/Web/API/Document/getElementsByTagName
func getElementsByTagName(doc *html.Node, tag string) []*html.Node {
	var list []*html.Node
	var find func(*html.Node)

	find = func(node *html.Node) {
		if node.Type == html.ElementNode && (tag == "*" || node.Data == tag) {
			list = append(list, node)
		}

		for c := node.FirstChild; c != nil; c = c.NextSibling {
			find(c)
		}
	}

	find(doc)

	return list

// textContent returns text content of a node and its descendants.
//
// See: https://developer.mozilla.org/en-US/docs/Web/API/Node/textContent
func textContent(node *html.Node) string {
	var buf bytes.Buffer
	var fun func(*html.Node)

	fun = func(n *html.Node) {
		if n.Type == html.TextNode {
			buf.WriteString(n.Data)
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			fun(c)
		}
	}

	fun(node)

	return buf.String()
}
