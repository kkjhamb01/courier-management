package pubsub

type Event struct {
	Data  interface{}
	Topic string
}

func (e *Event) OnboardingResultData() financePb.OnboardingResult {
	return e.Data.(financePb.OnboardingResult)
}

func OnboardingResult(onboardingResult financePb.OnboardingResult) Event {
	return Event{
		Data:  onboardingResult,
		Topic: TopicOnboardingResult,
	}
}
