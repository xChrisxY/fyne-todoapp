package main

import (
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/widget"
)

func main() {
    myApp := app.New()
    myWindow := myApp.NewWindow("To-Do List")

    input, listContainer := createUI()

    myWindow.SetContent(container.NewVBox(
        input,
        widget.NewButton("Agregar Tarea", func() {
            addTask(input, listContainer)
        }),
        widget.NewLabel("Lista de Tareas:"),
        listContainer,
    ))

    myWindow.Resize(fyne.NewSize(400, 600))
    myWindow.ShowAndRun()
}

func createUI() (*widget.Entry, *fyne.Container) {

    input := widget.NewEntry()
    input.SetPlaceHolder("Ingresa una nueva tarea")
    listContainer := container.NewVBox()

    return input, listContainer
}

func addTask(input *widget.Entry, listContainer *fyne.Container) {
    task := input.Text
    if task != "" {
        taskLabel := widget.NewLabel(task)
        completeButton := widget.NewButton("Completada", func() {
            taskLabel.SetText(task + " - Completada")
            
        })
        completeButton.Disable()

        listContainer.Add(container.NewHBox(
            taskLabel,
            completeButton,
        ))
        input.SetText("") 
    }
}
