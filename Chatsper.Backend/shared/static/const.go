package static

import (
	"fmt"

	"github.com/pterm/pterm"
)

const Version = "0.1.7"
const ConfigVersion = "1.0.3"
const TwitchBaseAPIUrl = "https://api.twitch.tv/helix"

func ContactTheTeam() {
	fmt.Println()
	fmt.Println(pterm.LightMagenta("!!! Please contact the Team for more information and help! Docs: https://doc.jespersen.zip/s/chatspers !!!"))
	fmt.Println()
}
