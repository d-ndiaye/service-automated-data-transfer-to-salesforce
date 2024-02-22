# Automatisierte Ordner-Überwachung, Auswertung und Export von Lizenzdateien in ein CRM-System

## Dependencies

* Go 1.20
* Mysql v1.7.0

## How to use

In the [configuration](./config/service.yaml) file give:
* the path of the folder to be watched 'folderPath: C:\Dame\...\'
* the path of the backup folder 'backupFolder: C:\Dame\...\'
* and the Mysql DB connection configurations(username, password, port, host, dbname)
