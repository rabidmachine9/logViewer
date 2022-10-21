package main

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var fileList []string = textLinesToSlice("storage.txt");
var storageFile = "storage.txt"


func main(){
	
	a := app.NewWithID("log.viewer")
	w := a.NewWindow("Log Viewer")
	w.SetMainMenu(makeMenu(a, w))
	fileOpenButton := dialogScreen(w)
	
	
	fileOpenButton.Resize(fyne.NewSize(150,30))
	fileOpenButton.Move(fyne.NewPos(40, 200))

	tabs := container.NewDocTabs(container.NewTabItem("Load New File", fileOpenButton))

	for _, filePath := range fileList { 
		filePathSplit := strings.Split(filePath, "/")
		filename := filePathSplit[len(filePathSplit) - 1]
		tabs.Append(container.NewTabItem(filename, widget.NewLabel(filePath)))
	}

	w.Resize(fyne.NewSize(640, 460))
	w.SetContent(tabs);
	w.ShowAndRun()

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