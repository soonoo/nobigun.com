package models

type Petition struct {
  From string `json:"name" form:"name" query:"name"`
  To string `json:"email" form:"email" query:"email"`
  Content string `json:"content" form:"content" query:"content"`
}
