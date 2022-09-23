package helper

import (
	"context"
	"fmt"

	"github.com/sky0621/cv-admin/src/ent"
)

func WithTransaction(ctx context.Context, cli *ent.Client, fn func(tx *ent.Tx) error) error {
	tx, err := cli.Tx(ctx)
	if err != nil {
		return err
	}

	fErr := fn(tx)
	if fErr != nil {
		if err := tx.Rollback(); err != nil {
			return fmt.Errorf("%w: %v", fErr, err)
		}
		return fErr
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("%w: %v", fErr, err)
	}

	return nil
}
