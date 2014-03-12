package btsync

// BTSync Client
type Client struct {
	Host, Port, User, Password string
}

type Preferences struct {
	DeviceName                 string `json:"device_name"`
	DiskLowPriority            bool   `json:"disk_low_priority"`
	DownloadLimit              int    `json:"download_limit"`
	FolderRescanInterval       int    `json:"folder_rescan_interval"`
	LANEncryptData             bool   `json:"lan_encrypt_data"`
	LANUseTcp                  bool   `json:"lan_use_tcp"`
	Lang                       int    `json:"lang"`
	ListeningPort              int    `json:"listening_port"`
	MaxFileSizeDiffForPatching int64  `json:"max_file_size_diff_for_patching"`
	MaxFileSizeForVersioning   int64  `json:"max_file_size_for_versioning"`
	RateLimitLocalPeers        bool   `json:"rate_limit_local_peers"`
	RecvBufSize                int64  `json:"recv_buf_size"`
	SendBufSize                int64  `json:"send_buf_size"`
	SyncMaxTimeDiff            int64  `json:"sync_max_time_diff"`
	SyncTrashTtl               int64  `json:"sync_trash_ttl"`
	UploadLimit                int    `json:"upload_limit"`
	UseUPnP                    bool   `json:"use_upnp"`
}

type Folder struct {
	Dir      string `json:"dir"`
	Secret   string `json:"secret"`
	Size     int64  `json:"size"`
	Type     string `json:"type"`
	Files    int64  `json:"files"`
	Error    int    `json:"error"`
	Indexing int    `json:"indexing"`
}

type Secrets struct {
	ReadOnly   string `json:"read_only"`
	ReadWrite  string `json:"read_write"`
	Encryption string `json:"encryption"`
}

type OS struct {
	Name string `json:"os"`
}

type Version struct {
	Version string `json:"version"`
}

type Speed struct {
	Download int64 `json:"download"`
	Upload   int64 `json:"upload"`
}

type OperationResult struct {
	Error   int    `json:"error"`
	Result   int   `json:"result"`
	Message string `json:"message"`
}

type Folders []*Folder
