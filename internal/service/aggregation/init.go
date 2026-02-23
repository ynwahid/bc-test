package aggregation

func NewProviderService(provider ProviderInterface) *ProviderService {
	return &ProviderService{
		provider: provider,
	}
}
