package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

type GitHubTag struct {
	Name       string          `json:"name"`
	ZipballUrl string          `json:"zipball_url"`
	Commit     GitHubTagCommit `json:"commit"`
	TarballUrl string          `json:"tarball_url"`
	NodeID     string          `json:"node_id"`
}
type GitHubTagCommit struct {
	Sha string `json:"sha"`
	Url string `json:"url"`
}

func Version() string {
	req, err := http.Get("https://api.github.com/repos/jesperrichert/Chatsper/tags")
	if err != nil {
		// Ignore
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(req.Body)

	var tags []GitHubTag
	err = json.NewDecoder(req.Body).Decode(&tags)
	if err != nil {
		return ""
	}
	sub := tags[0].Commit.Sha[0:7]
	return "" + tags[0].Name + "@" + sub
}
