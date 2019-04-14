package digit

import "fmt"

func ExampleLink() {

	link := NewLink("http://webfinger.example/rel/profile-page", "text/html", "https://www.example.com/~bob", "Bob Smith")

	fmt.Print(link.Titles["und"]) // Default language is "und" for undetermined.
	// Output: Bob Smith

	// You can also set a specific language on the link object itself
	link.Title("en-us", "The Magical World of Steve")
	link.Title("fr", "Le Mondo Magique de Steve")
}
