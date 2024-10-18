//  Copyright (c) 2015 Couchbase, Inc.
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

package config

import (
	// token maps
	_ "github.com/knights-analytics/indexer/analysis/tokenmap"

	// fragment formatters
	_ "github.com/knights-analytics/indexer/search/highlight/format/ansi"
	_ "github.com/knights-analytics/indexer/search/highlight/format/html"

	// fragmenters
	_ "github.com/knights-analytics/indexer/search/highlight/fragmenter/simple"

	// highlighters
	_ "github.com/knights-analytics/indexer/search/highlight/highlighter/ansi"
	_ "github.com/knights-analytics/indexer/search/highlight/highlighter/html"
	_ "github.com/knights-analytics/indexer/search/highlight/highlighter/simple"

	// char filters
	_ "github.com/knights-analytics/indexer/analysis/char/asciifolding"
	_ "github.com/knights-analytics/indexer/analysis/char/html"
	_ "github.com/knights-analytics/indexer/analysis/char/regexp"
	_ "github.com/knights-analytics/indexer/analysis/char/zerowidthnonjoiner"

	// analyzers
	_ "github.com/knights-analytics/indexer/analysis/analyzer/custom"
	_ "github.com/knights-analytics/indexer/analysis/analyzer/keyword"
	_ "github.com/knights-analytics/indexer/analysis/analyzer/simple"
	_ "github.com/knights-analytics/indexer/analysis/analyzer/standard"
	_ "github.com/knights-analytics/indexer/analysis/analyzer/web"

	// token filters
	_ "github.com/knights-analytics/indexer/analysis/token/apostrophe"
	_ "github.com/knights-analytics/indexer/analysis/token/camelcase"
	_ "github.com/knights-analytics/indexer/analysis/token/compound"
	_ "github.com/knights-analytics/indexer/analysis/token/edgengram"
	_ "github.com/knights-analytics/indexer/analysis/token/elision"
	_ "github.com/knights-analytics/indexer/analysis/token/keyword"
	_ "github.com/knights-analytics/indexer/analysis/token/length"
	_ "github.com/knights-analytics/indexer/analysis/token/lowercase"
	_ "github.com/knights-analytics/indexer/analysis/token/ngram"
	_ "github.com/knights-analytics/indexer/analysis/token/reverse"
	_ "github.com/knights-analytics/indexer/analysis/token/shingle"
	_ "github.com/knights-analytics/indexer/analysis/token/stop"
	_ "github.com/knights-analytics/indexer/analysis/token/truncate"
	_ "github.com/knights-analytics/indexer/analysis/token/unicodenorm"
	_ "github.com/knights-analytics/indexer/analysis/token/unique"

	// tokenizers
	_ "github.com/knights-analytics/indexer/analysis/tokenizer/exception"
	_ "github.com/knights-analytics/indexer/analysis/tokenizer/regexp"
	_ "github.com/knights-analytics/indexer/analysis/tokenizer/single"
	_ "github.com/knights-analytics/indexer/analysis/tokenizer/unicode"
	_ "github.com/knights-analytics/indexer/analysis/tokenizer/web"
	_ "github.com/knights-analytics/indexer/analysis/tokenizer/whitespace"

	// date time parsers
	_ "github.com/knights-analytics/indexer/analysis/datetime/flexible"
	_ "github.com/knights-analytics/indexer/analysis/datetime/iso"
	_ "github.com/knights-analytics/indexer/analysis/datetime/optional"
	_ "github.com/knights-analytics/indexer/analysis/datetime/percent"
	_ "github.com/knights-analytics/indexer/analysis/datetime/sanitized"
	_ "github.com/knights-analytics/indexer/analysis/datetime/timestamp/microseconds"
	_ "github.com/knights-analytics/indexer/analysis/datetime/timestamp/milliseconds"
	_ "github.com/knights-analytics/indexer/analysis/datetime/timestamp/nanoseconds"
	_ "github.com/knights-analytics/indexer/analysis/datetime/timestamp/seconds"

	// languages
	_ "github.com/knights-analytics/indexer/analysis/lang/ar"
	_ "github.com/knights-analytics/indexer/analysis/lang/bg"
	_ "github.com/knights-analytics/indexer/analysis/lang/ca"
	_ "github.com/knights-analytics/indexer/analysis/lang/cjk"
	_ "github.com/knights-analytics/indexer/analysis/lang/ckb"
	_ "github.com/knights-analytics/indexer/analysis/lang/cs"
	_ "github.com/knights-analytics/indexer/analysis/lang/da"
	_ "github.com/knights-analytics/indexer/analysis/lang/de"
	_ "github.com/knights-analytics/indexer/analysis/lang/el"
	_ "github.com/knights-analytics/indexer/analysis/lang/en"
	_ "github.com/knights-analytics/indexer/analysis/lang/es"
	_ "github.com/knights-analytics/indexer/analysis/lang/eu"
	_ "github.com/knights-analytics/indexer/analysis/lang/fa"
	_ "github.com/knights-analytics/indexer/analysis/lang/fi"
	_ "github.com/knights-analytics/indexer/analysis/lang/fr"
	_ "github.com/knights-analytics/indexer/analysis/lang/ga"
	_ "github.com/knights-analytics/indexer/analysis/lang/gl"
	_ "github.com/knights-analytics/indexer/analysis/lang/hi"
	_ "github.com/knights-analytics/indexer/analysis/lang/hr"
	_ "github.com/knights-analytics/indexer/analysis/lang/hu"
	_ "github.com/knights-analytics/indexer/analysis/lang/hy"
	_ "github.com/knights-analytics/indexer/analysis/lang/id"
	_ "github.com/knights-analytics/indexer/analysis/lang/in"
	_ "github.com/knights-analytics/indexer/analysis/lang/it"
	_ "github.com/knights-analytics/indexer/analysis/lang/nl"
	_ "github.com/knights-analytics/indexer/analysis/lang/no"
	_ "github.com/knights-analytics/indexer/analysis/lang/pl"
	_ "github.com/knights-analytics/indexer/analysis/lang/pt"
	_ "github.com/knights-analytics/indexer/analysis/lang/ro"
	_ "github.com/knights-analytics/indexer/analysis/lang/ru"
	_ "github.com/knights-analytics/indexer/analysis/lang/sv"
	_ "github.com/knights-analytics/indexer/analysis/lang/tr"

	// kv stores
	_ "github.com/knights-analytics/indexer/index/upsidedown/store/boltdb"
	_ "github.com/knights-analytics/indexer/index/upsidedown/store/goleveldb"
	_ "github.com/knights-analytics/indexer/index/upsidedown/store/gtreap"
	_ "github.com/knights-analytics/indexer/index/upsidedown/store/moss"

	// index types
	_ "github.com/knights-analytics/indexer/index/upsidedown"
)
