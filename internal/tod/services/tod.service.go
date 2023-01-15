package tod

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/alitdarmaputra/nadeshiko-bot/models"
)

type TodService struct {
}

func NewTodService() *TodService {
	return &TodService{}
}

func (t *TodService) GetTOD(option string, to string) (string, error) {
	option = strings.ToLower(option)
	var content string

	if option == "truth" || option == "dare" {
		req, _ := http.NewRequest(
			http.MethodGet,
			os.Getenv("TOD_URL")+option+"?rating="+os.Getenv("TOD_RATION"),
			nil,
		)

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return content, err
		}
		defer res.Body.Close()

		if res.StatusCode == http.StatusOK {
			var todResponse models.TODResponse
			_ = json.NewDecoder(res.Body).Decode(&todResponse)

			content = fmt.Sprintf("To: <@%s>\n%s: %s", to, todResponse.Type, todResponse.Question)
			return content, nil
		}
		return content, errors.New("Ups, somehting went wrong when getting user question")
	} else {
		return "Invalid option, only `truth` or `dare` are valid", nil
	}
}
