package asker

import "context"

type Asker interface {
	Ask(ctx context.Context, prompt, cluster string) (string, error)
}
