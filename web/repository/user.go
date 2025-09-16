package repository

import (
	"github.com/sreio/gold/web/common"
	"github.com/sreio/gold/web/dto"
	"github.com/sreio/gold/web/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo { return &UserRepo{db: db} }

func (r *UserRepo) List(q dto.QueryUser) (items []model.User, total int64, err error) {
	base := r.db.Model(&model.User{})
	base = model.FilterUsers(q)(base)

	err = base.Count(&total).Error
	if err != nil {
		return
	}

	err = base.Scopes(common.Paginate(q.Page, q.Size)).
		Preload("UserConf").
		Order("id DESC").
		Find(&items).Error
	return
}

func (r *UserRepo) GetByID(id uint) (*model.User, error) {
	var u model.User
	if err := r.db.Preload("UserConf").First(&u, id).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *UserRepo) CreateWithConf(u *model.User, conf []dto.UserConfDTO) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(u).Error; err != nil {
			return err
		}
		if len(conf) == 0 {
			return nil
		}

		rows := make([]model.UserConf, 0, len(conf))
		for _, c := range conf {
			rows = append(rows, model.UserConf{
				UserID: u.ID, Type: c.Type, Key: c.Key, Value: c.Value,
			})
		}

		// 幂等：基于唯一索引 (user_id, type, key) 做 UPSERT
		return tx.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "user_id"}, {Name: "type"}, {Name: "key"}},
			DoUpdates: clause.AssignmentColumns([]string{"value", "updated_at"}),
		}).Create(&rows).Error
	})
}

// UpdateAndSyncConf 覆盖式更新：基础信息可选更新 + 完整同步 user_conf
func (r *UserRepo) UpdateAndSyncConf(id uint, dto dto.UpdateUserDTO) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		var u model.User
		if err := tx.First(&u, id).Error; err != nil {
			return err
		}

		// 局部字段更新（PATCH）
		updates := map[string]any{}
		if dto.Name != nil {
			updates["name"] = *dto.Name
		}
		if dto.Cron != nil {
			updates["cron"] = *dto.Cron
		}
		if dto.SaveDay != nil {
			updates["save_day"] = *dto.SaveDay
		}
		if len(updates) > 0 {
			if err := tx.Model(&u).Updates(updates).Error; err != nil {
				return err
			}
		}

		// 同步子表：策略 = Upsert + 删除缺失
		if dto.UserConf != nil {
			// 1) Upsert 全量 payload
			if len(dto.UserConf) > 0 {
				rows := make([]model.UserConf, 0, len(dto.UserConf))
				for _, c := range dto.UserConf {
					rows = append(rows, model.UserConf{
						UserID: u.ID, Type: c.Type, Key: c.Key, Value: c.Value,
					})
				}
				if err := tx.Clauses(clause.OnConflict{
					Columns:   []clause.Column{{Name: "user_id"}, {Name: "type"}, {Name: "key"}},
					DoUpdates: clause.AssignmentColumns([]string{"value", "updated_at"}),
				}).Create(&rows).Error; err != nil {
					return err
				}
			}

			// 2) 删除“未在 payload 中出现”的旧记录（按 type+key 判断）
			// 若 payload 为空数组，下面的 NOT IN 会删除该用户所有 conf
			type pair struct{ Type, Key string }
			keep := make([]pair, 0, len(dto.UserConf))
			for _, c := range dto.UserConf {
				keep = append(keep, pair{c.Type, c.Key})
			}

			// 构造 NOT IN 条件
			if len(keep) == 0 {
				if err := tx.Where("user_id = ?", u.ID).Delete(&model.UserConf{}).Error; err != nil {
					return err
				}
			} else {
				// 构建 (type, key) NOT IN ((?, ?), (?, ?), ...)
				// GORM 没有直接的 tuple IN helper，这里用子查询规避
				// 简化实现：逐步删除
				var curr []model.UserConf
				if err := tx.Where("user_id = ?", u.ID).Find(&curr).Error; err != nil {
					return err
				}
				toKeep := map[string]struct{}{}
				for _, p := range keep {
					toKeep[p.Type+"|"+p.Key] = struct{}{}
				}
				var ids []uint
				for _, c := range curr {
					if _, ok := toKeep[c.Type+"|"+c.Key]; !ok {
						ids = append(ids, c.ID)
					}
				}
				if len(ids) > 0 {
					if err := tx.Where("id IN ?", ids).Delete(&model.UserConf{}).Error; err != nil {
						return err
					}
				}
			}
		}
		return nil
	})
}

func (r *UserRepo) Delete(id uint) error {
	// 如果你使用软删除，并希望子表一并删除，可在事务里手动删子表
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("user_id = ?", id).Delete(&model.UserConf{}).Error; err != nil {
			return err
		}
		if err := tx.Delete(&model.User{}, id).Error; err != nil {
			return err
		}
		return nil
	})
}
