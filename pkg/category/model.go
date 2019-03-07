package category

type Category struct {
	ID string           `xorm:"varchar(100) notnull pk 'uuid'"`
	Name string			`xorm:"varchar(100) notnull pk 'name'"`
	Description string	`xorm:"varchar(100) notnull pk 'des'"`
	Tag string			`xorm:"varchar(100) notnull pk 'tag'"`
}

