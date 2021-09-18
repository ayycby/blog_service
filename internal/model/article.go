package model

type Article struct {
	*Model
	Title       string `json:"title"`
	Desc        string `json:"desc"`
	Content     string `json:"content"`
	CoverImgUrl string `json:"cover_img_url"`
	State       uint8  `json:"state"`
}

func (a Article) TableName() string {
	return "blog_article"
}
