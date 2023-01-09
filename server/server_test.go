package server

import "testing"

func TestGetArtistById(t *testing.T) {

	artistReq := []int{1, 5, 8, 99}

	expectedMembers := []int{7,1,1,0}

	exepctedCreationDates := []int{1970,2013,2004,0}
	for i, v := range artistReq {
		artists := GetArtistById(v)
		if len(artists.Members) != expectedMembers[i] {
			t.Errorf(" ArtistId: %d - Expected %d members, got %d", v, expectedMembers[i], len(artists.Members))
		}
		if artists.CreationDate != exepctedCreationDates[i] {
			t.Errorf(" ArtistId: %d - Expected creation date %d, got %d", v, exepctedCreationDates[i], artists.CreationDate)
		}
	}
}
