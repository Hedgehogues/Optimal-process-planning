package main

import (
	"io"

	scheduler_pkg "gitlab.ozon.ru/eurvanov/optimal-process-planning/scheduler"
)

type Processor struct {
}

func (p *Processor) AppendFiles(filesForProcess scheduler_pkg.FileSlice) {

}

func (p *Processor) Process() int {
	// Processing filesForProcess ...
	return 0
}

func main() {
	files := scheduler_pkg.FileSlice{
		scheduler_pkg.File{0, 1, 2},
		scheduler_pkg.File{1, 1, 2},
		scheduler_pkg.File{2, 1, 2},
		scheduler_pkg.File{2, 18, 5},
	}
	scheduler, err := scheduler_pkg.NewScheduler(files, 5, 3)
	if err != nil {
		panic(err)
	}
	processor := Processor{}
	filesForProcess := scheduler.GetCurrentFiles()
	// New files will appended for processing
	processor.AppendFiles(filesForProcess)
	for {
		// Processing filesForProcess ...
		finishedFileId := processor.Process()
		// Processing one file was finished
		scheduler.Finished(finishedFileId)
		newFiles, err := scheduler.Next()
		if err != nil {
			panic(err)
		}
		if err == io.EOF {
			break
		}
		// New files will appended for processing
		processor.AppendFiles(newFiles)
	}
}
