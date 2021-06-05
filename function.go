package meat4thhhooking

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Nanai10a/hhooking"
)

func Gcf(w http.ResponseWriter, r *http.Request) {
	key := os.Getenv("PUB_KEY")
	if key == "" {
		key = "" // FIXME
	}

	hhooking.CreateInteractionHandler(key, func(i hhooking.Interaction) hhooking.InteractionReponse {
		var content string
		switch i.Data {
		case nil:
			content = fmt.Sprintf("Ok, received %v command.", "unknown")
		default:
			content = fmt.Sprintf("Ok, received %v command.", i.Data.Name)
		}

		return hhooking.InteractionReponse{Type: hhooking.IctChannelMessageWithSource, Data: &hhooking.InteractionApplicationCommandCallbackData{
			Content: &content,
		}}
	})(w, r)
}
