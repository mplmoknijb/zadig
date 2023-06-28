/*
Copyright 2023 The KodeRover Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LLMIntegration struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"  json:"id"`
	Name        string             `bson:"name"           json:"name"`
	Token       string             `bson:"token"          json:"token"`
	BaseURL     string             `bson:"base_url"       json:"base_url"`
	EnableProxy bool               `bson:"enable_proxy"   json:"enable_proxy"`
	UpdatedBy   string             `bson:"updated_by"     json:"updated_by"`
	UpdateTime  int64              `bson:"update_time"    json:"update_time"`
}

func (llm LLMIntegration) TableName() string {
	return "llm_integration"
}