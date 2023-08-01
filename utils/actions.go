package utils

import (
	album "example/web-service-gin/structs"
	"fmt"
)

func DeleteAlbum(albumItemId string, allAlbums []album.Album) []album.Album {

	newLength := 0
	oldLength := len(allAlbums)

	for index := range allAlbums {
		if albumItemId != allAlbums[index].ID {
			allAlbums[newLength] = allAlbums[index]
			newLength++
		}
	}

	if newLength == oldLength {
		return allAlbums
	} else {
		newArray := allAlbums[:newLength]
		fmt.Println("array after delete", newArray)
		return newArray
	}
}
