package ca

import (
	"github.com/knights-analytics/indexer/analysis"
	"github.com/knights-analytics/indexer/registry"
)

const ArticlesName = "articles_ca"

// this content was obtained from:
// lucene-4.7.2/analysis/common/src/resources/org/apache/lucene/analysis

var CatalanArticles = []byte(`
d
l
m
n
s
t
`)

func ArticlesTokenMapConstructor(config map[string]interface{}, cache *registry.Cache) (analysis.TokenMap, error) {
	rv := analysis.NewTokenMap()
	err := rv.LoadBytes(CatalanArticles)
	return rv, err
}

func init() {
	err := registry.RegisterTokenMap(ArticlesName, ArticlesTokenMapConstructor)
	if err != nil {
		panic(err)
	}
}
