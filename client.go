package api

type GarminClient struct {
	Activity    activity
	Sleep       sleep
	Hrv         hrv
	UserSummary userSummary
}

func NewGarminClient(token *string) *GarminClient {
	auth := auth{
		token: token,
	}

	gc := GarminClient{
		Activity: &activityClient{
			auth: &auth,
		},
		Sleep: &sleepClient{
			auth: &auth,
		},
		Hrv: &hrvClient{
			auth: &auth,
		},
		UserSummary: &userSummaryClient{
			auth: &auth,
		},
	}

	return &gc
}
