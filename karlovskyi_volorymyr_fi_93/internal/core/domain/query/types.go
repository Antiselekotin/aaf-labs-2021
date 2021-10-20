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

