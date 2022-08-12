package main

// made it all in 1 day
import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/Delta456/box-cli-maker/v2"
	// non-stdlib
)

type PoemInfo []struct {
	Title     string   `json:"title"` // these `json:"{JSON NAME}"` are so that it looks for the title not Title
	Author    string   `json:"author"`
	Lines     []string `json:"lines"`
	Linecount string   `json:"linecount"`
}

func returnPoemInfo() (author, title string, body string) {
	PoemPage := "https://poetrydb.org/random"
	response, err := http.Get(PoemPage) // Easy peasy

	// Error handling
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close() /* adds it to function stack so that before the
	returnPoemInfo function exits, the deferred function
	function  executes. */

	if response.StatusCode != 200 {
		log.Fatalf("Failed to fetch data: %d %s", response.StatusCode, response.Status)
	}

	respbody, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}
	// end retieval of poem info.

	// Todo is to convert the json string into Go variables. Starting from this line.
	var currentPoem PoemInfo

	err = json.Unmarshal(respbody, &currentPoem) // Converts it from string to json structure

	if err != nil {
		log.Fatal(err) // if error with json unmarshalling then panic.
	}

	var poemBody string
	for i := 0; i < len(currentPoem[0].Lines); i++ {
		poemBody += currentPoem[0].Lines[i] + "\n"
	}

	return currentPoem[0].Title, currentPoem[0].Author, poemBody
}

func main() {
	title, author, body := returnPoemInfo()
	Box := box.New(box.Config{Px: 8, Py: 5, Type: "Double", TitlePos: "Top", Color: "HiBlue"})
	Box.Print((title + " by " + author), body)

	fmt.Println("Press ENTER to exit")
	fmt.Scanln()
}
