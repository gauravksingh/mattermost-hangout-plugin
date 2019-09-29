package main

import (
	"fmt"
    "io/ioutil"
    "net/http"
	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/plugin"
)

type MMHangoutPlugin struct {
	plugin.MattermostPlugin
}

func (p *MMHangoutPlugin) ServeHTTP(c *plugin.Context, w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}

func generateFile(name string, data string) {
	ioutil.WriteFile("/tmp/mm-" + name, []byte(data), 0777)
}

func (p *MMHangoutPlugin) MessageHasBeenPosted(c *plugin.Context, post *model.Post) {
    user, err := p.API.GetUser(post.UserId)
    if err != nil {
        p.API.LogError("failed to query user", "user_id", post.UserId)
        return
    }

    channel, err := p.API.GetChannel(post.ChannelId)
    if err != nil {
        p.API.LogError("failed to query channel", "channel_id", post.ChannelId)
        return
    }

    msg := fmt.Sprintf("MessageHasBeenPosted:@%s---%s", user.Username, channel.Name)
    generateFile(msg, post.Message)
}

// This example demonstrates a plugin that handles HTTP requests which respond by greeting the
// world.
func main() {
	plugin.ClientMain(&MMHangoutPlugin{})
}
