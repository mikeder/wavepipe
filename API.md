API
===

wavepipe features a simple API which is used to retrieve metadata from media files, as well as endpoints
to retrieve a file stream from the server.

At this time, the API is **unstable**, and is subject to change.

## Albums
Used to retrieve information about albums from wavepipe.  If an ID is specified, information will be
retrieved about a single album.

**URL:** `/api/albums/:id`

**Examples:** `http://localhost:8080/api/albums/`, `http://localhost:8080/api/albums/1`

**Return JSON:**

| Name | Type | Description |
| :--: | :--: | :---------: |
| error | Error/null | Information about any errors that occurred.  Value is null if no error occurred. |
| albums | []Album | Array of Album objects returned by the API. |
| songs | []Song/null | If ID is specified, array of Song objects attached to this album.  Value is null if no ID specified. |

## Artists
Used to retrieve information about artists from wavepipe.  If an ID is specified, information will be
retrieved about a single artist.

**URL:** `/api/artists/:id`

**Examples:** `http://localhost:8080/api/artists/`, `http://localhost:8080/api/artists/1`

**Parameters:**

| Name | Type | Required | Description |
| :--: | :--: | :------: | :---------: |
| songs | bool | | If true, returns all songs attached to this artist. |

**Return JSON:**

| Name | Type | Description |
| :--: | :--: | :---------: |
| error | Error/null | Information about any errors that occurred.  Value is null if no error occurred. |
| artists | []Artist | Array of Artist objects returned by the API. |
| albums | []Album/null | If ID is specified, array of Album objects attached to this artist. |
| songs | []Song/null | If parameter `songs` is true, array of Song objects attached to this artist.  Value is null if parameter `songs` is false or not specified. |

## Songs
Used to retrieve information about songs from wavepipe.  If an ID is specified, information will be
retrieved about a single song.

**URL:** `/api/songs/:id`

**Examples:** `http://localhost:8080/api/songs/`, `http://localhost:8080/api/songs/1`

**Return JSON:**

| Name | Type | Description |
| :--: | :--: | :---------: |
| error | Error/null | Information about any errors that occurred.  Value is null if no error occurred. |
| songs | []Song | Array of Song objects returned by the API. |