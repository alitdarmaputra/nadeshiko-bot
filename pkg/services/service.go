package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/alitdarmaputra/nadeshiko-bot/pkg/structs"
	"github.com/bwmarrin/discordgo"
)

func HelpService(m *discordgo.MessageCreate) string {
	var content = fmt.Sprintf("Hello %s, nice to meet you ^_^\n**Keywords: help, stalk, lovecalc**\n\nFor help type:\n!usage <available keyword>", m.Author.Username)
	return content
}

func NotFoundService() string {
	var content = fmt.Sprintf(
		"Keyword not found " +
			"\n" +
			"\nTry **!help** to get list keywords",
	)
	return content
}

func GetUserId(username string) (*structs.UserInfo, error) {
	var userInfo structs.UserInfo

	req, _ := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("%s/get_user_id?username=%s", os.Getenv("INSTA_URL"), username),
		nil,
	)

	req.Header.Add("X-RapidAPI-Key", os.Getenv("X_RAPID_KEY"))
	req.Header.Add("X-RapidAPI-Host", os.Getenv("INSTA_HOST"))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("Something went wrong when getting user id")
	}

	err = json.NewDecoder(res.Body).Decode(&userInfo)

	if err != nil {
		return nil, err
	}

	return &userInfo, nil
}

func GetUserFeeds(userId string) ([]string, error) {
	var userPhotos []string

	if userId == "" {
		return userPhotos, errors.New("UserId not found")
	}

	var endCursor string
	for i := 0; i < 3; i++ {
		req, _ := http.NewRequest(
			http.MethodGet,
			fmt.Sprintf(
				"%s/public_user_posts?userid=%s&endcursor=%s",
				os.Getenv("INSTA_URL"),
				userId,
				endCursor,
			),
			nil,
		)

		req.Header.Add("X-RapidAPI-Key", os.Getenv("X_RAPID_KEY"))
		req.Header.Add("X-RapidAPI-Host", os.Getenv("INSTA_HOST"))

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return userPhotos, err
		}
		defer res.Body.Close()

		var feedsResponse structs.FeedsResponse
		if res.StatusCode == http.StatusOK {
			err = json.NewDecoder(res.Body).Decode(&feedsResponse)
			if err != nil {
				return userPhotos, err
			}
			feeds := feedsResponse.Body.Edges

			for _, feed := range feeds {
				// Filter photo only
				if !feed.Node.IsVideo {
					userPhotos = append(userPhotos, feed.Node.DisplayUrl)
				}
			}

			if !feedsResponse.Body.PageInfo.HasNextPage {
				break
			}
			endCursor = feedsResponse.Body.PageInfo.EndCursor
		}
	}

	if len(userPhotos) > 0 {
		return userPhotos, nil
	} else {
		return userPhotos, errors.New("Ups, something went wrong while getting user photos")
	}
}

func LoveCalc(name1 string, name2 string) (percent string) {
	var fullName = name1 + name2
	intFullNames := countChar(&fullName)
	intPercent := calcMatch(intFullNames)

	result := intPercent[0]*10 + intPercent[1]
	var react string

	if result > 80 {
		react = "You are a good parner"
	} else if result > 40 {
		react = "I think you should give a try"
	} else {
		react = "I think it's better just be a friend"
	}

	return fmt.Sprintf(
		"%s and %s is %d%d%% match, **%s**",
		name1,
		name2,
		intPercent[0],
		intPercent[1],
		react,
	)
}

func countChar(fullName *string) (intFullNames []int) {
	var count = map[string]int{}
	var keys []rune

	*fullName = strings.ToLower(*fullName)

	for _, e := range *fullName {
		_, isExist := count[string(e)]
		if isExist {
			count[string(e)]++
		} else {
			keys = append(keys, e)
			count[string(e)] = 1
		}
	}

	for _, key := range keys {
		intFullNames = append(intFullNames, count[string(key)])
	}
	return
}

func calcMatch(intFullNames []int) []int {
	if len(intFullNames) == 2 {
		return intFullNames
	} else {
		var subIntFullNames []int
		var left, right = 0, len(intFullNames) - 1

		for left <= right {
			if left == right {
				subIntFullNames = append(subIntFullNames, intFullNames[left])
			} else {
				subIntFullNames = append(subIntFullNames, intFullNames[left]+intFullNames[right])
			}
			left++
			right--
		}

		return calcMatch(subIntFullNames)
	}
}

func GetTOD(option string, to string) (string, error) {
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
			var todResponse structs.TODResponse
			_ = json.NewDecoder(res.Body).Decode(&todResponse)

			content = fmt.Sprintf("To: <@%s>\n%s: %s", to, todResponse.Type, todResponse.Question)
			return content, nil
		}
		return content, errors.New("Ups, somehting went wrong when getting user question")
	} else {
		return "Invalid option, only `truth` or `dare` are valid", nil
	}
}

func GetHelp(command string) string {
	if command == "lovecalc" {
		return "Type:\n!lovecalc <name1> <name2>\n**Ex: !lovecalc Nadeshiko Kagamihara**"
	} else if command == "stalk" {
		return "Type:\n!stalk <instagram_username>\n**to show 3 recent photo from instagram**"
	} else if command == "help" {
		return "Type:\n!help\n**to show help**"
	} else if command == "usage" {
		return "Type:\n!usage <command_name>\n**to show how to use a command**"
	} else if command == "tod" {
		return "Type:\n!tod <truth | dare>\n**To get a truth or dare question**"
	} else {
		return "Sorry, keyword not found"
	}
}
