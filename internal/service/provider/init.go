package provider

func NewProviderService(api ProviderInterface) *ProviderService {
	return &ProviderService{
		api: api,
	}
}
