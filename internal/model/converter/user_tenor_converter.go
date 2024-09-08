package converter

import (
	"github/golang-developer-technical-test/internal/entity"
	"github/golang-developer-technical-test/internal/model"

	"github.com/google/uuid"
)

func UserTenorToEntity(req model.UserLimitTenor, userId uuid.UUID, createBy uuid.UUID, updateBy uuid.UUID) (*entity.MapUserTenor, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	return &entity.MapUserTenor{
		PkMapUserTenor: id,
		FkMsUser:       userId,
		FkMsTenor:      req.TenorId,
		Amount:         req.Amount,
		Stamp: entity.Stamp{
			CreatedBy: createBy.String(),
			UpdatedBy: updateBy.String(),
		},
	}, nil
}

func UserTenorsToEntities(req []model.UserLimitTenor, userId uuid.UUID, createBy uuid.UUID, updateBy uuid.UUID) ([]entity.MapUserTenor, error) {
	records := make([]entity.MapUserTenor, 0)
	for _, val := range req {
		record, err := UserTenorToEntity(val, userId, createBy, updateBy)
		if err != nil {
			return records, err
		}
		records = append(records, *record)
	}
	return records, nil
}
