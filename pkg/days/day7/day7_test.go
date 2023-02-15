package day7

import (
	"testing"
)

func TestFileSystem_changeDirectory(t *testing.T) {
	childDir := &Directory{
		Path: []string{"childDir"},
		Size: 0,
	}
	rootDir := &Directory{
		Subdirectories: map[string]*Directory{
			"childDir": childDir,
		},
	}
	childDir.Parent = rootDir
	type fields struct {
		Root             *Directory
		CurrentDirectory *Directory
	}
	type args struct {
		d string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		wantDir *Directory
	}{
		{
			name: "fromChildToRoot",
			fields: fields{
				Root:             rootDir,
				CurrentDirectory: childDir,
			},
			args:    args{"/"},
			wantErr: false,
			wantDir: rootDir,
		},
		{
			name: "fromRootToChild",
			fields: fields{
				Root:             rootDir,
				CurrentDirectory: rootDir,
			},
			args:    args{"childDir"},
			wantErr: false,
			wantDir: childDir,
		},
		{
			name: "fromChildToParent",
			fields: fields{
				Root:             rootDir,
				CurrentDirectory: childDir,
			},
			args:    args{".."},
			wantErr: false,
			wantDir: rootDir,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &FileSystem{
				Root:             tt.fields.Root,
				CurrentDirectory: tt.fields.CurrentDirectory,
			}
			if err := fs.changeDirectory(tt.args.d); (err != nil) != tt.wantErr {
				t.Errorf("FileSystem.changeDirectory() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantDir != fs.CurrentDirectory {
				t.Errorf("%v is not %v", tt.wantDir, fs.CurrentDirectory)
			}
		})
	}
}

func Test_setSizes(t *testing.T) {
	type args struct {
		dir *Directory
	}
	tests := []struct {
		name     string
		args     args
		wantSize int
	}{
		{
			name: "EmptySubdirs",
			args: args{
				dir: &Directory{
					Subdirectories: map[string]*Directory{
						"a": {Size: 1},
						"b": {Size: 2},
						"c": {Size: 3},
					},
					Size: 0,
				},
			},
			wantSize: 6,
		},
		{
			name: "SubdirsWithSubdirs",
			args: args{
				dir: &Directory{
					Subdirectories: map[string]*Directory{
						"a": {
							Size: 1,
							Subdirectories: map[string]*Directory{
								"a1": {Size: 1},
							},
						},
						"b": {
							Size: 2,
							Subdirectories: map[string]*Directory{
								"b1": {Size: 2},
							},
						},
						"c": {
							Size: 3,
							Subdirectories: map[string]*Directory{
								"c1": {Size: 3},
							},
						},
					},
					Size: 0,
				},
			},
			wantSize: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := setSizes(tt.args.dir)
			if err != nil {
				t.Errorf("setSizes() error = %v", err)
			}

			if tt.args.dir.Size != tt.wantSize {
				t.Errorf("rootDir size %v != %v", tt.args.dir.Size, tt.wantSize)
			}
		})
	}
}
