# URLShortener
This is the backend of url shortener for saving & retrieving the shorten links. The shorten links can be randomized or customized.

### New To docker

- `
docker run -p 127.0.0.1:5432:5432 --name url-container -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=test -d postgres:14
`

- `cd app`
- `go run main.go`

table name : urlshorts <br />
	`ID       uint64`<br />
	`Redirect string`<br />
	`URLshort string`<br />
	`Clicked  uint64`<br />
`Random   bool`<br />
