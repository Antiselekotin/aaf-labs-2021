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

type Print struct {
	CollectionName string
}

// Visitor

type Where interface {

}

type WhereNone struct {

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