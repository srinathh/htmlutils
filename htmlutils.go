package htmlutils

import (
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// MatchFunc matches HTML nodes.
type MatchFunc func(*html.Node) bool

// AppendAll recursively traverses the parse tree rooted under the provided
// node and appends all nodes matched by the MatchFunc to dst.
func AppendAll(dst []*html.Node, n *html.Node, mf MatchFunc) []*html.Node {
	if mf(n) {
		dst = append(dst, n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		dst = AppendAll(dst, c, mf)
	}
	return dst
}

// MatchAtom returns a MatchFunc that matches a Node with the specified Atom.
func MatchAtom(a atom.Atom) MatchFunc {
	return func(n *html.Node) bool {
		return n.DataAtom == a
	}
}

// MatchAtomAttr returns a MatchFunc that matches a Node with the specified
// Atom and a html.Attribute's namespace, key and value.
func MatchAtomAttr(a atom.Atom, namespace, key, value string) MatchFunc {
	return func(n *html.Node) bool {
		return n.DataAtom == a && GetAttr(n, namespace, key) == value
	}
}

// GetAttr fetches the value of a html.Attribute for a given namespace and key.
func GetAttr(n *html.Node, namespace, key string) string {
	for _, attr := range n.Attr {
		if attr.Namespace == namespace && attr.Key == key {
			return attr.Val
		}
	}
	return ""
}

// FindNode recursively searches for the node matched by
// MatchFunc and returns it if found.
func FindNode(n *html.Node, mf MatchFunc) *html.Node {
	if mf(n) {
		return n
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret := FindNode(c, mf)
		if ret != nil {
			return ret
		}
	}
	return nil
}

// GetData searches for all the text nodes under the provided
// node and returns concatenation of text data
func GetData(n *html.Node) string {
	ret := ""
	if n.Type == html.TextNode {
		ret = ret + n.Data
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret = ret + GetData(c)
	}
	return ret
}
