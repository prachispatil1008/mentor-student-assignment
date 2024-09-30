package repository

import (
	"fmt"
	"mentor-student-assignment/internal/db"
	"mentor-student-assignment/models"
)

func GetMentors() ([]models.Mentor, error) {
	rows, err := db.DB.Query("SELECT uid, name, capacity FROM mentors")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var mentors []models.Mentor
	for rows.Next() {
		var mentor models.Mentor
		err := rows.Scan(&mentor.UID, &mentor.Name, &mentor.Capacity)

		if err != nil {
			return nil, err
		}
		mentors = append(mentors, mentor)
	}
	decreaseMentorCapacity("1")
	return mentors, nil

}

func decreaseMentorCapacity(mentorUID string) error {
	query := "UPDATE mentors SET capacity = capacity - 1 WHERE uid = $1"
	_, err := db.DB.Exec(query, mentorUID)
	if err != nil {
		return err
	}
	fmt.Printf("Decreased capacity of mentor %s\n", mentorUID)
	return nil
}
