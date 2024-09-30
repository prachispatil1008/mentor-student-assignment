package repository

import (
	"mentor-student-assignment/internal/db"
	"mentor-student-assignment/models"
)

func RegisterStudent(student models.Student) error {
	query := "INSERT INTO students (sid, name, lang, standard, mentorid) VALUES ($1, $2, $3, $4, $5)"
	_, err := db.DB.Exec(query, student.SID, student.Name, student.Lang, student.Standard, student.MentorID)
	if err != nil {
		return err
	}
	return nil
}
