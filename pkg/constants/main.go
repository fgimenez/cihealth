package constants

const (
	DateFormat = "2006-01-02T15:04:05Z"

	LGTMLabel              = "lgtm"
	ApprovedLabel          = "approved"
	HoldLabel              = "do-not-merge/hold"
	WorkInProgressLabel    = "do-not-merge/work-in-progress"
	DoNotMergeLabelPattern = "do-not-merge/*"

	NeedsRebaseLabel = "needs-rebase"

	MergeQueueLengthName = "AverageMergeQueueLength"
	TimeToMergeName      = "AverageTimeToMerge"

	DefaultPath                        = "/tmp/test"
	DefaultTokenPath                   = ""
	DefaultSource                      = "kubevirt/kubevirt"
	DefaultDataDays                    = 7
	DefaultLogLevel                    = "info"
	DefaultTimeToMergeYellowLevel      = 1.0
	DefaultTimeToMergeRedLevel         = 3.0
	DefaultMergeQueueLengthYellowLevel = 8.0
	DefaultMergeQueueLengthRedLevel    = 15.0

	TimeToMergeBadgeFileName      = "time-to-merge.svg"
	MergeQueueLengthBadgeFileName = "merge-queue-length.svg"
	JSONResultsFileName           = "results.json"
	TimeToMergeBadgeName          = "avg time to merge"
	MergeQueueLengthBadgeName     = "avg merge queue length"
	BadgeDataFormat               = "%.2f ± std %.2f"
)

func DoNotMergeLabels() []string {
	return []string{
		NeedsRebaseLabel,
		HoldLabel,
		WorkInProgressLabel,
		DoNotMergeLabelPattern,
	}
}

func RequiredForMergeLabels() []string {
	return []string{
		LGTMLabel,
		ApprovedLabel,
	}
}
