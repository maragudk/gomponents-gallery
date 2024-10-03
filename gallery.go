package gallery

import (
	. "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

// OpenGraph tags for a website.
// Title is always set, og:description, og:image, and og:url only if non-empty.
func OpenGraph(title, description, image, url string) Node {
	return Group([]Node{
		Meta(Attr("property", "og:type"), Content("website")),
		Meta(Attr("property", "og:title"), Content(title)),
		If(description != "", Meta(Attr("property", "og:description"), Content(description))),
		If(image != "", Meta(Attr("property", "og:image"), Content(image))),
		If(url != "", Meta(Attr("property", "og:url"), Content(url))),
	})
}
