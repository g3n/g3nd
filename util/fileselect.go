package util

import (
	"os"
	"path/filepath"
	"sort"

	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/gui/assets/icon"
	"github.com/g3n/engine/math32"
)

type FileSelectButton struct {
	*gui.Button
	FS       *FileSelect
	path     string
	errLabel *gui.Label
}

func NewFileSelectButton(path, text string, width, height float32) *FileSelectButton {

	// Initialize file select button
	fsb := new(FileSelectButton)
	fsb.Button = gui.NewButton(text)
	fsb.path = path

	// Creates error label
	fsb.errLabel = gui.NewLabel("")
	fsb.errLabel.SetBounded(false)
	fsb.errLabel.SetBgColor(&math32.Color{1, 1, 1})
	fsb.errLabel.SetColor(&math32.Color{1, 0, 0})
	fsb.errLabel.SetPosition(4, 40)
	fsb.errLabel.SetBorders(1, 1, 1, 1)
	fsb.errLabel.SetPaddings(2, 6, 2, 6)
	fsb.errLabel.SetFontSize(18)
	fsb.errLabel.SetVisible(false)
	fsb.Button.Add(fsb.errLabel)

	// Creates file select panel and add it to the button
	fsb.FS = NewFileSelect(width, height)
	fsb.FS.SetPosition(0, 0)
	fsb.FS.SetVisible(false)
	fsb.FS.SetBounded(false)
	fsb.Button.Add(fsb.FS)

	// Subscribe to file select panel buttons
	fsb.FS.Subscribe("OnOK", func(evname string, ev interface{}) {
		fpath := fsb.FS.Selected()
		if fpath == "" {
			fsb.FS.SetVisible(false)
			return
		}
		fsb.Button.Dispatch("OnSelect", fpath)
		fsb.FS.SetVisible(false)
	})
	// Hides file select if Cancel button is clicked
	fsb.FS.Subscribe("OnCancel", func(evname string, ev interface{}) {
		fsb.FS.SetVisible(false)
	})

	// When button is clicked shows the file selector panel
	fsb.Button.Subscribe(gui.OnClick, func(evname string, ev interface{}) {
		err := fsb.FS.SetPath(fsb.path)
		if err != nil {
			panic(err)
		}
		fsb.FS.SetVisible(true)
	})

	return fsb
}

func (fsb *FileSelectButton) SetError(text string) {

	if text == "" {
		fsb.errLabel.SetVisible(false)
		return
	}
	fsb.errLabel.SetText(text)
	fsb.errLabel.SetVisible(true)
}

type FileSelect struct {
	gui.Panel
	path        string
	pathLabel   *gui.Label
	list        *gui.List
	bok         *gui.Button
	bcan        *gui.Button
	fileFilters []string
}

func NewFileSelect(width, height float32) *FileSelect {

	fs := new(FileSelect)
	fs.Panel.Initialize(fs, width, height)
	fs.SetBorders(1, 1, 1, 1)
	fs.SetPaddings(4, 4, 4, 4)
	fs.SetColor(math32.NewColor("white"))
	fs.SetVisible(false)

	// Set vertical box layout for the whole panel
	l := gui.NewVBoxLayout()
	l.SetSpacing(4)
	fs.SetLayout(l)

	// Creates path label
	fs.pathLabel = gui.NewLabel("path")
	fs.Add(fs.pathLabel)

	// Creates list
	fs.list = gui.NewVList(0, 0)
	fs.list.SetLayoutParams(&gui.VBoxLayoutParams{Expand: 5, AlignH: gui.AlignWidth})
	fs.list.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		fs.onSelect()
	})
	fs.Add(fs.list)

	// Button container panel
	bc := gui.NewPanel(0, 0)
	bcl := gui.NewHBoxLayout()
	bcl.SetAlignH(gui.AlignWidth)
	bc.SetLayout(bcl)
	bc.SetLayoutParams(&gui.VBoxLayoutParams{Expand: 1, AlignH: gui.AlignWidth})
	fs.Add(bc)

	// Creates OK button
	fs.bok = gui.NewButton("OK")
	fs.bok.SetLayoutParams(&gui.HBoxLayoutParams{Expand: 0, AlignV: gui.AlignCenter})
	fs.bok.Subscribe(gui.OnClick, func(evname string, ev interface{}) {
		fs.Dispatch("OnOK", nil)
	})
	bc.Add(fs.bok)

	// Creates Cancel button
	fs.bcan = gui.NewButton("Cancel")
	fs.bcan.SetLayoutParams(&gui.HBoxLayoutParams{Expand: 0, AlignV: gui.AlignCenter})
	fs.bcan.Subscribe(gui.OnClick, func(evname string, ev interface{}) {
		fs.Dispatch("OnCancel", nil)
	})
	bc.Add(fs.bcan)

	// Sets initial directory
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	} else {
		fs.SetPath(path)
	}
	return fs
}

