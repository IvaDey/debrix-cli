package mdGenerator

import (
	"fmt"
	"time"
)

func generateHeader(lang string, overview overview) string {
	localizedLabels := getLabels(lang)

	return fmt.Sprintf(`# 📌 %s
[![Powered by Debrix](https://img.shields.io/badge/Powered%%20by-Debrix-6E40C9?style=flat&logo=task)](https://debrix.dev)

_%s: %s_

## 📊 %s

- 🔧 %s: %d
- 🟡 %s: %d
- ✅ %s: %d
- ⏰ %s: %d
- 🔔 %s: %d
- 👤 %s: %d`,
		localizedLabels.title, localizedLabels.lastUpdate, time.Now().Format(time.DateOnly),
		localizedLabels.overview,
		localizedLabels.total, overview.Total,
		localizedLabels.opened, overview.Opened,
		localizedLabels.done, overview.Done,
		localizedLabels.withDeadline, overview.WithDeadline,
		localizedLabels.withReminder, overview.WithReminder,
		localizedLabels.withAssignee, overview.WithAssignee,
	)
}

type overview struct {
	Total        int
	Opened       int
	Done         int
	WithDeadline int
	WithReminder int
	WithAssignee int
}
