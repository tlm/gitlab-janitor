package backup

import (
	"context"
	"testing"
	"time"

	"gocloud.dev/blob"
	_ "gocloud.dev/blob/memblob"
)

var (
	DummyData = []byte("dummy data")
)

func comparePruneLists(l1 []string, l2 []string) bool {
	if len(l1) != len(l2) {
		return false
	}
	lookup := make(map[string]bool, len(l1))
	for _, l := range l1 {
		lookup[l] = false
	}
	for _, l := range l2 {
		val, ok := lookup[l]
		if !ok || val {
			return false
		}
		lookup[l] = true
	}
	return true
}

func createDummyFiles(bucket *blob.Bucket, files []string) error {
	ctx := context.Background()
	for _, file := range files {
		if err := bucket.WriteAll(ctx, file, DummyData, nil); err != nil {
			return err
		}
	}
	return nil
}

func TestWithKeepNumberOfVersions(t *testing.T) {
	tests := []struct {
		Decision   Decider
		PreFiles   []string
		PruneFiles []string
	}{
		{
			Decision: WithKeepNumberOfVersions(3),
			PreFiles: []string{
				"1540174211_2018_10_22_11.3.6-ee_gitlab_backup.tar",
				"1540174453_2018_10_22_11.3.6-ee_gitlab_backup.tar",
				"1543197673_2018_11_26_11.4.0-ee_gitlab_backup.tar",
				"1543284005_2018_11_27_11.4.0-ee_gitlab_backup.tar",
				"1543370414_2018_11_28_11.4.0-ee_gitlab_backup.tar",
				"1561696174_2019_06_28_11.7.0-ee_gitlab_backup.tar",
				"1561776781_2019_06_29_11.7.0-ee_gitlab_backup.tar",
				"1561863334_2019_06_30_11.7.0-ee_gitlab_backup.tar",
				"1564884018_2019_08_04_12.0.3-ee_gitlab_backup.tar",
				"1564970415_2019_08_05_12.0.3-ee_gitlab_backup.tar",
				"1565056820_2019_08_06_12.0.3-ee_gitlab_backup.tar",
			},
			PruneFiles: []string{
				"1540174211_2018_10_22_11.3.6-ee_gitlab_backup.tar",
				"1540174453_2018_10_22_11.3.6-ee_gitlab_backup.tar",
			}},
	}

	for i, test := range tests {
		bucket, err := blob.OpenBucket(context.Background(), "mem://")
		if err != nil {
			t.Fatalf("unexpected error opening memory bucket for test: %v", err)
		}
		defer bucket.Close()

		if err := createDummyFiles(bucket, test.PreFiles); err != nil {
			t.Fatalf("unexpected error seeding bucket with files: %v", err)
		}

		res, err := CreatePruneList(bucket, test.Decision)
		if err != nil {
			t.Fatalf("unexpected error creating prune list: %v", err)
		}
		if !comparePruneLists(res, test.PruneFiles) {
			t.Errorf("test %d prune list does not match expected post list", i)
		}
	}
}

func TestWithKeepTime(t *testing.T) {
	tests := []struct {
		Decision   Decider
		PreFiles   []string
		PruneFiles []string
	}{
		{
			Decision: WithKeepAfterTime(time.Date(2019, 06, 29, 0, 0, 0, 0, time.UTC)),
			PreFiles: []string{
				"1540174211_2018_10_22_11.3.6-ee_gitlab_backup.tar",
				"1540174453_2018_10_22_11.3.6-ee_gitlab_backup.tar",
				"1543197673_2018_11_26_11.4.0-ee_gitlab_backup.tar",
				"1543284005_2018_11_27_11.4.0-ee_gitlab_backup.tar",
				"1543370414_2018_11_28_11.4.0-ee_gitlab_backup.tar",
				"1561696174_2019_06_28_11.7.0-ee_gitlab_backup.tar",
				"1561776781_2019_06_29_11.7.0-ee_gitlab_backup.tar",
				"1561863334_2019_06_30_11.7.0-ee_gitlab_backup.tar",
				"1564884018_2019_08_04_12.0.3-ee_gitlab_backup.tar",
				"1564970415_2019_08_05_12.0.3-ee_gitlab_backup.tar",
				"1565056820_2019_08_06_12.0.3-ee_gitlab_backup.tar",
			},
			PruneFiles: []string{
				"1540174211_2018_10_22_11.3.6-ee_gitlab_backup.tar",
				"1540174453_2018_10_22_11.3.6-ee_gitlab_backup.tar",
				"1543197673_2018_11_26_11.4.0-ee_gitlab_backup.tar",
				"1543284005_2018_11_27_11.4.0-ee_gitlab_backup.tar",
				"1543370414_2018_11_28_11.4.0-ee_gitlab_backup.tar",
				"1561696174_2019_06_28_11.7.0-ee_gitlab_backup.tar",
			}},
	}

	for i, test := range tests {
		bucket, err := blob.OpenBucket(context.Background(), "mem://")
		if err != nil {
			t.Fatalf("unexpected error opening memory bucket for test: %v", err)
		}
		defer bucket.Close()

		if err := createDummyFiles(bucket, test.PreFiles); err != nil {
			t.Fatalf("unexpected error seeding bucket with files: %v", err)
		}

		res, err := CreatePruneList(bucket, test.Decision)
		if err != nil {
			t.Fatalf("unexpected error creating prune list: %v", err)
		}
		if !comparePruneLists(res, test.PruneFiles) {
			t.Errorf("test %d prune list does not match expected post list", i)
		}
	}
}

