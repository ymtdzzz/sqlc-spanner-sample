package main

import (
	"context"
	"database/sql"
	"fmt"
	"sync"

	"github.com/ymtdzzz/sqlc-spanner-sample/db"

	_ "github.com/googleapis/go-sql-spanner"
	spannerdriver "github.com/googleapis/go-sql-spanner"
)

const (
	project  = "zerocloc-dev-test-env"
	instance = "sqlc-test"
	database = "sqlc-test"

	counterID = "05a7f30c-823c-4502-a866-6ac783050e4f"
)

var (
	userIDs = []string{
		"199f8059-558a-4c6f-aad3-526859cfa88e",
		"7d586bac-1c9e-4c3d-aefe-53c8649352f0",
		"a94c740a-3791-492d-88be-b05cf0b85252",
	}
)

func main() {
	sqldb, err := sql.Open("spanner", fmt.Sprintf("projects/%s/instances/%s/databases/%s", project, instance, database))
	if err != nil {
		panic(err)
	}

	queries := db.New(sqldb)

	ctx := context.Background()

	fmt.Println("------- query: Users -------")
	for _, id := range userIDs {
		user, err := queries.GetUser(ctx, id)
		if err != nil {
			panic(err)
		}

		fmt.Printf("%#v\n", user)
	}

	fmt.Println("------- query: UsersWithAddresses -------")
	for _, id := range userIDs {
		useraddrRows, err := queries.GetUserWithAddresses(ctx, id)
		if err != nil {
			panic(err)
		}

		for _, useraddr := range useraddrRows {
			fmt.Printf("%#v\n", useraddr)
		}
	}

	waitNum := 10
	var wg sync.WaitGroup
	for i := 0; i < waitNum; i++ {
		fmt.Printf("------- query: UpdateCounter (%d) -------\n", i)
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				err := incrementCounter(ctx, queries, sqldb)
				if err == spannerdriver.ErrAbortedDueToConcurrentModification {
					continue
				}
				if err != nil {
					panic(err)
				}
				break
			}
		}()
	}

	wg.Wait()

	c, err := queries.GetCounter(ctx, counterID)
	if err != nil {
		panic(err)
	}
	fmt.Printf("counter: %d\n", c.Count)
}

func incrementCounter(ctx context.Context, queries *db.Queries, sqldb *sql.DB) error {
	// Read-Write Transaction
	tx, err := sqldb.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback()

	qtx := queries.WithTx(tx)
	c, err := qtx.GetCounter(ctx, counterID)
	if err != nil {
		return err
	}

	if err := qtx.UpdateCounter(ctx, db.UpdateCounterParams{
		ID:    c.ID,
		Count: c.Count + 1,
	}); err != nil {
		return err
	}

	return tx.Commit()
}
