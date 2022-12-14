package reporting

type RedisClient interface {
	GetArray(key string, out *[]string) error
}
type Service struct {
	client RedisClient
}

func NewService(client RedisClient) *Service {
	return &Service{client: client}
}
func (s *Service) Report(sender, receiver string) ([]Report, error) {
	reportsStr := make([]string, 0)
	reports := make([]Report, len(reportsStr))
	err := s.client.GetArray(sender, &reportsStr)
	if err != nil {
		return nil, err
	}
	for _, item := range reportsStr {
		reportDTO := ConstructReportDTOFromStr(item)
		if reportDTO.IsMessageSentTo(receiver) {
			reports = append(reports, Report{
				Sender:   sender,
				Receiver: reportDTO.Receiver,
				Message:  reportDTO.Message,
			})
		}
	}
	return reports, nil
}
