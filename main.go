package main

import (
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

	fileOpenButton := dialogScreen(w)
	
	
	fileOpenButton.Resize(fyne.NewSize(150,30))
	fileOpenButton.Move(fyne.NewPos(40, 200))

	tabs := container.NewAppTabs(container.NewTabItem("Load New File", fileOpenButton))

	for _, filePath := range fileList { 
		tabs.Append(container.NewTabItem("Load New File", widget.NewLabel(filePath)))
	}

	w.Resize(fyne.NewSize(640, 460))
	w.SetContent(tabs);
	w.ShowAndRun()

}



