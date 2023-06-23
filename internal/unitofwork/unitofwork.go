package unitofwork

type DomainObject interface{}
type Database interface{}

type UnitOfWorkInterface interface {
	RegisterNew(newObject DomainObject)
	RegisterDirty(modifiedObject DomainObject)
	RegisterClean() //Optional function
	RegisterDeleted(deletedObject DomainObject)
	Commit()
	Rollback()
}

func New(database *Database) *UnitOfWork {
	return &UnitOfWork{
		database: database,
	}
}

type (
	Worker interface {
		RegisterNew(newObject DomainObject)
		RegisterDirty(modifiedObject DomainObject)
		RegisterDeleted(deletedObject DomainObject)
		Commit()
	}

	UnitOfWork struct {
		database        *Database
		newObjects      []DomainObject
		modifiedObjects []DomainObject
		deletedObjects  []DomainObject
	}
)

func (u *UnitOfWork) RegisterNew(newObject DomainObject) {
	// Validate domain object
	// ...
	u.newObjects = append(u.newObjects, newObject)
}

func (u *UnitOfWork) RegisterDirty(modifiedObject DomainObject) {
	// Validate domain object
	// ...
	u.modifiedObjects = append(u.modifiedObjects, modifiedObject)
}

func (u *UnitOfWork) RegisterDeleted(deletedObject DomainObject) {
	// Validate domain object
	// ...
	u.deletedObjects = append(u.deletedObjects, deletedObject)
}

func (u UnitOfWork) Commit() {
	// Handles the actual persistence logic
}
