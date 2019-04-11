package main

import (
	"context"
	"fmt"
	"os"

	"github.com/muxinc/mux-go"
)

// Exercises all asset operations:
//   get-asset
//   delete-asset
//   create-asset
//   list-assets
//   get-asset-input-info
//   create-asset-playback-id
//   get-asset-playback-id
//   delete-asset-playback-id
//   update-asset-mp4-support

func checkError(err error) {
	if err != nil {
		fmt.Printf("err: %s \n\n", err)
		os.Exit(255)
	}
}

func main() {
	// API Client Initialization
	client := muxgo.NewAPIClient(
		muxgo.NewConfiguration(
			muxgo.WithBasicAuth(os.Getenv("ACCESS_TOKEN_ID"), os.Getenv("ACCESS_TOKEN_SECRET")),
		))

	// ========== create-asset ==========
	cr, err := client.AssetsApi.CreateAsset(context.Background(), muxgo.CreateAssetRequest{
		Input: []muxgo.InputSettings{
			muxgo.InputSettings{
				Url: "https://storage.googleapis.com/muxdemofiles/mux-video-intro.mp4",
			},
		},
	})
	checkError(err)
	fmt.Println("create-asset OK ✅")
	fmt.Printf("created asset with id: %s\n\n", cr.Data.Id)

	// ========== list-assets ==========
	lr, err := client.AssetsApi.ListAssets(context.Background(), nil)
	checkError(err)
	fmt.Println("list-assets OK ✅")
	fmt.Printf("loaded a list of %d assets.\n\n", len(lr.Data))

	// ========== get-asset ==========
	gr, err := client.AssetsApi.GetAsset(context.Background(), cr.Data.Id)
	checkError(err)
	fmt.Println("get-asset OK ✅")
	fmt.Printf("got the asset with id %s\n\n", gr.Data.Id)

	// ========== get-asset-input-info ==========
	ir, err := client.AssetsApi.GetAssetInputInfo(context.Background(), cr.Data.Id)
	checkError(err)
	fmt.Println("get-asset-input-info OK ✅")
	fmt.Printf("%+v\n\n", ir.Data)

	// ========== create-asset-playback-id ==========

	// ========== get-asset-playback-id ==========

	// ========== update-asset-mp4-support ==========

	// ========== delete-asset-playback-id ==========

	// ========== delete-asset ==========
}
