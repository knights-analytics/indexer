//  Copyright (c) 2020 Couchbase, Inc.
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

package hr

import (
	"reflect"
	"testing"

	"github.com/knights-analytics/indexer/analysis"
	"github.com/knights-analytics/indexer/registry"
)

func TestCroatianAnalyzer(t *testing.T) {
	tests := []struct {
		input  []byte
		output analysis.TokenStream
	}{
		// stemming
		{
			input: []byte("Hrvatska"),
			output: analysis.TokenStream{
				&analysis.Token{
					Term: []byte("hrvatsk"),
				},
			},
		},
		{
			input: []byte("Hrvatski"),
			output: analysis.TokenStream{
				&analysis.Token{
					Term: []byte("hrvatsk"),
				},
			},
		},
		// uppercase letters
		{
			input: []byte("KOMARAC"),
			output: analysis.TokenStream{
				&analysis.Token{
					Term: []byte("komarc"),
				},
			},
		},
		// vowelR
		{
			input: []byte("crvi"),
			output: analysis.TokenStream{
				&analysis.Token{
					Term: []byte("crv"),
				},
			},
		},
		// stop word
		{
			input:  []byte("biti"),
			output: analysis.TokenStream{},
		},
		// suffix transformation
		{
			input: []byte("zaključcima"),
			output: analysis.TokenStream{
				&analysis.Token{
					Term: []byte("zaključk"),
				},
			},
		},
	}

	cache := registry.NewCache()
	analyzer, err := cache.AnalyzerNamed(AnalyzerName)
	if err != nil {
		t.Fatal(err)
	}
	for _, test := range tests {
		actual := analyzer.Analyze(test.input)
		if len(actual) != len(test.output) {
			t.Fatalf("expected length: %d, got %d", len(test.output), len(actual))
		}
		for i, tok := range actual {
			if !reflect.DeepEqual(tok.Term, test.output[i].Term) {
				t.Errorf("expected term %s (% x) got %s (% x)", test.output[i].Term, test.output[i].Term, tok.Term, tok.Term)
			}
		}
	}
}
