package businessobjects_test

import (
	"context"
)

const stateInitial = "initial"
const stateStaged = "staged"
const statePublished = "published"

func validTokenFromContext(ctx context.Context) (string, error) {
	return "abc", nil
}

func validUriFromContext(ctx context.Context) (string, error) {
	return "abc", nil
}
