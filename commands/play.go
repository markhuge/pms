package commands

import (
	"fmt"

	"github.com/ambientsound/pms/api"
	"github.com/ambientsound/pms/input/lexer"
	"github.com/ambientsound/pms/song"
)

// Play plays songs in the MPD playlist.
type Play struct {
	api  api.API
	song *song.Song
	id   int
	pos  int
}

func NewPlay(api api.API) Command {
	return &Play{
		api: api,
		pos: -1,
	}
}

func (cmd *Play) Execute(class int, s string) error {
	var err error

	switch class {
	case lexer.TokenIdentifier:
		switch s {
		case "cursor":

			cmd.song = cmd.api.Songlist().CursorSong()
			if cmd.song == nil {
				return fmt.Errorf("Cannot play: no song under cursor")
			}
		default:
			return nil
		}

	case lexer.TokenEnd:
		client := cmd.api.MpdClient()
		if client == nil {
			return fmt.Errorf("Cannot play: not connected to MPD")
		}

		if cmd.song == nil {
			err = client.Play(-1)
			return err
		}

		id := cmd.song.ID

		if cmd.song.NullID() {
			id, err = client.AddID(cmd.song.StringTags["file"], -1)
			if err != nil {
				return err
			}
		}

		err = client.PlayID(id)
		return err

	default:
		return fmt.Errorf("Unknown input '%s', expected END", s)
	}

	return nil
}