func TestWithAggregateAgree(t *testing.T) {
	tests := []struct {
		Decision   Decider
		PreFiles   []string
		PruneFiles []string
	}{
		{
			Decision: WithAggregateAgree(),
			PreFiles: []string{
				"1540174211_2018_10_22_11.3.6-ee_gitlab_backup.tar",
				"1540174453_2018_10_22_11.3.6-ee_gitlab_backup.tar",
				"1543197673_2018_11_26_11.4.0-ee_gitlab_backup.tar",
				"1543284005_2018_11_27_11.4.0-ee_gitlab_backup.tar",
				"1543370414_2018_11_28_11.4.0-ee_gitlab_backup.tar",
				"1561696174_2019_06_28_11.7.0-ee_gitlab_backup.tar",
				"1561776781_2019_06_29_11.7.0-ee_gitlab_backup.tar",
				"1561863334_2019_06_30_11.7.0-ee_gitlab_backup.tar",
				"1564884018_2019_08_04_12.0.3-ee_gitlab_backup.tar",
				"1564970415_2019_08_05_12.0.3-ee_gitlab_backup.tar",
				"1565056820_2019_08_06_12.0.3-ee_gitlab_backup.tar",
			},
			PruneFiles: []string{}},
		{
			Decision: WithAggregateAgree(
				WithKeepAfterTime(time.Date(2019, 06, 29, 0, 0, 0, 0, time.UTC)),
				WithKeepNumberOfVersions(1),
			),
			PreFiles: []string{
				"1540174211_2018_10_22_11.3.6-ee_gitlab_backup.tar",
				"1540174453_2018_10_22_11.3.6-ee_gitlab_backup.tar",
				"1543197673_2018_11_26_11.4.0-ee_gitlab_backup.tar",
				"1543284005_2018_11_27_11.4.0-ee_gitlab_backup.tar",
				"1543370414_2018_11_28_11.4.0-ee_gitlab_backup.tar",
				"1561696174_2019_06_28_11.7.0-ee_gitlab_backup.tar",
				"1561776781_2019_06_29_11.7.0-ee_gitlab_backup.tar",
				"1561863334_2019_06_30_11.7.0-ee_gitlab_backup.tar",
				"1564884018_2019_08_04_12.0.3-ee_gitlab_backup.tar",
				"1564970415_2019_08_05_12.0.3-ee_gitlab_backup.tar",
				"1565056820_2019_08_06_12.0.3-ee_gitlab_backup.tar",
			},
			PruneFiles: []string{
				"1540174211_2018_10_22_11.3.6-ee_gitlab_backup.tar",
				"1540174453_2018_10_22_11.3.6-ee_gitlab_backup.tar",
				"1543197673_2018_11_26_11.4.0-ee_gitlab_backup.tar",
				"1543284005_2018_11_27_11.4.0-ee_gitlab_backup.tar",
				"1543370414_2018_11_28_11.4.0-ee_gitlab_backup.tar",
				"1561696174_2019_06_28_11.7.0-ee_gitlab_backup.tar",
			}},
	}

	for i, test := range tests {
		bucket, err := blob.OpenBucket(context.Background(), "mem://")
		if err != nil {
			t.Fatalf("unexpected error opening memory bucket for test: %v", err)
		}
		defer bucket.Close()

		if err := createDummyFiles(bucket, test.PreFiles); err != nil {
			t.Fatalf("unexpected error seeding bucket with files: %v", err)
		}

		res, err := CreatePruneList(bucket, test.Decision)
		if err != nil {
			t.Fatalf("unexpected error creating prune list: %v", err)
		}
		if !comparePruneLists(res, test.PruneFiles) {
			t.Errorf("test %d prune list does not match expected post list", i)
		}
	}
}
