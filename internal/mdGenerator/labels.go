package mdGenerator

func getLabels(lang string) labels {
	switch lang {
	case "en":
		return enLabels
	case "es":
		return esLabels
	case "fr":
		return frLabels
	case "de":
		return deLabels
	case "ua":
		return uaLabels
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

var esLabels = labels{
	title:        "TODO en el proyecto",
	lastUpdate:   "Última actualización",
	overview:     "Resumen",
	total:        "Total",
	opened:       "Abierto",
	done:         "Hecho",
	withDeadline: "Con fecha límite",
	withReminder: "Con recordatorio",
	withAssignee: "Con asignado",

	general: "General",
	scope:   "Ámbito",

	deadline: "Fecha límite",
	reminder: "Recordatorio",
	assignee: "Responsable",
	file:     "Archivo",
}

var frLabels = labels{
	title:        "TODO dans le projet",
	lastUpdate:   "Dernière mise à jour",
	overview:     "Aperçu",
	total:        "Total",
	opened:       "Ouvert",
	done:         "Terminé",
	withDeadline: "Avec échéance",
	withReminder: "Avec rappel",
	withAssignee: "Avec responsable",

	general: "Général",
	scope:   "Portée",

	deadline: "Date limite",
	reminder: "Rappel",
	assignee: "Responsable",
	file:     "Fichier",
}

var deLabels = labels{
	title:        "TODO im Projekt",
	lastUpdate:   "Letzte Aktualisierung",
	overview:     "Übersicht",
	total:        "Gesamt",
	opened:       "Offen",
	done:         "Erledigt",
	withDeadline: "Mit Frist",
	withReminder: "Mit Erinnerung",
	withAssignee: "Mit Verantwortlichem",

	general: "Allgemein",
	scope:   "Bereich",

	deadline: "Frist",
	reminder: "Erinnerung",
	assignee: "Verantwortlicher",
	file:     "Datei",
}

var uaLabels = labels{
	title:        "TODO в проєкті",
	lastUpdate:   "Останнє оновлення",
	overview:     "Огляд",
	total:        "Всього",
	opened:       "Відкрите",
	done:         "Завершено",
	withDeadline: "З терміном виконання",
	withReminder: "З нагадуванням",
	withAssignee: "З виконавцем",

	general: "Загальне",
	scope:   "Область",

	deadline: "Термін виконання",
	reminder: "Нагадування",
	assignee: "Виконавець",
	file:     "Файл",
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
	scope:   "Область",

	deadline: "Дедлайн",
	reminder: "Напоминание",
	assignee: "Ответственный",
	file:     "Файл",
}
