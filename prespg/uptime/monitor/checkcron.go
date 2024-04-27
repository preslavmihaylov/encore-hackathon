package monitor

import "encore.dev/cron"

// Check all tracked sites every 1 hour.
var _ = cron.NewJob("check-all", cron.JobConfig{
	Title:    "Check all sites",
	Endpoint: CheckAll,
	Every:    1 * cron.Hour,
})
