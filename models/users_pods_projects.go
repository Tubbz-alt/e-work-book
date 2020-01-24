package models

import (
	"time"
)

// UsersPodsProjects Model Struct
type UsersPodsProjects struct {
	Id        int
	PodId     int       `orm:"size(20)"`
	ProjectId int       `orm:"size(20)"`
	UserId    int       `orm:"size(20)"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)"`
}
