package message

// # https://docs.aws.amazon.com/codebuild/latest/APIReference/API_BuildPhase.html
//
type BuildPhasesIcon string

const (
	BuildPhasesUnknown    = ":grey_question:"
	BuildPhasesFailed     = ":x:"
	BuildPhasesFault      = ":exclamation:"
	BuildPhasesInProgress = ":building_construction:"
	BuildPhasesQueued     = ":building_construction:"
	BuildPhasesStopped    = ":octagonal_sign:"
	BuildPhasesSucceeded  = ":white_check_mark:"
	BuildPhasesTimedOut   = ":stop_watch:"
)
