package repository

import (
	"context"
	"fmt"
	"log"
	"sanatorioApp/internal/domain/cites/entities"
)

func (cr *citesRepository) RegisterOffice(ctx context.Context, of entities.Office) (string, error) {
	query := `
		INSERT INTO office (name, status_id)
		VALUES ($1, $2, $3)`

	_, err := cr.storage.DbPool.Exec(ctx, query, of.Name, entities.OfficeStatusUnassigned)
	if err != nil {
		log.Printf("error al registrar el consultorio '%s' en la db: %v", of.Name, err)
		return "", err
	}

	return fmt.Sprintf("Consultorio '%s' registrado con éxito", of.Name), nil
}

func (cr *citesRepository) UpdateOffice(ctx context.Context, of entities.Office) (string, error) {
	query := `
		UPDATE office
		SET name = $1, updated_at = $2
		WHERE id = $3
	`

	// Ejecutar la consulta
	result, err := cr.storage.DbPool.Exec(ctx, query, of.Name, of.UpdatedAt, of.ID)
	if err != nil {
		log.Printf("error al actualizar office: %v", err)
		return "", fmt.Errorf("error al ejecutar la consulta de actualización: %w", err)
	}

	// Obtener el número de filas afectadas
	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return "", fmt.Errorf("no se encontró una oficina con el id %d", of.ID)
	}

	return fmt.Sprintf("Consultorio '%s' actualizado con éxito", of.Name), nil
}

func (cr *citesRepository) GetOffices(ctx context.Context) ([]entities.Office, error) {
	query := `
		SELECT 
			o.id AS office_id,
			status_id,
			o.name AS office_name,
			st.name AS status_name
		FROM office o
		INNER JOIN office_status st
		ON o.status_id = st.id
	`

	rows, err := cr.storage.DbPool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var offices []entities.Office

	for rows.Next() {
		var office entities.Office

		err := rows.Scan(
			&office.ID,
			&office.OfficeStatus.ID,
			&office.Name,
			&office.OfficeStatus.Name,
		)
		if err != nil {
			return nil, err
		}

		offices = append(offices, office)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return offices, nil
}
