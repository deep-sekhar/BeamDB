package storage

import (
	"fmt"
	"os"
	"time"
)

// save data implementation
func SaveData(path string, data []byte) error {
	// create a temporary file to save the data to support multiple reads
	// this makes it atomic wrt reads
	tmp := fmt.Sprintf("%s.tmp.%d", path, time.Now().UnixNano())

	// open the file for writing
	// O_WRONLY: write only
	// O_CREATE: create the file if it does not exist
	// O_EXCL: ensure that the file does not exist
	// 0664: file permissions - rw-rw-r--
	fp, err := os.OpenFile(tmp, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0664)

	// could not open the file
	if err != nil {
		return err
	}

	// delete the temporary file if anything goes wrong
	defer func() {
		if err != nil {
			// delete the temporary file if it exists
			os.Remove(tmp)
		}
	}()

	//  write the data to the file
	_, err = fp.Write(data)

	// could not write the data
	if err != nil {
		return err
	}

	// sync to ensure durability
	err = fp.Sync()

	// could not write to disk
	if err != nil {
		return err
	}

	// close the file before renaming
	err = fp.Close()
	if err != nil {
		return err
	}

	// rename the temporary file to the actual file name
	return os.Rename(tmp, path)
}
