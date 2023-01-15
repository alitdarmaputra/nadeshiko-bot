package instagram

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/alitdarmaputra/nadeshiko-bot/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type InstagramService struct {
	instagramRepo models.InstagramRepository
}

func NewInstagramService(instagramRepo models.InstagramRepository) *InstagramService {
	return &InstagramService{instagramRepo: instagramRepo}
}

func (i *InstagramService) GetUserId(username string) (*models.Instagram, error) {
	// check if exist in database
	instagram, err := i.instagramRepo.FindOne(username)

	if err == mongo.ErrNoDocuments {
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

		instagram = &models.Instagram{}
		err = json.NewDecoder(res.Body).Decode(instagram)

		if err != nil {
			return nil, err
		}

		return instagram, nil
	}

	if err != nil {
		return nil, err
	}

	return instagram, nil
}

func (i *InstagramService) GetUserFeeds(instagram *models.Instagram) error {
	var userPhotos []string

	var endCursor string
	for i := 0; i < 3; i++ {
		req, _ := http.NewRequest(
			http.MethodGet,
			fmt.Sprintf(
				"%s/public_user_posts?userid=%s&endcursor=%s",
				os.Getenv("INSTA_URL"),
				instagram.UserID,
				endCursor,
			),
			nil,
		)

		req.Header.Add("X-RapidAPI-Key", os.Getenv("X_RAPID_KEY"))
		req.Header.Add("X-RapidAPI-Host", os.Getenv("INSTA_HOST"))

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}
		defer res.Body.Close()

		var feedsResponse models.FeedsResponse
		if res.StatusCode == http.StatusOK {
			err = json.NewDecoder(res.Body).Decode(&feedsResponse)
			if err != nil {
				return err
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
		instagram.UserFeeds = userPhotos
		// save result to database
		err := i.instagramRepo.Save(instagram)
		if err != nil {
			fmt.Println(err)
			return err
		}
		return nil
	} else {
		return errors.New("Ups, something went wrong while getting user photos")
	}
}
