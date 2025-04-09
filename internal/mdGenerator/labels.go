package mdGenerator

func getLabels(lang string) labels {
	switch lang {
	case "en":
		return enLabels
	case "ru":
		return ruLabels
	default:
		return enLabels
	}
}

type labels struct {
	title        string
	lastUpdate   string
	overview     string
	total        string
	opened       string
	done         string
	withDeadline string
	withReminder string
	withAssignee string

	general string
	scope   string

	deadline string
	reminder string
	assignee string
	file     string
}

var enLabels = labels{
	title:        "Project TODOs",
	lastUpdate:   "Last update",
	overview:     "Overview",
	total:        "Total",
	opened:       "Opened",
	done:         "Done",
	withDeadline: "With deadline",
	withReminder: "With reminder",
	withAssignee: "With assignee",

	general: "General",
	scope:   "Scope",

	deadline: "Deadline",
	reminder: "Reminder",
	assignee: "Assignee",
	file:     "File",
}

var ruLabels = labels{
	title:        "TODO в проекте",
	lastUpdate:   "Последнее обновление",
	overview:     "Обзор",
	total:        "Всего",
	opened:       "Открыто",
	done:         "Выполнено",
	withDeadline: "С дедлайном",
	withReminder: "С напоминанием",
	withAssignee: "С исполнителем",

	general: "Общие",
	scope:   "Скоуп",

	deadline: "Дедлайн",
	reminder: "Напоминание",
	assignee: "Ответственный",
	file:     "Файл",
}

// todo: add more locales
