package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/alitdarmaputra/nadeshiko-bot/pkg/structs"
	"github.com/bwmarrin/discordgo"
)

func HelpService(m *discordgo.MessageCreate) string {
	var content = fmt.Sprintf("Hello %s, nice to meet you ^_^\n**Keywords: help**\n\nFor help type:\n!usage <available keyword>", m.Author.Username)
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

	for {
		var endCursor string
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

		if res.StatusCode == http.StatusOK {
			var feedsResponse structs.FeedsResponse

			err = json.NewDecoder(res.Body).Decode(&feedsResponse)
			if err != nil {
				return userPhotos, err
			}
			fmt.Println(feedsResponse.Body.PageInfo.EndCursor)

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
		}
	}

	if len(userPhotos) > 0 {
		return userPhotos, nil
	} else {
		return userPhotos, errors.New("Ups, something went wrong while getting user photos")
	}
}
