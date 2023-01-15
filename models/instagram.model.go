package models

type FeedsBody struct {
	Edges []struct {
		Node struct {
			IsVideo    bool   `json:"is_video"`
			DisplayUrl string `json:"display_url"`
		}
	}
	PageInfo struct {
		HasNextPage bool   `json:"has_next_page"`
		EndCursor   string `json:"end_cursor"`
	} `json:"page_info"`
}

type FeedsResponse struct {
	Body FeedsBody
}

type Instagram struct {
	UserID    string   `json:"user_id"  bson:"user_id"`
	Username  string   `json:"username" bson:"username"`
	UserFeeds []string `                bson:"user_feeds"`
}

type InstagramRepository interface {
	FindOne(username string) (*Instagram, error)
	Save(instagram *Instagram) error
}
