package query

import (
	"errors"

	"github.com/expanse-agency/tycho/utils/helpers"
)

// ["relation", "otherrelation"]
type Relation []string

func ParseRelation(raw string) (*Relation, error) {
	relation, err := helpers.Unmarshal[Relation](raw)
	if err != nil {
		return nil, err
	}

	if relation == nil {
		return nil, errors.New("relation is nil")
	}

	return relation, nil
}
