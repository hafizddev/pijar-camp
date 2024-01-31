package provinsimodel

import (
	"pijar_camp/config"
	"pijar_camp/entities"
)

func GetAll() (provinsis []entities.Provinsi) {
	rows, err := config.DB.Query(`SELECT * FROM provinsi`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var provinsi entities.Provinsi
		if err := rows.Scan(&provinsi.Id, &provinsi.Name, &provinsi.CreatedAt, &provinsi.UpddatedAt); err != nil {
			panic(err)
		}

		provinsis = append(provinsis, provinsi)
	}
	return
}

func Create(provinsi entities.Provinsi) (bool bool) {
	result, err := config.DB.Exec(`
		INSERT INTO provinsi (name, created_at, updated_at)
		VALUE (?, ?, ?)`,
		provinsi.Name,
		provinsi.CreatedAt,
		provinsi.UpddatedAt,
	)
	if err != nil {
		panic(err)
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	return lastInsertId > 0
}

func Detail(id int) (provinsi entities.Provinsi) {
	row := config.DB.QueryRow(`SELECT id, name FROM provinsi WHERE id = ? `, id)

	if err := row.Scan(&provinsi.Id, &provinsi.Name); err != nil {
		panic(err.Error())
	}
	return
}

func Update(id int, provinsi entities.Provinsi) (bool bool) {
	query, err := config.DB.Exec(`UPDATE provinsi SET name = ?, updated_at = ? where id = ?`, provinsi.Name, provinsi.UpddatedAt, id)
	if err != nil {
		panic(err)
	}

	result, err := query.RowsAffected()
	if err != nil {
		panic(err)
	}
	return result > 0
}

func Delete(id int) (err error) {
	_, err = config.DB.Exec("DELETE FROM provinsi WHERE id = ?", id)
	return
}
