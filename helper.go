package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)



func dialogScreen(win fyne.Window) fyne.CanvasObject {
	return widget.NewButton("Open Log File", func() {
		fileDialog(win)
	})
}

func fileDialog(win fyne.Window) {
	dialog.ShowFileOpen(func(file fyne.URIReadCloser, err error) {
		if err != nil {
			dialog.ShowError(err, win)
			return
		}
		if file == nil {
			log.Println("Cancelled")
			return
		}
		fmt.Printf("%s",file.URI().Path())
		
		addToFileList(&fileList, file.URI().Path() )
		sliceToText(fileList, storageFile)
	}, win)
	

}

func addToFileList(list *[]string, filePath string) {
	*list = append(*list, filePath)
}


func removeFileFromList(list  []string, filePath string) []string {
	fileIndex := SliceIndex(len(list), func(i int) bool { return list[i] == filePath })
	return remove(list, fileIndex)
}


func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func SliceIndex(limit int, predicate func(i int) bool) int {
	for i := 0; i < limit; i++ {
			if predicate(i) {
					return i
			}
	}
	return -1
}


func textLinesToSlice(filename string) [] string {
	lines := [] string {};
	f, err := os.Open(filename)
	if err != nil {
			panic(err)
	}

	defer f.Close()
	r := bufio.NewReader(f)

	for line, prefix, err := r.ReadLine(); err != io.EOF; line, prefix, err = r.ReadLine() {
			if prefix {
				 lines=  append(lines, string(line))
					
			} else {
					lines = append(lines, string(line))
			}
			
	}

	return lines;
}


func sliceToText(lines []string, filename string){
	f, err := os.Create(filename) 
	if err != nil {
		panic(err)
	}

	for _, line := range lines {
    // element is the element from someSlice for where we are
		f.WriteString(line)
	}
	 
	f.Close()
}


func updateFileList(sliceOld []string, sliceNew []string) []string {
	if(len(sliceOld) > len(sliceNew)){
		return sliceOld
	}else {
		return sliceNew
	}
		
}


func addNewTab(tab *container.TabItem, content fyne.CanvasObject){

}