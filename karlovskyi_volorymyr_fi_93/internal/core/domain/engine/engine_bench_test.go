package engine

import (
	"fmt"
	"labdb/internal/core/domain/query"
	"math/rand"
	"strings"
	"testing"
)

var db = New()
var colName = "test_collection"
var baseSize = 1000
var prefixSearch = "s"
var contentTest = []string{
	"scaremonger",
	"dextrogastria",
	"scaremonger",
	"famulus",
	"connected",
	"involved",
	"with",
	"the",
	"pursuit",
	"of",
	"knowledge",
	"especially",
	"scholarly",
	"nature",
	"learned",
	"journal",
}

func init()  {
	create := query.Create{Name: colName}
	db.Create(create)
	insert := query.Insert{CollectionName: colName, Content: ""}
	for i := 1; i < baseSize; i++ {
		strArr := []string{}
		for j := 0; j < i; j++ {
			index := rand.Int() % (len(contentTest) - 1)
			strArr = append(strArr, contentTest[index])
		}
		insert.Content = strings.Join(strArr, " ")
		db.Insert(insert)
	}
}

func BenchmarkDatabase_Create(b *testing.B) {
	create := query.Create{Name: ""}
	for i := 0; i < b.N; i++ {
		create.Name = fmt.Sprintf("%v_%v", colName, i)
		db.Create(create)
	}
}



func BenchmarkDatabase_SearchWhereNone(b *testing.B) {
	search := query.Search{CollectionName: colName, Where: &query.WhereNone{}}
	for i := 0; i < b.N; i++ {
		db.Search(search)
	}
}

func BenchmarkDatabase_SearchWhereWord(b *testing.B) {
	search := query.Search{CollectionName: colName, Where: &query.WhereWord{Word: contentTest[0]}}
	for i := 0; i < b.N; i++ {
		db.Search(search)
	}
}

func BenchmarkDatabase_SearchWhereInterval(b *testing.B) {
	search := query.Search{
		CollectionName: colName,
		Where: &query.WhereInterval{
			FirstWord: contentTest[0],
			LastWord: contentTest[1],
			Interval: 3}}
	for i := 0; i < b.N; i++ {
		db.Search(search)
	}
}

func BenchmarkDatabase_SearchWherePrefix(b *testing.B) {
	search := query.Search{
		CollectionName: colName,
		Where: &query.WherePrefix{
			Prefix: prefixSearch}}
	for i := 0; i < b.N; i++ {
		db.Search(search)
	}
}

func BenchmarkDatabase_Insert(b *testing.B) {
	insert := query.Insert{CollectionName: colName, Content: "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum."}
	for i := 0; i < b.N; i++ {
		db.Insert(insert)
	}
}