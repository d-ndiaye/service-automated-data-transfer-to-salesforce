package internal

import (
	"github.com/fsnotify/fsnotify"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"service-automatisierte-daten-in-salesforce/pkg/config"
	"sort"
	"strings"
)

type Watcher struct {
	fileManager FileManager
	repository  Repository
}

func NewWatcher(fileManager FileManager, repository Repository) Watcher {
	return Watcher{
		fileManager: fileManager,
		repository:  repository,
	}
}

// WatchFolder This method watches a given folder to see if there is a new .pem file.
// return LicenceInfo{}, error
func (w Watcher) WatchFolder(conf config.Config) (LicenceInfo, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return LicenceInfo{}, err
	}
	defer watcher.Close()
	err = watcher.Add(conf.FolderPath)
	if err != nil {
		return LicenceInfo{}, err
	}
	log.Println("Watching the folder :", conf.FolderPath)
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return LicenceInfo{}, err
			}
			if event.Op&fsnotify.Create == fsnotify.Create {
				if strings.HasSuffix(event.Name, ".pem") {
					log.Println("New .pem file found :", event.Name)
					licenceID := w.fileManager.ExtractFileName(event.Name)
					licenceInfo, LicenceProperties, _ := w.fileManager.ReadInfosPemFile(event.Name, licenceID)
					w.repository.InsertLicenceInfo(licenceInfo)
					w.repository.InsertLicenceProperties(LicenceProperties)
					w.moveFile(event.Name, conf.BackupFolder+"\\"+w.fileManager.ExtractFileName(event.Name)+".pem")
				}
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return LicenceInfo{}, err
			}
			log.Println("Error :", err)
		}
	}
}

// moveFile this method moves a pem file to another folder "backupFolder".
func (w Watcher) moveFile(folderPath string, backupFolder string) {
	dirPath := filepath.Dir(backupFolder)
	er := os.Mkdir(dirPath, os.ModePerm)
	if er != nil {
		err := os.Rename(folderPath, backupFolder)
		if err != nil {
			log.Println("Error moveFile")
		}
	}
	w.sortFolderByDate(dirPath)
}

// sortFolderByDate this method sorts the backup folder by date
func (w Watcher) sortFolderByDate(backupFolder string) {
	files, _ := ioutil.ReadDir(backupFolder)
	sort.Slice(files, func(i, j int) bool {
		return files[i].ModTime().Before(files[j].ModTime())
	})
}
