package day7

import (
	"aoc2022/pkg/inputreader"
	"strconv"
	"strings"
)

type File struct {
	Name string
	Size int
}

func (d Directory) getSubDirs() []*Directory {
	var subDirs []*Directory
	for _, subDir := range d.Subdirectories {
		subDirs = append(subDirs, subDir)
	}

	return subDirs
}

type Directory struct {
	Parent         *Directory
	Path           []string
	Files          []File
	Subdirectories map[string]*Directory
	Size           int
}

type FileSystem struct {
	Root             *Directory
	CurrentDirectory *Directory
}

func (fs *FileSystem) createSubdirectory(name string) error {
	fs.CurrentDirectory.Subdirectories[name] = &Directory{
		Parent:         fs.CurrentDirectory,
		Path:           append(fs.CurrentDirectory.Path, name),
		Subdirectories: make(map[string]*Directory),
	}

	return nil
}

func (fs *FileSystem) changeDirectory(dir string) error {
	switch dir {
	case "/":
		fs.CurrentDirectory = fs.Root
	case "..":
		fs.CurrentDirectory = fs.CurrentDirectory.Parent
	default:
		if _, ok := fs.CurrentDirectory.Subdirectories[dir]; !ok {
			fs.createSubdirectory(dir)
		}
		fs.CurrentDirectory = fs.CurrentDirectory.Subdirectories[dir]
	}

	return nil
}

func (fs *FileSystem) addItemToDirectory(fileInfo []string) error {
	if fileInfo[0] == "dir" {
		subDir := fileInfo[1]
		if _, ok := fs.CurrentDirectory.Subdirectories[subDir]; !ok {
			fs.createSubdirectory(subDir)
		}

		return nil
	}

	fileName := fileInfo[1]
	size, err := strconv.Atoi(fileInfo[0])
	if err != nil {
		return err
	}
	newFile := File{
		Name: fileName,
		Size: size,
	}
	fs.CurrentDirectory.Files = append(fs.CurrentDirectory.Files, newFile)
	fs.CurrentDirectory.Size += size

	return nil
}

func (fs *FileSystem) handleCommand(words []string) error {
	switch words[0] {
	case "$":
		if words[1] == "cd" {
			fs.changeDirectory(words[2])
		}
	default:
		fs.addItemToDirectory(words)
	}
	return nil
}

func (fs *FileSystem) findDirectoriesUnderMaxSize(maxSize int) ([]*Directory, error) {
	var dirsUnderMaxSize []*Directory

	currDir := fs.Root
	dirsToCheck := currDir.getSubDirs()
	for {
		if len(dirsToCheck) == 0 {
			return dirsUnderMaxSize, nil
		}
		currDir = dirsToCheck[0]
		dirsToCheck = dirsToCheck[1:]
		if currDir.Size <= maxSize {
			dirsUnderMaxSize = append(dirsUnderMaxSize, currDir)
		}

		dirsToCheck = append(dirsToCheck, currDir.getSubDirs()...)
	}
}

func (fs *FileSystem) findDirectoriesOverMinSize(minSize int) ([]*Directory, error) {
	var dirsOverMinSize []*Directory

	currDir := fs.Root
	dirsToCheck := currDir.getSubDirs()
	for {
		if len(dirsToCheck) == 0 {
			return dirsOverMinSize, nil
		}
		currDir = dirsToCheck[0]
		dirsToCheck = dirsToCheck[1:]
		if currDir.Size >= minSize {
			dirsOverMinSize = append(dirsOverMinSize, currDir)
			dirsToCheck = append(dirsToCheck, currDir.getSubDirs()...)
		}
	}
}

// setSizes will recursively increment the directory size with the sizes of the subdirectories from the leaves to the root
func setSizes(dir *Directory) error {
	if len(dir.Subdirectories) > 0 {
		for _, subdir := range dir.Subdirectories {
			err := setSizes(subdir)
			if err != nil {
				return err
			}

			dir.Size += subdir.Size
		}
	}

	return nil
}

func getTotalSize(dirs []*Directory) int {
	totalSize := 0
	for _, dir := range dirs {
		totalSize += dir.Size
	}

	return totalSize
}

func buildFileSystem(lines []string) (*FileSystem, error) {
	rootDir := &Directory{
		Subdirectories: make(map[string]*Directory),
	}
	fs := FileSystem{
		Root:             rootDir,
		CurrentDirectory: rootDir,
	}
	for _, line := range lines {
		words := strings.Split(line, " ")
		fs.handleCommand(words)
	}

	err := setSizes(fs.Root)
	if err != nil {
		return nil, err
	}

	return &fs, nil
}

func Part1() (int, error) {
	lines, err := inputreader.ReadLines("inputs/day7/1.txt")
	if err != nil {
		return 0, err
	}

	fs, err := buildFileSystem(lines)
	if err != nil {
		return 0, err
	}

	dirsUnderMaxSize, err := fs.findDirectoriesUnderMaxSize(100000)
	if err != nil {
		return 0, err
	}

	totalSize := getTotalSize(dirsUnderMaxSize)

	return totalSize, nil
}

func Part2() (int, error) {
	lines, err := inputreader.ReadLines("inputs/day7/1.txt")
	if err != nil {
		return 0, err
	}

	fs, err := buildFileSystem(lines)
	if err != nil {
		return 0, err
	}

	spaceToReclaim := 30000000 - (70000000 - fs.Root.Size)
	dirsOverMinSize, err := fs.findDirectoriesOverMinSize(spaceToReclaim)
	if err != nil {
		return 0, err
	}

	minSize := 30000000
	for _, dir := range dirsOverMinSize {
		if dir.Size < minSize {
			minSize = dir.Size
		}
	}

	return minSize, nil
}
