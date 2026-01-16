package af_sailor_agent

import (
	"context"
)

type AFSailorAgent interface {
	SailorGetSessionEngine(ctx context.Context, req map[string]any) (*SailorSessionResp, error)
}
