//  Copyright (c) 2014 Couchbase, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 		http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package web

import (
	"github.com/knights-analytics/indexer/analysis"
	"github.com/knights-analytics/indexer/analysis/lang/en"
	"github.com/knights-analytics/indexer/analysis/token/lowercase"
	"github.com/knights-analytics/indexer/analysis/tokenizer/web"
	"github.com/knights-analytics/indexer/registry"
)

const Name = "web"

func AnalyzerConstructor(config map[string]interface{}, cache *registry.Cache) (analysis.Analyzer, error) {
	tokenizer, err := cache.TokenizerNamed(web.Name)
	if err != nil {
		return nil, err
	}
	toLowerFilter, err := cache.TokenFilterNamed(lowercase.Name)
	if err != nil {
		return nil, err
	}
	stopEnFilter, err := cache.TokenFilterNamed(en.StopName)
	if err != nil {
		return nil, err
	}
	rv := analysis.DefaultAnalyzer{
		Tokenizer: tokenizer,
		TokenFilters: []analysis.TokenFilter{
			toLowerFilter,
			stopEnFilter,
		},
	}
	return &rv, nil
}

func init() {
	err := registry.RegisterAnalyzer(Name, AnalyzerConstructor)
	if err != nil {
		panic(err)
	}
}
