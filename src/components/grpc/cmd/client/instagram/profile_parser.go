package main

import "encoding/json"

func NewProfileParser() *ProfileParser {
	return &ProfileParser{
		profilePage: &ProfilePage{},
	}
}

type ProfileParser struct {
	profilePage *ProfilePage
}

func (p *ProfileParser) Parse(rawBytes []byte) error {
	w := &WrappedBytes{}
	if err := json.Unmarshal(rawBytes, w); err != nil {
		return err
	}
	if err := json.Unmarshal(w.Bytes, p.profilePage); err != nil {
		return err
	}
	return nil
}

func (p *ProfileParser) FirstNShortcodes() (cursor string, shortcodes []string, err error) {
	media := p.profilePage.GraphQl.User.EdgeOwnerToTimelineMedia
	for _, edge := range media.Edges {
		shortcodes = append(shortcodes, edge.Node.ShortCode)
	}
	if media.PageInfo.HasNextPage {
		return media.PageInfo.EndCursor, shortcodes, nil
	}
	return "", shortcodes, nil
}

type WrappedBytes struct {
	CrawledTs int32           `json:"crawled_ts"`
	Bytes     json.RawMessage `json:"bytes"`
}

type ProfilePage struct {
	SeoCategoryInfos [][]string `json:"seo_category_infos"`
	LoggingPageId    string     `json:"logging_page_id"`
	GraphQl          struct {
		User struct {
			Biography              string          `json:"biography"`
			BlockedByViewer        bool            `json:"blocked_by_viewer"`
			RestrictedByViewer     bool            `json:"restricted_by_viewer"`
			EdgeFollowedBy         CountInt64      `json:"edge_followed_by"`
			FbId                   string          `json:"fbid"`
			Id                     string          `json:"id"`
			EdgeFollow             CountInt64      `json:"edge_follow"`
			FollowsViewer          bool            `json:"follows_viewer"`
			UserName               string          `json:"username"`
			ProfilePicUrl          string          `json:"profile_pic_url"`
			IsBusinessAccount      bool            `json:"is_business_account"`
			IsProfessionalAccount  bool            `json:"is_professional_account"`
			BusinessAddressJson    json.RawMessage `json:"business_address_json"`
			BusinessContactMethod  json.RawMessage `json:"business_contact_method"`
			BusinessEmail          json.RawMessage `json:"business_email"`
			BusinessPhoneNumber    json.RawMessage `json:"business_phone_number"`
			BusinessCategoryName   string          `json:"business_category_name"`
			EdgeFelxiVideoTimeline struct {
				Count    int64    `json:"count"`
				PageInfo PageInfo `json:"page_info"`
				Edges    []Edge   `json:"edges"`
			} `json:"edge_felix_video_timeline"`
			EdgeOwnerToTimelineMedia struct {
				Count    int64    `json:"count"`
				PageInfo PageInfo `json:"page_info"`
				Edges    []Edge   `json:"edges"`
			} `json:"edge_owner_to_timeline_media"`
		} `json:"user"`
	} `json:"graphql"`
}

type EdgeNodeType string

const (
	GraphVideo   EdgeNodeType = "GraphVideo"
	GraphSidecar EdgeNodeType = "GraphSidecar"
	GraphImage   EdgeNodeType = "GraphImage"
)

type Edge struct {
	Node struct {
		TypeName           EdgeNodeType `json:"__typename"` // GraphVideo | GraphSidecar
		Id                 string       `json:"id"`
		ShortCode          string       `json:"shortcode"`
		DashInfo           struct{}     `json:"dash_info,omitempty"`
		EdgeMediaToCaption struct {
			Edges []struct {
				Node struct {
					Text string `json:"text"`
				} `json:"node"`
			} `json:"edges"`
		} `json:"edge_media_to_caption"`
		EdgeMediaToComment    CountInt64 `json:"edge_media_to_comment"`
		EdgeMediaToToggedUser struct {
			Edges []struct {
				Node struct {
					User struct {
						Id            string `json:"id"`
						FullName      string `json:"full_name"`
						UserName      string `json:"username"`
						ProfilePicUrl string `json:"profile_pic_url"`
					} `json:"user"`
				} `json:"node"`
			} `json:"edges"`
		} `json:"edge_media_to_tagged_user"`
		EdgeSidecarToChildren struct {
			Edges []struct {
				Node struct {
					TypeName   EdgeNodeType `json:"__typename"` // GraphImage
					DisplayURL string       `json:"display_url"`
					ShortCode  string       `json:"shortcode"`
					Owner      struct {
						Id       string `json:"id"`
						Username string `json:"username"`
					} `json:"owner"`
				} `json:"node"`
			} `json:"edges"`
		} `json:"edge_sidecar_to_children,omitempty"`
		CommentsDisable      bool       `json:"comments_disabled"`
		EdgeLikedBy          CountInt64 `json:"edge_liked_by"`
		EdgeMediaPreviewLike CountInt64 `json:"edge_media_preview_like"`
		Location             struct {
			Id   string `json:"id"`
			Name string `json:"name"`
		} `json:"location,omitempty"`
	} `json:"node"`
}

type CountInt64 struct {
	Count int64 `json:"count"`
}

type PageInfo struct {
	HasNextPage bool   `json:"has_next_page"`
	EndCursor   string `json:"end_cursor"`
}

type PostPagination struct {
	Data struct {
		User struct {
			EdgeOwnerToTimelineMedia struct {
				Count    int64    `json:"count"`
				PageInfo PageInfo `json:"page_info"`
				Edges    []Edge   `json:"edges"`
			} `json:"edge_owner_to_timeline_media"`
		} `json:"user"`
	} `json:"data"`
	Status string `json:"status"`
}

type VideoPagination struct {
	Data struct {
		User struct {
			EdgeFelxiVideoTimeline struct {
				Count    int64    `json:"count"`
				PageInfo PageInfo `json:"page_info"`
				Edges    []Edge   `json:"edges"`
			} `json:"edge_felix_video_timeline"`
		} `json:"user"`
	} `json:"data"`
	Status string `json:"status"`
}
