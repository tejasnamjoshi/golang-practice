package db

import (
	"encoding/binary"
	"time"

	"github.com/boltdb/bolt"
)

var taskBucket = []byte("tasks")
var db *bolt.DB

type Task struct {
	Key   int
	Value string
}

func Init(path string) error {
	// db, err := db.Open(path, 0600, &bolt.Options{Timeout: 1 * time.Second}) // this is not correct. This means db is a local var in this function

	var err error
	db, err = bolt.Open(path, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}

	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskBucket)
		if err != nil {
			return err
		}

		return nil
	})
}

func CreateTask(task string) (int, error) {
	var id int
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		id64, _ := b.NextSequence()
		key := itob(int(id64))
		value := []byte(task)
		err := b.Put(key, value)
		if err != nil {
			return err
		}

		id = int(id64)
		return nil
	})

	if err != nil {
		return -1, err
	}

	return id, nil
}

func GetTasks() ([]Task, error) {
	tasks := []Task{}
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			tasks = append(tasks, Task{Key: btoi(k), Value: string(v)})
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func RemoveTask(key int) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		err := b.Delete(itob(key))

		return err
	})
}

func itob(id int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(id))
	return b
}

func btoi(id []byte) int {
	return int(binary.BigEndian.Uint64(id))
}
