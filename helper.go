package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"github.com/atotto/clipboard"
)

//open file button
func dialogScreen(win fyne.Window) fyne.CanvasObject {
	return widget.NewButton("Open Log File", func() {
		fileDialog(win)
	})
}

//open file dialog
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
		
		addToFileList(fileMap, file.URI().Path() )
		sliceToText(fileMap, storageFile)
	}, win)
}
//add filepath to storage
func addToFileList(list map[string]string, filePath string) {
	filename := filenameFromPath(filePath)
	list[filename] = filePath
}

//remove filepath from storage
func removeFileFromList(list  []string, filePath string) []string {
	fileIndex := SliceIndex(len(list), func(i int) bool { return list[i] == filePath })
	return remove(list, fileIndex)
}

//remove slice element by index
func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}
//return slice index from value
func SliceIndex(limit int, predicate func(i int) bool) int {
	for i := 0; i < limit; i++ {
			if predicate(i) {
					return i
			}
	}
	return -1
}

//returns a map with filenames as keys and filepaths as values
func textLinesToMap(filename string) map[string] string {
	lines := make(map[string]string)
	f, err := os.Open(filename)
	if err != nil {
			panic(err)
	}

	defer f.Close()
	r := bufio.NewReader(f)

	for line, prefix, err := r.ReadLine(); err != io.EOF; line, prefix, err = r.ReadLine() {
			name := filenameFromPath(string(line))
			if prefix {	 
				lines[name] = string(line);
			} else {
				lines[name] = string(line);
			}
			
	}

	return lines;
}

//slice elements become lines for textfile
func sliceToText(lines map[string]string, filename string){
	f, err := os.Create(filename) 
	if err != nil {
		panic(err)
	}

	for _, path := range lines {
    // element is the element from someSlice for where we are
		f.WriteString(path+"\n")
	}
	 
	f.Close()
}

//lines of text -> slice elements
func textLinesToSlice(filename string) [] string {
	lines := [] string {};
	f, err := os.Open(filename)
	if err != nil {
			panic(err)
	}

	defer f.Close()
	r := bufio.NewReader(f)

	for line, _, err := r.ReadLine(); err != io.EOF; line, _, err = r.ReadLine() {
		if string(line) == "" {
			continue
		}
		lines=  append(lines, strings.TrimSuffix(string(line), "\n"))
	}

	return lines;
}

//process lines before sending them to tab
func getLinesForTab(filename string, linesNum int) []string {
	lines := textLinesToSlice(filename)
	reverse(lines)
	if(len(lines) < linesNum){
		return lines
	}else {
		return lines[0:(linesNum-1)]
	}
}

//reverse a slice
func reverse[S ~[]E, E any](s S)  {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
			s[i], s[j] = s[j], s[i]
	}
}


func createNewTab(tabs *container.DocTabs,filePath string, logs map[string][]string, widgetLists map[string]*widget.List) {
	logs[filePath] = getLinesForTab(filePath,20)
	widgetLists[filePath] =  widget.NewList(
		func() int {
				return len(logs[filePath])
		},
		func() fyne.CanvasObject {
			return container.NewHBox (
				widget.NewLabel(filePath),
				layout.NewSpacer(),
				widget.NewButton("Copy", nil),
			)
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*fyne.Container).Objects[0].(*widget.Label).SetText((logs[filePath])[i])
			o.(*fyne.Container).Objects[1].(*layout.Spacer).Show()

			o.(*fyne.Container).Objects[2].(*widget.Button).OnTapped = func() {
				clipboard.WriteAll(logs[filePath][i])
			}				
		},
	)
	tabs.Append(container.NewTabItem(filePath, widgetLists[filePath]))
}

func makeMenu(a fyne.App, w fyne.Window) *fyne.MainMenu { 
	fileItem := fyne.NewMenuItem("File", func() {  
		fileDialog(w)
	})
	
	file := fyne.NewMenu("File", fileItem)

	main := fyne.NewMainMenu(
		file,
	)
	 
	return main
}


func filenameFromPath(path string) string{
	filePathSplit := strings.Split(path, "/")

	return filePathSplit[len(filePathSplit) - 1]
}