package database

import (
	"fmt"

	"github.com/adaickalavan/Go-Rest-Kafka-Mongo/document"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//Dictionary contains server and database strings
type Dictionary struct {
	connection
}

//EnsureIndex creates an index field in the collection
func (dictionary *Dictionary) EnsureIndex(fields []string) {
	//Ensure index in MongoDB
	index := mgo.Index{
		Key:        fields, //Index key fields; prefix name with (-) dash for descending order
		Unique:     true,   //Prevent two documents from having the same key
		DropDups:   true,   //Drop documents with same index
		Background: true,   //Build index in background and return immediately
		Sparse:     true,   //Only index documents containing the Key fields
	}
	err := dictionary.c.EnsureIndex(index)
	checkError(err)
}

//FindAll retrieves all doc by its Value from dictionary
func (dictionary *Dictionary) FindAll() ([]document.Word, error) {
	var docs []document.Word
	err := dictionary.c.Find(bson.M{}).All(&docs)
	return docs, err
}

//FindByValue retrieves the doc by its Value from dictionary
func (dictionary *Dictionary) FindByValue(value string) (document.Word, error) {
	var doc document.Word
	err := dictionary.c.Find(bson.M{"value": value}).One(&doc)
	return doc, err
}

//Insert inserts the doc into the dictionary
func (dictionary *Dictionary) Insert(doc document.Word) error {
	err := dictionary.c.Insert(&doc)
	return err
}

//Delete deletes the doc from dictionary
func (dictionary *Dictionary) Delete(doc document.Word) error {
	err := dictionary.c.Remove(&doc)
	return err
}

func checkError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
		return true
	}
	return false
}
