package structs

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
	}
}

type FeedsResponse struct {
	Body FeedsBody
}

type UserInfo struct {
	UserID    string `json:"user_id"`
	Username  string `json:"username"`
	UserFeeds []string
}
