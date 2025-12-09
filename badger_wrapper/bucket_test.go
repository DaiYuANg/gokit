package badger_wrapper

import (
	"testing"

	"github.com/dgraph-io/badger/v4"
	"github.com/stretchr/testify/require"
)

type TestValue struct {
	Name  string
	Age   int
	Valid bool
}

func newTestDB(t *testing.T) *badger.DB {
	opts := badger.DefaultOptions("").WithInMemory(true)
	db, err := badger.Open(opts)
	require.NoError(t, err)
	t.Cleanup(func() {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	})
	return db
}

func TestBucketBasicOperations(t *testing.T) {
	db := newTestDB(t)
	bkt := NewBucket[TestValue](db, "test")

	// Set & Get
	val := TestValue{Name: "Alice", Age: 30, Valid: true}
	err := bkt.Set("alice", val)
	require.NoError(t, err)

	got, err := bkt.Get("alice")
	require.NoError(t, err)
	require.Equal(t, val, got)

	// Exists
	exists, err := bkt.Exists("alice")
	require.NoError(t, err)
	require.True(t, exists)

	exists, err = bkt.Exists("bob")
	require.NoError(t, err)
	require.False(t, exists)

	// Count
	count, err := bkt.Count()
	require.NoError(t, err)
	require.Equal(t, 1, count)

	// ListKeys
	keys, err := bkt.ListKeys()
	require.NoError(t, err)
	require.Equal(t, []string{"alice"}, keys)

	// ForEach
	err = bkt.ForEach(func(key string, val TestValue) error {
		require.Equal(t, "alice", key)
		require.Equal(t, TestValue{Name: "Alice", Age: 30, Valid: true}, val)
		return nil
	})
	require.NoError(t, err)

	// BatchSet
	items := map[string]TestValue{
		"bob":   {Name: "Bob", Age: 25, Valid: true},
		"carol": {Name: "Carol", Age: 28, Valid: false},
	}
	err = bkt.BatchSet(items)
	require.NoError(t, err)

	count, err = bkt.Count()
	require.NoError(t, err)
	require.Equal(t, 3, count)

	// PrefixScan
	err = bkt.PrefixScan("b", func(key string, val TestValue) (stop bool, err error) {
		require.Equal(t, "bob", key)
		require.Equal(t, "Bob", val.Name)
		return false, nil
	})
	require.NoError(t, err)

	// Delete
	err = bkt.Delete("alice")
	require.NoError(t, err)

	_, err = bkt.Get("alice")
	require.Error(t, err)
}
