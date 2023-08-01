package utils

import (
	album "example/web-service-gin/structs"
	"fmt"
)

func BinarySearch(itemId string, allAlbums []album.Album) album.Album {

	low := 0
	high := len(allAlbums)
	var album album.Album

	for low <= high {
		middle := (low + high) / 2

		if allAlbums[middle].ID == itemId {
			album = allAlbums[middle]
			fmt.Println("album is in stock!")
			break
		}

		if allAlbums[middle].ID < itemId {
			low = middle + 1
		} else {
			high = middle - 1
		}
	}

	return album
}
