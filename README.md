# bencode2gostruct
A simple utility to parse the fields from a bencode encoded file,
into a go `struct`.

```
Usage of bencode2gostruct:
  -in string
        Bencoded file to generate model from
  -model-name string
        Model name for the struct
  -out string
        Go model file to output
  -package package
        The package name for the Go file
```

Say you were to parse a qBittorrent `fastresume` file, you would
get the following output

```
$ bencode2gostruct -in sample.fastresume -out model.go -model-name FastResumeData -package fastresume

package fastresume

type FastResumeData struct {
        ActiveTime                  int64         `bencode:"active_time"`
        AddedTime                   int64         `bencode:"added_time"`
        Allocation                  string        `bencode:"allocation"`
        ApplyIpFilter               int64         `bencode:"apply_ip_filter"`
        AutoManaged                 int64         `bencode:"auto_managed"`
        CompletedTime               int64         `bencode:"completed_time"`
        DisableDht                  int64         `bencode:"disable_dht"`
        DisableLsd                  int64         `bencode:"disable_lsd"`
        DisablePex                  int64         `bencode:"disable_pex"`
        DownloadRateLimit           int64         `bencode:"download_rate_limit"`
        FileFormat                  string        `bencode:"file-format"`
        FileVersion                 int64         `bencode:"file-version"`
        FilePriority                []interface{} `bencode:"file_priority"`
        FinishedTime                int64         `bencode:"finished_time"`
        Httpseeds                   []interface{} `bencode:"httpseeds"`
        I2p                         int64         `bencode:"i2p"`
        InfoHash                    string        `bencode:"info-hash"`
        InfoHash2                   string        `bencode:"info-hash2"`
        LastDownload                int64         `bencode:"last_download"`
        LastSeenComplete            int64         `bencode:"last_seen_complete"`
        LastUpload                  int64         `bencode:"last_upload"`
        LibtorrentVersion           string        `bencode:"libtorrent-version"`
        MaxConnections              int64         `bencode:"max_connections"`
        MaxUploads                  int64         `bencode:"max_uploads"`
        Name                        string        `bencode:"name"`
        NumComplete                 int64         `bencode:"num_complete"`
        NumDownloaded               int64         `bencode:"num_downloaded"`
        NumIncomplete               int64         `bencode:"num_incomplete"`
        Paused                      int64         `bencode:"paused"`
        Peers                       string        `bencode:"peers"`
        Peers6                      string        `bencode:"peers6"`
        PiecePriority               string        `bencode:"piece_priority"`
        Pieces                      string        `bencode:"pieces"`
        QBtCategory                 string        `bencode:"qBt-category"`
        QBtContentLayout            string        `bencode:"qBt-contentLayout"`
        QBtDownloadPath             string        `bencode:"qBt-downloadPath"`
        QBtFirstLastPiecePriority   int64         `bencode:"qBt-firstLastPiecePriority"`
        QBtInactiveSeedingTimeLimit int64         `bencode:"qBt-inactiveSeedingTimeLimit"`
        QBtName                     string        `bencode:"qBt-name"`
        QBtRatioLimit               int64         `bencode:"qBt-ratioLimit"`
        QBtSavePath                 string        `bencode:"qBt-savePath"`
        QBtSeedStatus               int64         `bencode:"qBt-seedStatus"`
        QBtSeedingTimeLimit         int64         `bencode:"qBt-seedingTimeLimit"`
        QBtStopCondition            string        `bencode:"qBt-stopCondition"`
        QBtTags                     []interface{} `bencode:"qBt-tags"`
        SavePath                    string        `bencode:"save_path"`
        SeedMode                    int64         `bencode:"seed_mode"`
        SeedingTime                 int64         `bencode:"seeding_time"`
        SequentialDownload          int64         `bencode:"sequential_download"`
        ShareMode                   int64         `bencode:"share_mode"`
        StopWhenReady               int64         `bencode:"stop_when_ready"`
        SuperSeeding                int64         `bencode:"super_seeding"`
        TotalDownloaded             int64         `bencode:"total_downloaded"`
        TotalUploaded               int64         `bencode:"total_uploaded"`
        Trackers                    []interface{} `bencode:"trackers"`
        Unfinished                  []interface{} `bencode:"unfinished"`
        UploadMode                  int64         `bencode:"upload_mode"`
        UploadRateLimit             int64         `bencode:"upload_rate_limit"`
        UrlList                     []interface{} `bencode:"url-list"`
}
```

## Building
With go installed

```
git clone https://github.com/jslay88/bencode2gostruct.git
cd bencode2gostruct
go build -o ./bin/bencode2gostruct ./cmd/bencode2gostruct
```

