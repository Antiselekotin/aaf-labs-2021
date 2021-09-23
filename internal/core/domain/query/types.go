package query

type Query interface {

}

type Create struct {
	Name string
}

type Insert struct {
	CollectionName string
	Content string

}

type Search struct {
	CollectionName string
	Where
}

// Visitor

type Where interface {

}

type WhereWord struct {
	Word string
}

type WherePrefix struct {
	Prefix string
}

type WhereInterval struct {
	FirstWord, LastWord string
	Interval int
}