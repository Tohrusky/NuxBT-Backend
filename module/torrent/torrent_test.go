package torrent

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const testTorrent = "test_torrents/test.torrent"
const testTorrentFolder = "test_torrents/test_folder.torrent"
const testTorrentSave = "test_torrents/test_save.torrent"

func TestTorrentHash(t *testing.T) {
	torrentFilePath := testTorrent

	torrent, err := NewBitTorrentFilePath(torrentFilePath)
	if err != nil {
		t.Error(err)
		return
	}

	err = torrent.Repack(&BitTorrentFileEditStrategy{})
	if err != nil {
		t.Error(err)
		return
	}
	assert.Equal(t, torrent.GetHash(), "7f3956e5a15b34b62159727c08f944c7e433ad1e")
	assert.Equal(t, torrent.Info.Name, "lenna.jpg")
}

func TestFolderTorrentHash(t *testing.T) {
	torrentFilePath := testTorrentFolder

	torrent, err := NewBitTorrentFilePath(torrentFilePath)
	if err != nil {
		t.Error(err)
		return
	}

	err = torrent.Repack(&BitTorrentFileEditStrategy{})
	if err != nil {
		t.Error(err)
		return
	}
	assert.Equal(t, torrent.GetHash(), "0f8cd84ebb514a4d6975f217c1df129bba080a01")
	assert.Equal(t, torrent.Info.Name, "cxkcxk")
}

func TestRepackTorrent(t *testing.T) {
	torrentFilePath := testTorrent

	torrent, err := NewBitTorrentFilePath(torrentFilePath)
	if err != nil {
		t.Error(err)
		return
	}
	comment := "TensoRaws"
	infoSource := "https://github.com/TensoRaws"

	err = torrent.Repack(&BitTorrentFileEditStrategy{
		Comment:    comment,
		InfoSource: infoSource,
	})

	if err != nil {
		t.Error(err)
		return
	}

	assert.Equal(t, torrent.Comment, comment)
	assert.Equal(t, torrent.Info.Source, infoSource)
}

func TestSaveTorrent(t *testing.T) {
	torrentFilePath := testTorrent
	saveFilePath := testTorrentSave

	torrent, err := NewBitTorrentFilePath(torrentFilePath)
	if err != nil {
		t.Error(err)
		return
	}

	err = torrent.SaveTo(saveFilePath)
	if err != nil {
		t.Error(err)
		return
	}

	torrent, err = NewBitTorrentFilePath(saveFilePath)
	if err != nil {
		t.Error(err)
		return
	}

	assert.Equal(t, torrent.GetHash(), "7f3956e5a15b34b62159727c08f944c7e433ad1e")
	assert.Equal(t, torrent.Info.Name, "lenna.jpg")
}