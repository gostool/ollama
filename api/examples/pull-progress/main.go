package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ollama/ollama/api"
)

func main() {
	client, err := api.ClientFromEnvironment()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	req := &api.PullRequest{
		Model: "mistral",
	}
	progressFunc := func(resp api.ProgressResponse) error {
		fmt.Printf("Progress: status=%v, total=%v, completed=%v\n", resp.Status, resp.Total, resp.Completed)
		return nil
	}

	err = client.Pull(ctx, req, progressFunc)
	if err != nil {
		log.Fatal(err)
	}
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
