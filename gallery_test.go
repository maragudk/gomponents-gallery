package gallery_test

import (
	"os"
	"testing"

	gallery "maragu.dev/gomponents-gallery"
	"maragu.dev/gomponents-gallery/internal/assert"
)

func TestOpenGraph(t *testing.T) {
	t.Run("can generate OpenGraph tags", func(t *testing.T) {
		og := gallery.OpenGraph("Partyhats for everyone!", "Bring your own glitter.", "https://www.example.com/partyhat.jpg", "https://www.example.com/partyhats-for-everyone")
		assert.Equal(t,
			`<meta property="og:type" content="website"><meta property="og:title" content="Partyhats for everyone!"><meta property="og:description" content="Bring your own glitter."><meta property="og:image" content="https://www.example.com/partyhat.jpg"><meta property="og:url" content="https://www.example.com/partyhats-for-everyone">`,
			og)
	})
}

func ExampleOpenGraph() {
	og := gallery.OpenGraph(
		"Partyhats for everyone!",
		"Bring your own glitter.",
		"https://www.example.com/partyhat.jpg",
		"https://www.example.com/partyhats-for-everyone")
	_ = og.Render(os.Stdout)
	// Output: <meta property="og:type" content="website"><meta property="og:title" content="Partyhats for everyone!"><meta property="og:description" content="Bring your own glitter."><meta property="og:image" content="https://www.example.com/partyhat.jpg"><meta property="og:url" content="https://www.example.com/partyhats-for-everyone">
}
