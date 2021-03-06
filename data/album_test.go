package data

import (
	"testing"
)

// Mock album
var album = Album{
	Artist:   "TestArtist",
	ArtistID: 1,
	Title:    "TestAlbum",
	Year:     2014,
}

// TestAlbumDatabase verifies that an Album can be saved and loaded from the database
func TestAlbumDatabase(t *testing.T) {
	// Load database configuration
	DB = new(SqliteBackend)
	DB.DSN("~/.config/wavepipe/wavepipe.db")
	if err := DB.Open(); err != nil {
		t.Fatalf("Could not open database connection: %s", err.Error())
	}
	defer DB.Close()

	// Attempt to save the album
	if err := album.Save(); err != nil {
		t.Fatalf("Could not save album: %s", err.Error())
	}

	// Attempt to load the album
	if err := album.Load(); err != nil {
		t.Fatalf("Could not load album: %s", err.Error())
	}

	// Attempt to delete the album
	if err := album.Delete(); err != nil {
		t.Fatalf("Could not delete album: %s", err.Error())
	}
}
