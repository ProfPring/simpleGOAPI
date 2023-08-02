package utils

import (
	album "example/web-service-gin/structs"
	"fmt"
	"strconv"
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

func MergeSort(allAlbums []album.Album) []album.Album {
	if len(allAlbums) < 2 {
		return allAlbums
	}

	//split the array into two halves and call merge sort recursively
	first := MergeSort(allAlbums[:len(allAlbums)/2])
	second := MergeSort(allAlbums[len(allAlbums)/2:])

	//return sorted array
	return merge(first, second)
}

func merge(a []album.Album, b []album.Album) []album.Album {

	final := []album.Album{}
	i := 0
	j := 0

	for i < len(a) && j < len(b) {
		if ConvStr(a[i].ID) < ConvStr(b[j].ID) {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}

	//stich the two halves back together
	for ; i < len(a); i++ {
		final = append(final, a[i])
	}
	for ; j < len(b); j++ {
		final = append(final, b[j])
	}

	return final
}

func ConvStr(stringToConvert string) int {
	marks, err := strconv.Atoi(stringToConvert)

	if err != nil {
		return -1
	} else {
		return marks
	}
}
