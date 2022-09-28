package database

import (
	"os"
)

// structure describing a database page
type Page struct {
	Number uint64
	Data   []byte
}

// structure describing the data layer
type Layer struct {
	File     *os.File
	PageSize int
	Manager  *PageManager
}

/*
New data Layer constructor
path: 		where to store the database data
pageSize: 	how large should the datbase page size be
*/
func NewLayer(path string, pageSize int) (*Layer, error) {

	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}

	dal := &Layer{
		File:     file,
		PageSize: pageSize,
		Manager:  NewPageManager(),
	}

	return dal, nil
}

/*
Allocate empty page
initialize an empty page structure
*/
func (l *Layer) AllocEmptyPage() *Page {
	return &Page{
		Data: make([]byte, l.PageSize),
	}
}

/*
Read operation for the layer
*/
func (l *Layer) ReadPage(pageNum uint64) (*Page, error) {
	newPage := l.AllocEmptyPage()

	// calculate the correct offset given the page number and page size
	offset := int(pageNum) * l.PageSize

	// read the data stored at that offset
	if _, err := l.File.ReadAt(newPage.Data, int64(offset)); err != nil {
		return nil, err
	}

	return newPage, nil
}

/*
Write operation for the layer
*/
func (l *Layer) WritePage(p *Page) error {
	// write the data to the correct page number with the page size
	offset := int64(p.Number) * int64(l.PageSize)
	_, err := l.File.WriteAt(p.Data, offset)
	return err
}

/*
Close file handle for the layer
*/
func (l *Layer) Close() error {
	if l.File != nil {
		if err := l.File.Close(); err != nil {
			return ErrorClosingFile
		}
		l.File = nil
	}
	return nil
}
