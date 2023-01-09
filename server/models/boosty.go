package models

type MediaAlbum struct {
	Data  `json:"data"`
	Extra `json:"extra"`
}

type Extra struct {
	Offset string `json:"offset"`
	IsLast bool   `json:"isLast"`
}

type Data struct {
	MediaPosts `json:"mediaPosts"`
}

type MediaPosts []MediaPost

type Media []XMedia

type MediaPost struct {
	Media `json:"media"`
	Post  `json:"post"`
}

type XMedia struct {
	Height    int    `json:"height"`
	Type      string `json:"type"`
	Width     int    `json:"width"`
	Url       string `json:"url"`
	Id        string `json:"id"`
	Rendition string `json:"rendition"`
}

type Post struct {
	Id                  string `json:"id"`
	SignedQuery         string `json:"signedQuery"`
	Price               int    `json:"price"`
	Teaser              `json:"teaser"`
	HasAccess           bool   `json:"hasAccess"`
	SubscriptionLevelId int    `json:"subscriptionLevelId"`
	Title               string `json:"title"`
}

type Teaser []XTeaser

type XTeaser struct {
	Height    int    `json:"height"`
	Width     int    `json:"width"`
	Type      string `json:"type"`
	Url       string `json:"url"`
	Id        string `json:"id"`
	Rendition string `json:"rendition"`
}
