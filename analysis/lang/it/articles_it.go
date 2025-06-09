package it

import (
	"github.com/knights-analytics/indexer/analysis"
	"github.com/knights-analytics/indexer/registry"
)

const ArticlesName = "articles_it"

// this content was obtained from:
// lucene-4.7.2/analysis/common/src/resources/org/apache/lucene/analysis

var ItalianArticles = []byte(`
c
l
all
dall
dell
nell
sull
coll
pell
gl
agl
dagl
degl
negl
sugl
un
m
t
s
v
d
`)

func ArticlesTokenMapConstructor(config map[string]interface{}, cache *registry.Cache) (analysis.TokenMap, error) {
	rv := analysis.NewTokenMap()
	err := rv.LoadBytes(ItalianArticles)
	return rv, err
}

func init() {
	err := registry.RegisterTokenMap(ArticlesName, ArticlesTokenMapConstructor)
	if err != nil {
		panic(err)
	}
}
