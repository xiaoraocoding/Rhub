// Package user 存放用户 Model 相关逻辑
package user

import (
	"Rhub/app/models"
)

// User 用户模型
type User struct {
	models.BaseModel

	Name     string `json:"name,omitempty"`
	Email    string `json:"-"` //使用"-" 是为了将隐私信息不进行传递
	Phone    string `json:"-"`
	Password string `json:"-"`

	models.CommonTimestampsField
}
