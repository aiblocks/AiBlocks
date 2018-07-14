// Copyright 2016 The go-aiblocks Authors
// This file is part of the go-aiblocks library.
//
// The go-aiblocks library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-aiblocks library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-aiblocks library. If not, see <http://www.gnu.org/licenses/>.

// +build ios linux,arm64 !darwin,!freebsd,!linux,!netbsd,!solaris,!windows

// This is the fallback implementation of directory watching.
// It is used on unsupported platforms.

package accounts

type watcher struct{ running bool }

func newWatcher(*addrCache) *watcher { return new(watcher) }
func (*watcher) start()              {}
func (*watcher) close()              {}
