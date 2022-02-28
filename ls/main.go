package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	pathlib "path/filepath"
	"regexp"
	"strings"

	"golang.org/x/term"
)

func round(x float64) int {
	return int(math.Floor(x + 0/5))
}
func getdircontent(dir string) (content []string) {
	content = make([]string, 0)
	files, _ := ioutil.ReadDir(dir)

	for _, f := range files {
		content = append(content, f.Name())
	}
	return content
}

const (
	red   = "\033[31m"
	grn   = "\033[32m"
	white = "\033[37m"
)

func insert(a []string, c string, i int) []string {
	return append(a[:i], append([]string{c}, a[i:]...)...)
}

func main() {
	// Get the current directory

	inpdir := os.Args

	path, err := os.Getwd()
	if len(inpdir) > 1 {
		usrdir := inpdir[1]
		// Check if the specified directory is valid when appeneded to the current working directory
		if _, err := os.Stat(path + usrdir); !os.IsNotExist(err) {
			path = path + usrdir
		} else {
			if _, err := os.Stat(usrdir); !os.IsNotExist(err) {
				path = usrdir
			} else {
				fmt.Println("Invalid directory")
				os.Exit(1)
			}

		}
	}
	// remove any dot slashes from the path
	path = strings.Replace(path, "./", "", -1)

	print("Contents of " + path + ":\n")

	if err != nil {
		print("Error reading dir: ")
		log.Fatal(err)
	}
	files := getdircontent(path)
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		log.Fatal(err)
	}
	/*Using golang and the term, os, fmt and ioutil packages, create a script that displays all the files and folders from the terminal's working directory in 4 collumns, using red for folders and green for files. make sure the collumns always stretch across the entire terminal's width.*/

	spacing := math.Floor(float64(width / 4))
	var printstr string
	for i, file := range files {
		rawfile := path + "/" + file
		icon := iconer(rawfile)
		namelen := len(file)
		if icon != "" {
			namelen += 2
			file = icon + "\u2002" + file
		}

		var displayname string
		if namelen < int(spacing) {
			displayname = file + strings.Repeat(" ", round(spacing-float64(namelen)))
		} else if namelen < int(spacing)*2 {
			displayname = file + strings.Repeat(" ", round(spacing*2-float64(namelen)))
			files = insert(files, strings.Repeat(" ", int(spacing)), i+1)
		} else if namelen < int(spacing)*3 {
			displayname = file + strings.Repeat(" ", round(spacing*3-float64(namelen)))
			files = insert(files, strings.Repeat(" ", int(spacing)), i+1)
			files = insert(files, strings.Repeat(" ", int(spacing)), i+2)
		} else if namelen < int(spacing)*4 {
			displayname = file + strings.Repeat(" ", round(spacing*4-float64(namelen)))
			files = insert(files, strings.Repeat(" ", int(spacing)), i+1)
			files = insert(files, strings.Repeat(" ", int(spacing)), i+2)
			files = insert(files, strings.Repeat(" ", int(spacing)), i+3)
		} else {
			// if it is more than 4 collumns, it will be cut off
			displayname = file[:(int(spacing)*4)-3] + "..."
			files = insert(files, strings.Repeat(" ", int(spacing)), i+1)
			files = insert(files, strings.Repeat(" ", int(spacing)), i+2)
			files = insert(files, strings.Repeat(" ", int(spacing)), i+3)
		}

		if m, err := regexp.MatchString("\\s+", file); m == false {
			if err != nil {
				log.Fatal(err)
			}
			fst, err := os.Stat(rawfile)
			if err != nil {
				log.Fatal(err)
			}
			if fst.IsDir() {

				printstr += red + displayname
			} else {
				printstr += grn + displayname
			}
		}
	}
	printstr += white
	fmt.Println(printstr)
}

func iconer(dirpath string) string {
	// get the file extension

	var icon string
	// Get the name of the file/folder without the path
	path := pathlib.Base(dirpath)
	f, err := os.Stat(dirpath)
	// if the error is an ENOENT, skip the file
	if err != nil {
		log.Fatal(err)
	}
	if f.IsDir() == false {

		ext := path[strings.LastIndex(path, ".")+1:]
		switch ext {
		case "py":
			icon = "\ue73c"
		case "go", "mod", "sum":
			icon = "\ufcd1"
		case "exe":
			icon = "\uf085"
		case "jar", "java", "class":
			icon = "\ue256"
		case "cpp":
			icon = "\ue61d"
		case "c":
			icon = "\ue61e"
		case "cs":
			icon = "\uf81a"
		case "h":
			icon = "\uf471"
		case "sh", "bat", "cmd":
			icon = "\uf120"
		case "sql", "db", "sqlite3":
			icon = "\uf472"
		case "js":
			icon = "\ue74e"
		case "ts":
			icon = "\ufbe4"
		case "xml", "json", "yml", "yaml", "toml":
			icon = "\ufb25"
		case "txt":
			icon = "\ue612"
		case "html", "htm":
			icon = "\ue736"
		case "css":
			icon = "\ue749"
		case "md":
			icon = "\ue73e"
		case "log":
			icon = "\uf71d"
		case "lua":
			icon = "\ue620"
		case "php":
			icon = "\uf81e"
		case "img", "png", "jpg", "jpeg", "gif", "bmp", "ico":
			icon = "\uf03e"
		case "woff", "woff2", "ttf", "otf", "eot":
			icon = "\ufb68"
		case "sass", "scss":
			icon = "\ufcea"
		case "mp3", "wav", "ogg", "flac", "aac", "m4a", "mp4", "m4v", "avi", "mpg", "mpeg", "mov", "webm", "wmv", "flv", "3gp", "3g2":
			icon = "\uf1c8"
		case "zip", "rar", "7z", "tar", "gz", "bz2", "z", "lz", "lzma", "lzo", "lz4", "lzop", "lzma2", "lz4hc", "xz", "zst", "zstd", "zpaq":
			icon = "\uf1c6"
		case "csv", "xls", "xlt", "xltx", "xltm", "xltb", "xlsx", "xlsm", "xlsb", "xla", "xlam", "xll", "xlw", "xlr":
			icon = "\uf0ce"
		case "dll", "so":
			icon = "\uf471"
		default:
			icon = ""
		}
	} else {
		// Get the directory name
		dir := path[strings.LastIndex(path, "/")+1:]
		switch dir {
		case ".git":
			icon = "\ue5fd"
		case ".vscode":
			icon = "\ue70c"
		case "logs":
			icon = "\uf71d"
		case "backups", "backup":
			icon = "\uf650"
		case "env", "envs", "environment", ".env":
			icon = "\ue73c"
		case "json":
			icon = "\ufb25"
		case "config", "configs":
			icon = "\uf085"
		default:
			icon = "\uf07c"
		}
	}
	return icon

}
