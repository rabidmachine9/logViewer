package main

import (
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

var fileList []string


type FileDialog struct {
}

func main(){
	
	a := app.NewWithID("log.viewer")
	w := a.NewWindow("Log Viewer")

	fileOpenButton := dialogScreen(w)
	
	
	fileOpenButton.Resize(fyne.NewSize(150,30))
	fileOpenButton.Move(fyne.NewPos(40, 200))

	w.Resize(fyne.NewSize(640, 460))
	w.SetContent(container.NewWithoutLayout(fileOpenButton));
	w.ShowAndRun()

}



func dialogScreen(win fyne.Window) fyne.CanvasObject {
	return widget.NewButton("Open Log File", func() {
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
			//fileSaved(writer, win)
			addToFileList(&fileList, file.URI().Path() )
		}, win)
	})
	
}

func addToFileList(list *[]string, filePath string) {
	*list = append(*list, filePath)
}


func fileSaved(f fyne.URIWriteCloser, w fyne.Window) {
	defer f.Close()
	_, err := f.Write([]byte("Written by Fyne demo\n"))
	if err != nil {
		dialog.ShowError(err, w)
	}
	err = f.Close()
	if err != nil {
		dialog.ShowError(err, w)
	}
	log.Println("Saved to...", f.URI())
}