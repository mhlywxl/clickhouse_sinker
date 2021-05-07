/*Copyright [2019] housepower

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

package model

import (
	"sync"
)

// Metric interface for metric collection
type Metric interface {
	GetString(key string, nullable bool) (val interface{})
	GetFloat(key string, nullable bool) (val interface{})
	GetInt(key string, nullable bool) (val interface{})
	GetDate(key string, nullable bool) (val interface{})
	GetDateTime(key string, nullable bool) (val interface{})
	GetDateTime64(key string, nullable bool) (val interface{})
	GetElasticDateTime(key string, nullable bool) (val interface{})
	GetArray(key string, t string) (val interface{})
	GetNewKeys(knownKeys *sync.Map, newKeys *sync.Map) bool
}

// DimMetrics
type DimMetrics struct {
	Dims   []*ColumnWithType
	Fields []*ColumnWithType
}

// ColumnWithType
type ColumnWithType struct {
	Name       string
	Type       int
	Nullable   bool
	SourceName string
}
