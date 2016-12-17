package storage

import (
	//"encoding/binary"
	"github.com/boltdb/bolt"
	"log"
	"time"
)

//itob function
// func itob(v int) []byte {
// 	buf := make([]byte, 8)
// 	binary.BigEndian.PutUint64(buf, uint64(v))
// 	return buf
// }

type Store struct {
	Path string
	db   *bolt.DB
}

func (s *Store) Open(name string) error {
	// 1 - Opens the bolt db
	db, err := bolt.Open(s.Path, 0666, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Println(err)
		return err
	}
	s.db = db
	// 2 - Starts a writable transaction
	tx, err := db.Begin(true)
	if err != nil {
		log.Println(err)
		return err
	}
	defer tx.Rollback()
	// 3 - Initialized buckets to make sure they exist
	if _, err := tx.CreateBucketIfNotExists([]byte(name)); err != nil {
		log.Println(err)
		return err
	}
	// 4 - Commits the transaction
	return tx.Commit()
}

func (s *Store) Close() error {
	if err := s.db.Close(); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (s *Store) CreateApplicant(a *Applicant) error {
	// 1 - Start a writable transaction
	tx, err := s.db.Begin(true)
	if err != nil {
		log.Println(err)
		return err
	}
	defer tx.Rollback()
	// 2 - retrieve bucket
	bkt := tx.Bucket([]byte("Applicant"))

	// 3 - Get sequence for new applicant
	// seq, _ := bkt.NextSequence()
	// a.SeqID = int(seq)
	// if a.ID == nil {
	// 	a.ID = randId(20)
	// }
	// 4 - Marshall note to the bucket
	buf, err := a.MarshalBinary()
	if err != nil {
		log.Println(err)
		return err
	}
	// 5 = Save Applicant to bucket
	if err := bkt.Put([]byte(a.Email), buf); err != nil {
		log.Println(err)
		return err
	}
	return tx.Commit()
}

func (s *Store) GetApplicant(email string) (*Applicant, error) {
	// 1 - Start a readonly transaction
	tx, err := s.db.Begin(false)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer tx.Rollback()
	// 2 - Retrieve bucket and read encoded Applicant bytes
	buf := tx.Bucket([]byte("Applicant")).Get([]byte(email))
	if buf == nil {
		return nil, nil
	}
	// 3 - Unmarshall Applicant bytes into struct
	var a Applicant
	if err := a.UnmarshalBinary(buf); err != nil {
		log.Println(err)
		return nil, err
	}
	return &a, nil
}

func (s *Store) GetApplicants() ([]*Applicant, error) {
	// 1 - Start a readonly transaction
	tx, err := s.db.Begin(false)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer tx.Rollback()
	// 2 - Create a cursor on the db
	c := tx.Bucket([]byte("Applicant")).Cursor()

	var a []*Applicant

	// Iterate over the cursor and put contents
	for k, v := c.First(); k != nil; k, v = c.Next() {
		var aC Applicant
		if err := aC.UnmarshalBinary(v); err != nil {
			log.Println(err)
			return nil, err
		}
		a = append(a, &aC)
	}
	return a, nil
}
