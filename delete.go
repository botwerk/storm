package storm

import "github.com/boltdb/bolt"

// Delete deletes a key from a bucket
func (s *DB) Delete(bucketName string, key interface{}) error {
	id, err := toBytes(key)
	if err != nil {
		return err
	}

	return s.Bolt.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		if bucket == nil {
			return ErrNotFound
		}

		return bucket.Delete(id)
	})
}
