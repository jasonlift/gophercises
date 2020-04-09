package db

import (
	"encoding/binary"
	"time"

	bolt "go.etcd.io/bbolt"
)

var dbProxy *BoltDb

type BoltDb struct {
	Db         *bolt.DB
	TaskBucket []byte
}

func Init(dbPath string) error {
	boltdb, err := bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	} else {
		dbProxy = &BoltDb{
			Db:         boltdb,
			TaskBucket: []byte("tasks"),
		}
		return dbProxy.Db.Update(func(tx *bolt.Tx) error {
			_, err := tx.CreateBucketIfNotExists(dbProxy.TaskBucket)
			return err
		})
	}
}

func CreateTask(task string) (int, error) {
	var id int
	err := dbProxy.Db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(dbProxy.TaskBucket)
		id64, _ := b.NextSequence()
		id = int(id64)
		key := itob(id)
		return b.Put(key, []byte(task))
	})
	if err != nil {
		return -1, err
	}
	return id, nil
}

func ReadAllTasks() ([]Task, error) {
	var tasks []Task
	err := dbProxy.Db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(dbProxy.TaskBucket)
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			tasks = append(tasks, Task{
				Key:   btoi(k),
				Value: string(v),
			})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func DeleteTask(id int) error {
	return dbProxy.Db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(dbProxy.TaskBucket)
		return b.Delete(itob(id))
	})
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}
