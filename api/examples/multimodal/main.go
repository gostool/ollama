package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ollama/ollama/api"
)

func main() {
	if len(os.Args) <= 1 {
		log.Fatal("usage: <image name>")
	}

	imgData, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	client, err := api.ClientFromEnvironment()
	if err != nil {
		log.Fatal(err)
	}

	req := &api.GenerateRequest{
		Model:  "llama3.2-vision:11b",
		Prompt: "describe this image",
		Images: []api.ImageData{imgData},
	}

	ctx := context.Background()
	respFunc := func(resp api.GenerateResponse) error {
		// In streaming mode, responses are partial so we call fmt.Print (and not
		// Println) in order to avoid spurious newlines being introduced. The
		// model will insert its own newlines if it wants.
		fmt.Print(resp.Response)
		return nil
	}

	err = client.Generate(ctx, req, respFunc)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println()
}

//➜  multimodal git:(hyh) ✗ go build
//➜  multimodal git:(hyh) ✗ ./multimodal ~/Downloads/snow.png
//The image is a digital illustration of a snowman wearing a red knit hat and scarf, with a festive Christmas wreath in the background.
//
//*   The snowman:
//*   Wears a red knit hat with white trim.
//*   Has a carrot nose.
//*   Is adorned with a red scarf.
//*   Has coal eyes and mouth.
//*   Smiles at the viewer.
//*   Stands on a snowy surface.
//*   The Christmas wreath:
//*   Is made of evergreen branches.
//*   Features holly berries.
//*   Includes pinecones.
//*   Is hung on a door or wall.
//*   The background:
//*   Features a blue sky with snowflakes falling.
//*   Has a few trees in the distance.
//*   Includes a house or building in the far background.
//
//The image is a cheerful and festive representation of winter, perfect for the holiday season.
