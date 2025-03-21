package projection

import (
	"context"
	"fmt"

	opensearch_repo "github.com/openshift-assisted/assisted-events-streams/internal/repository/opensearch"
	redis_repo "github.com/openshift-assisted/assisted-events-streams/internal/repository/redis"
	kafka "github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

func NewEnrichedEventsProjectionFromEnv(ctx context.Context, logger *logrus.Logger, ackChannel chan kafka.Message) (*EnrichedEventsProjection, error) {
	enrichedEventRepository := opensearch_repo.NewEnrichedEventRepositoryFromEnv(logger, ackChannel)
	snapshotRepository, err := redis_repo.NewSnapshotRepositoryFromEnv(ctx, logger)
	if err != nil {
		return nil, fmt.Errorf("failed to create snapshot repository: %w", err)
	}

	return NewEnrichedEventsProjection(
		logger,
		snapshotRepository,
		enrichedEventRepository,
		ackChannel,
	)
}
