package main

import (
	"fmt"
	"reflect"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

//global variables
var fileMap = make(map[string]string)
var storageFile = "storage.txt"


func main(){
	//var myTabsInfo = make(map[string]LogTab)
	var logs = make(map[string][]string)
	var numberOfLines = 20
	fileMap = textLinesToMap(storageFile)
	fmt.Println("file Map:",fileMap )
	a := app.NewWithID("log.viewer")
	w := a.NewWindow("Log Viewer")
	w.SetMainMenu(makeMenu(a, w))
	fileOpenButton := dialogScreen(w)
	
	fileOpenButton.Resize(fyne.NewSize(150,30))
	fileOpenButton.Move(fyne.NewPos(40, 200))

	tabs := container.NewDocTabs(container.NewTabItem("Load New File", fileOpenButton))

	var widgetLists = make(map[string]*widget.List)


	for _, filePath := range fileMap {
		createNewTab(tabs, filePath, logs, widgetLists)
	}
	
	
	go func() {
		for range time.Tick(time.Second) {
			for _, filePath := range fileMap { 
				newLog := getLinesForTab(filePath,numberOfLines)
				if _, exists := logs[filePath]; !exists {
					createNewTab(tabs, filePath, logs, widgetLists)
				} else if(!reflect.DeepEqual(newLog, logs[filePath])){
					logs[filePath] = getLinesForTab(filePath,numberOfLines)
					widgetLists[filePath].Refresh()
				}	
			}
		}
	}()
	

	tabs.OnClosed = func(tab *container.TabItem) {
		//first delete buffer from logs
		delete(logs, fileMap[tab.Text])
		//then delete the filepath from the map
		delete(fileMap, tab.Text)
		sliceToText(fileMap, storageFile)
	}
	w.Resize(fyne.NewSize(640, 460))
	w.SetContent(tabs);
	w.ShowAndRun()

}



