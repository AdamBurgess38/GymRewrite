package gymstore

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type jsonStore struct {
}

func (store *jsonStore) SaveUser(ue User, path string) error {
	dat, _ := json.MarshalIndent(ue, "", " ")

	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			err := os.MkdirAll(path, os.ModePerm)
			if err != nil {
				return err
			}
		}
	}
	err := os.WriteFile(path+ue.Name+".json", dat, 0644)
	if err != nil {
		fmt.Printf("Unsuccesful save of: %s, in directory %s\n", ue.Name, path)
		return err
	}

	return nil
}

func (store *jsonStore) LoadUser(username string, path string) (User, error) {
	var ue = &User{Name: "",
		Exercises: make(map[string]Exercise)}

	var data []byte

	fmt.Println(username + ".json")
	dat, err := os.ReadFile(path + username + ".json")
	if err != nil {
		fmt.Printf("Not a valid username, please try again\n")
		return User{}, err
	}
	data = dat

	json.Unmarshal([]byte(data), &ue)

	backupfile(*ue, path)
	return *ue, nil
}

func backupfile(ue User, path string) error {
	getDate := func() string {
		currentTime := time.Now()
		day := currentTime.Local().Day()
		month := currentTime.Local().Month()
		year := currentTime.Local().Year()
		return strconv.Itoa(day) + "/" + strconv.Itoa(int(month)) + "/" + strconv.Itoa(year)
	}

	dat, _ := os.ReadFile(ue.Name + ".json")
	dateFileFormat := strings.Replace(getDate(), "/", ":", 2)
	path = path + "/backups/"
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			err := os.MkdirAll(path, os.ModePerm)
			if err != nil {
				return err
			}
		}
	}
	err := os.WriteFile(path+ue.Name+"-BACKUP:"+dateFileFormat+".json", dat, 0644)
	if err != nil {
		fmt.Printf("Unsuccesful backup of: %s", ue.Name)
		return err
	}

	return nil
}

func savefile(ue User) error {
	getDate := func() string {
		currentTime := time.Now()
		day := currentTime.Local().Day()
		month := currentTime.Local().Month()
		year := currentTime.Local().Year()
		return strconv.Itoa(day) + "/" + strconv.Itoa(int(month)) + "/" + strconv.Itoa(year)
	}

	dat, _ := os.ReadFile(ue.Name + ".json")
	dateFileFormat := strings.Replace(getDate(), "/", ":", 2)
	path := "./backups"
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			err := os.MkdirAll(path, os.ModePerm)
			if err != nil {
				return err
			}
		}
	}
	err := os.WriteFile("./backups/"+ue.Name+"-BACKUP:"+dateFileFormat+".json", dat, 0644)
	if err != nil {
		fmt.Printf("Unsuccesful backup of: %s", ue.Name)
		return err
	}

	return nil
}
