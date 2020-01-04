// Copyright 2012-present Oliver Eilhard. All rights reserved.
// Use of this source code is governed by a MIT-license.
// See http://olivere.mit-license.org/license.txt for details.

package elastic

// IntervalQueryRule represents the generic matching interval rule interface. Interval Rule is actually
// just a Query, but may be used only inside IntervalQuery. An extra method is added
// just to separate all its implementations (*Rule objects) from other query objects
type IntervalQueryRule interface {
	Query

	// isIntervalQueryRule is never actually called, and is used just for Rule to differ from standard Query
	isIntervalQueryRule() bool
}

// IntervalQuery returns documents based on the order and proximity of matching terms
//
// For more details, see
// https://www.elastic.co/guide/en/elasticsearch/reference/7.4/query-dsl-intervals-query.html
type IntervalQuery struct {
	field string
	rule  IntervalQueryRule
}

// NewIntervalQuery creates and initializes a new IntervalQuery.
func NewIntervalQuery(field string, rule IntervalQueryRule) *IntervalQuery {
	return &IntervalQuery{field: field, rule: rule}
}

// Source returns JSON for the function score query.
func (q *IntervalQuery) Source() (interface{}, error) {
	source := make(map[string]interface{})

	ruleSrc, err := q.rule.Source()
	if err != nil {
		return nil, err
	}

	source[q.field] = ruleSrc

	return source, nil
}