func (fs *FileSelect) SetPath(path string) error {

	// Open path file or dir
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	// Checks if it is a directory
	files, err := f.Readdir(0)
	if err != nil {
		return err
	}
	fs.pathLabel.SetText(path)

	// Sort files by name
	sort.Sort(listFileInfo(files))

	// Reads directory contents and loads into the list
	fs.list.Clear()
	// Adds previous directory
	prev := gui.NewImageLabel("..")
	prev.SetIcon(string(icon.FolderOpen))
	fs.list.Add(prev)
	// Adds directory files
	for i := 0; i < len(files); i++ {
		// Checks if is directory
		fname := files[i].Name()
		if files[i].IsDir() {
			item := gui.NewImageLabel(fname)
			item.SetIcon(string(icon.FolderOpen))
			fs.list.Add(item)
			continue
		}
		// Checks file filters
		if len(fs.fileFilters) > 0 {
			show := false
			for _, f := range fs.fileFilters {
				match, err := filepath.Match(f, fname)
				if err != nil {
					return err
				}
				if match {
					show = true
				}
			}
			if !show {
				continue
			}
		}
		// Adds file item to the list
		item := gui.NewImageLabel(fname)
		item.SetIcon(string(icon.InsertPhoto))
		fs.list.Add(item)
	}
	fs.path = path
	return nil
}

func (fs *FileSelect) Selected() string {

	selist := fs.list.Selected()
	if len(selist) == 0 {
		return ""
	}
	label := selist[0].(*gui.ImageLabel)
	text := label.Text()
	return filepath.Join(fs.path, text)
}

// SetFileFilters sets filters for file names which should be shown.
// Each filter must be a valid regular expression.
func (fs *FileSelect) SetFileFilters(filter ...string) {

	for _, f := range filter {
		fs.fileFilters = append(fs.fileFilters, f)
	}
}

func (fs *FileSelect) onSelect() {

	// Get selected image label and its txt
	sel := fs.list.Selected()[0]
	label := sel.(*gui.ImageLabel)
	text := label.Text()

	// Checks if previous directory
	if text == ".." {
		dir, _ := filepath.Split(fs.pathLabel.Text())
		fs.SetPath(filepath.Dir(dir))
		return
	}

	// Checks if it is a directory
	path := filepath.Join(fs.pathLabel.Text(), text)
	s, err := os.Stat(path)
	if err != nil {
		panic(err)
	}
	if s.IsDir() {
		fs.SetPath(path)
	}
}

// For sorting array of FileInfo by Name
type listFileInfo []os.FileInfo

func (fi listFileInfo) Len() int      { return len(fi) }
func (fi listFileInfo) Swap(i, j int) { fi[i], fi[j] = fi[j], fi[i] }
func (fi listFileInfo) Less(i, j int) bool {

	return fi[i].Name() < fi[j].Name()
}
