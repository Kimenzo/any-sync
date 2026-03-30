package commonspace

import (
	"context"
	"os"
	"path"
	"time"

	anystore "github.com/anyproto/any-store"

	"github.com/Kimenzo/any-sync/app"
	"github.com/Kimenzo/any-sync/commonspace/object/tree/objecttree"
	"github.com/Kimenzo/any-sync/commonspace/spacestorage"
)

type spaceStorageProvider struct {
	rootPath  string
	anyStores map[string]anystore.DB
}

func (s *spaceStorageProvider) Run(ctx context.Context) (err error) {
	return nil
}

func (s *spaceStorageProvider) Close(ctx context.Context) (err error) {
	if s.anyStores == nil {
		return s.removeRootPath()
	}
	for id, store := range s.anyStores {
		if closeErr := store.Close(); closeErr != nil && err == nil {
			err = closeErr
		}
		delete(s.anyStores, id)
	}
	if err != nil {
		return err
	}
	return s.removeRootPath()
}

func (s *spaceStorageProvider) Init(a *app.App) (err error) {
	return nil
}

func (s *spaceStorageProvider) Name() (name string) {
	return spacestorage.CName
}

func (s *spaceStorageProvider) WaitSpaceStorage(ctx context.Context, id string) (spacestorage.SpaceStorage, error) {
	if s.anyStores == nil {
		s.anyStores = make(map[string]anystore.DB)
	}
	if store, ok := s.anyStores[id]; ok {
		delete(s.anyStores, id)
		return spacestorage.New(ctx, id, store)
	}
	dbPath := path.Join(s.rootPath, id)
	if _, err := os.Stat(dbPath); err != nil {
		return nil, spacestorage.ErrSpaceStorageMissing
	}
	db, err := anystore.Open(ctx, dbPath, &anystore.Config{SQLiteConnectionOptions: map[string]string{"synchronous": "off"}})
	if err != nil {
		return nil, err
	}
	testStore := objecttree.TestStore{
		DB:   db,
		Path: dbPath,
	}
	return spacestorage.New(ctx, id, testStore)
}

func (s *spaceStorageProvider) SetStore(id string, store anystore.DB) {
	if s.anyStores == nil {
		s.anyStores = make(map[string]anystore.DB)
	}
	s.anyStores[id] = store
}

func (s *spaceStorageProvider) SpaceExists(id string) bool {
	if id == "" {
		return false
	}
	dbPath := path.Join(s.rootPath, id)
	if _, err := os.Stat(dbPath); err != nil {
		return false
	}
	return true
}

func (s *spaceStorageProvider) CreateSpaceStorage(ctx context.Context, payload spacestorage.SpaceStorageCreatePayload) (spacestorage.SpaceStorage, error) {
	id := payload.SpaceHeaderWithId.Id
	if s.SpaceExists(id) {
		return nil, spacestorage.ErrSpaceStorageExists
	}
	if s.anyStores == nil {
		s.anyStores = make(map[string]anystore.DB)
	}
	dbPath := path.Join(s.rootPath, id)
	db, err := anystore.Open(ctx, dbPath, nil)
	if err != nil {
		return nil, err
	}
	return spacestorage.Create(ctx, db, payload)
}

func (s *spaceStorageProvider) removeRootPath() error {
	var err error
	for i := 0; i < 100; i++ {
		err = os.RemoveAll(s.rootPath)
		if err == nil || os.IsNotExist(err) {
			return nil
		}
		time.Sleep(20 * time.Millisecond)
	}
	return nil
}
