package instagram

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/alitdarmaputra/nadeshiko-bot/models"
)

type InstagramService struct {
}

func NewInstagramService() *InstagramService {
	return &InstagramService{}
}

func (i *InstagramService) GetUserId(username string) (*models.Instagram, error) {
	var instagram models.Instagram

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

	err = json.NewDecoder(res.Body).Decode(&instagram)

	if err != nil {
		return nil, err
	}

	return &instagram, nil
}

func (i *InstagramService) GetUserFeeds(userId string) ([]string, error) {
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

		var feedsResponse models.FeedsResponse
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
