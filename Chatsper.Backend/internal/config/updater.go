package config

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/user"
	"runtime"
	"strings"

	"github.com/pterm/pterm"
	"zip.jespersen.chatsper/internal/utils"
	"zip.jespersen.chatsper/shared/static"
)

type GitHubRelease struct {
	URL             string         `json:"url"`
	AssetsURL       string         `json:"assets_url"`
	UploadURL       string         `json:"upload_url"`
	HTMLURL         string         `json:"html_url"`
	ID              int64          `json:"id"`
	Author          GitHubUser     `json:"author"`
	NodeID          string         `json:"node_id"`
	TagName         string         `json:"tag_name"`
	TargetCommitish string         `json:"target_commitish"`
	Name            string         `json:"name"`
	Draft           bool           `json:"draft"`
	Immutable       bool           `json:"immutable"`
	Prerelease      bool           `json:"prerelease"`
	CreatedAt       string         `json:"created_at"`
	UpdatedAt       string         `json:"updated_at"`
	PublishedAt     string         `json:"published_at"`
	Assets          []ReleaseAsset `json:"assets"`
	TarballURL      string         `json:"tarball_url"`
	ZipballURL      string         `json:"zipball_url"`
	Body            string         `json:"body"`
}

type GitHubUser struct {
	Login             string `json:"login"`
	ID                int64  `json:"id"`
	NodeID            string `json:"node_id"`
	AvatarURL         string `json:"avatar_url"`
	GravatarID        string `json:"gravatar_id"`
	URL               string `json:"url"`
	HTMLURL           string `json:"html_url"`
	FollowersURL      string `json:"followers_url"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	OrganizationsURL  string `json:"organizations_url"`
	ReposURL          string `json:"repos_url"`
	EventsURL         string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	Type              string `json:"type"`
	UserViewType      string `json:"user_view_type"`
	SiteAdmin         bool   `json:"site_admin"`
}

type ReleaseAsset struct {
	URL                string     `json:"url"`
	ID                 int64      `json:"id"`
	NodeID             string     `json:"node_id"`
	Name               string     `json:"name"`
	Label              *string    `json:"label"`
	Uploader           GitHubUser `json:"uploader"`
	ContentType        string     `json:"content_type"`
	State              string     `json:"state"`
	Size               int64      `json:"size"`
	Digest             string     `json:"digest"`
	DownloadCount      int        `json:"download_count"`
	CreatedAt          string     `json:"created_at"`
	UpdatedAt          string     `json:"updated_at"`
	BrowserDownloadURL string     `json:"browser_download_url"`
}

func getDownloadUrl(asset ReleaseAsset, os string, url *string) {
	if strings.Contains(asset.Name, os) {
		if len(*url) > 1 {
			return
		}
		*url = asset.BrowserDownloadURL
	}
}

func runUpdate(release GitHubRelease) {
	osType := runtime.GOOS

	var downloadUrl string
	for id := range release.Assets {
		asset := release.Assets[id]
		switch osType {
		case "linux":
			getDownloadUrl(asset, "linux", &downloadUrl)
		case "windows":
			getDownloadUrl(asset, "windows", &downloadUrl)
		case "darwin":
			getDownloadUrl(asset, "macos", &downloadUrl)
		}
	}

	user, _ := user.Current()
	version := utils.Version()
	chatsperDocumentPath := user.HomeDir + "/Dokumente/Chatsper/" + version
	os.MkdirAll(chatsperDocumentPath, os.FileMode(0777))

	fmt.Println("Download the latest Chatsper Release!")
	req, err := http.Get(downloadUrl)
	if err != nil {
		static.ContactTheTeam()
		os.Exit(0)
	}
	defer req.Body.Close()

	bytes, _ := io.ReadAll(req.Body)
	releaseFile := chatsperDocumentPath + "/chatsper"

	os.WriteFile(releaseFile, bytes, os.FileMode(0777))
	fmt.Println("Downloaded successfuly to " + chatsperDocumentPath)
	os.Exit(0)
}

func Updater() {
	if static.Version != strings.Split(utils.Version(), "@")[0] {
		pterm.Bold.Println(pterm.Yellow("You are running version " + static.Version + ". Latest release is " + utils.Version()))
		fmt.Println("Running updater and exiting...")

		var releases []GitHubRelease
		req, err := http.Get("https://api.github.com/repos/jesperrichert/Chatsper/releases")
		if err != nil {
			static.ContactTheTeam()
			os.Exit(0)
		}
		defer req.Body.Close()
		json.NewDecoder(req.Body).Decode(&releases)

		fmt.Println("Loading latest release \"" + releases[0].Name + "\" (" + releases[0].TagName + ")")
		fmt.Println("Read Changelog: " + releases[0].HTMLURL)

		fmt.Println()
		fmt.Println(pterm.Gray("Note: Chatsper will update and download the new version in the " + pterm.LightBlue("~/Documents/Chatsper/"+utils.Version()) + " and after this you can copy the old data or re-setup Chatsper to use the new Version!"))
		fmt.Println()

		runUpdate(releases[0])
	} else {
		fmt.Println(pterm.White("Chatsper running the latest version and is up-to-date! (" + utils.Version() + ""))
	}
}
