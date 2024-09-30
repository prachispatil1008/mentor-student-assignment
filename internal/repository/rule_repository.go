package repository

import (
	"mentor-student-assignment/internal/db"
	"mentor-student-assignment/models"
)

func GetRules() ([]models.Rule, error) {
	rows, err := db.DB.Query("SELECT id, name, exp, groupid, priority FROM rules ORDER BY priority")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rules []models.Rule
	for rows.Next() {
		var rule models.Rule
		err := rows.Scan(&rule.ID, &rule.Name, &rule.Exp, &rule.GroupID, &rule.Priority)
		if err != nil {
			return nil, err
		}
		rules = append(rules, rule)
	}
	return rules, nil
}
