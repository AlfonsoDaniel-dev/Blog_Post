package migrations

type MigrateService interface {
	Migrate() error
}

type Migrator struct {
	ms MigrateService
}

func NewMigrator(MS MigrateService) Migrator {
	return Migrator{
		ms: MS,
	}
}

func (m *Migrator) MigrateService() error {
	return m.ms.Migrate()
}
