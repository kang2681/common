package sessions

import (
	"github.com/boj/redistore"
	"github.com/kang2681/sessions"
)

type FileStoreOptionFunc func(*sessions.FilesystemStore) error

func NewFileStore(path string, keyPairs []string, opts ...FileStoreOptionFunc) (*sessions.FilesystemStore, error) {
	keyPairsArr := make([][]byte, 0, len(keyPairs))
	for _, v := range keyPairs {
		keyPairsArr = append(keyPairsArr, []byte(v))
	}
	fs := sessions.NewFilesystemStore(path, keyPairsArr...)
	for _, opt := range opts {
		if err := opt(fs); err != nil {
			return nil, err
		}
	}
	return fs, nil
}

func SetFileStoreMaxLength(l int) FileStoreOptionFunc {
	return func(fs *sessions.FilesystemStore) error {
		fs.MaxLength(l)
		return nil
	}
}

func SetFileStoreMaxAge(age int) FileStoreOptionFunc {
	return func(fs *sessions.FilesystemStore) error {
		if fs.Options != nil {
			fs.MaxAge(age)
		}
		return nil
	}
}

func SetFileStoreOption(opt *sessions.Options) FileStoreOptionFunc {
	return func(fs *sessions.FilesystemStore) error {
		if opt != nil {
			fs.Options = opt
			fs.MaxAge(opt.MaxAge)
		}
		return nil
	}
}

func NewRedisStore(maxIdle int, network, address, password string, keyPairs []string) (*redistore.RediStore, error) {
	keyPairsArr := make([][]byte, 0, len(keyPairs))
	for _, v := range keyPairs {
		keyPairsArr = append(keyPairsArr, []byte(v))
	}
	return redistore.NewRediStore(maxIdle, network, address, password, keyPairsArr...)
}
