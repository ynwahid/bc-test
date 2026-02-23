package aggregation

func NewAggregationService(provider ProviderInterface) *AggregationService {
	return &AggregationService{
		provider: provider,
	}
}
