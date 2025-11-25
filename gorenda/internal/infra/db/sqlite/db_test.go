package sqlite

import (
	"testing"
)

func TestDatabaseConnection(t *testing.T) {
	db, err := Open(":memory:")
	if err != nil {
		t.Fatal("Не удалось открыть базу:", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		t.Fatal("База не отвечает:", err)
	}

	t.Log("работает")
}
