// See http://www.bittorrent.com/intl/ru/sync/developers/api
// Inspired by https://github.com/vole/btsync-api
package btsync

import (
	"net/url"
	"strconv"
)

// Returns an array with folders info.
// http://[address]:[port]/api?method=get_folders[&secret=(secret)]
// - secret (optional) - if a secret is specified, will return info about the folder with this secret
func (c Client) Folders() (Folders, error) {
	v := &Folders{}
	err := c.call("get_folders", nil, v)
	return *v, err
}

// Returns Folder info about the folder with this secret.
//http://[address]:[port]/api?method=get_folders[&secret=(secret)]
func (c Client) Folder(secret string) (*Folder, error) {
	v := &Folders{}
	err := c.call("get_folders", url.Values{"secret": {secret}}, v)
	if v != nil {
		return (*v)[0], err
	}

	return nil, err
}

// Adds a folder to Sync. If a secret is not specified, it will be generated automatically. The folder will have to pre-exist on the disk and Sync will add it into a list of syncing folders.
// Returns '0' if no errors, error code and error message otherwise.
// http://[address]:[port]/api?method=add_folder&dir=(folderPath)[&secret=(secret)&selective_sync=1]
// - dir (required) - specify path to the sync folder
// - secret (optional) - specify folder secret
// - selective_sync (optional) - specify sync mode, selective - 1, all files (default) - 0
func (c Client) AddFolder(dir, secret string, selectiveSync int) (*OperationResult, error) {
	v := &OperationResult{}
	p := url.Values{"dir": {dir}, "selective_sync": {strconv.Itoa(selectiveSync)}}
	if secret != "" {
		p.Add( "secret", secret )
	}

	err := c.call("add_folder", p, v)
	return v, err
}

// Removes folder from Sync while leaving actual folder and files on disk. It will remove a folder from the Sync
// list of folders and does not touch any files or folders on disk. Returns '0' if no error, '1' if there’s no folder with specified secret.
// { "error": 0 }
// http://[address]:[port]/api?method=remove_folder&secret=(secret)
// - secret (required) - specify folder secret
func (c Client) RemoveFolder(secret string) (*OperationResult, error) {
	v := &OperationResult{}
	err := c.call("remove_folder", url.Values{"secret": {secret}}, v)
	return v, err
}

// Returns list of files within the specified directory.
// If a directory is not specified, will return list of files and folders within the root folder.
// Note that the Selective Sync function is only available in the API at this time.
// http://[address]:[port]/api?method=get_files&secret=(secret)[&path=(path)]
// - secret (required) - must specify folder secret
// - path (optional) - specify path to a subfolder of the sync folder.
func (c Client) Files(secret, path string) {
	// TODO
}

// Selects file for download for selective sync folders. Returns file information with applied preferences.
// http://[address]:[port]/api?method=set_file_prefs&secret=(secret)&path=(path)&download=1
// - secret (required) - must specify folder secret
// - path (required) - specify path to a subfolder of the sync folder.
// - download (required) - specify if file should be downloaded (yes - 1, no - 0)
func (c Client) SelectFile(secret, path string, download bool) {
	// TODO
}

// Returns list of peers connected to the specified folder.
// http://[address]:[port]/api?method=get_folder_peers&secret=(secret)
// - secret (required) - must specify folder secret
func (c Client) FolderPeers(secret string) {
	// TODO
}

// Generates read-write, read-only and encryption read-only secrets. If ‘secret’ parameter is specified, will return secrets available for sharing under this secret.
// The Encryption Secret is new functionality. This is a secret for a read-only peer with encrypted content (the peer can sync files but can not see their content).
// One example use is if a user wanted to backup files to an untrusted, unsecure, or public location. This is set to disabled by default for all users but included in the API.
// http://[address]:[port]/api?method=get_secrets[&secret=(secret)&type=encryption]
// - secret (required) - must specify folder secret
// - type (optional) - if type=encrypted, generate secret with support of encrypted peer
func (c Client) Secrets(secret string, encrypted bool) (*Secrets, error) {
	v := &Secrets{}
	p := url.Values{}

	if secret != "" {
		p.Add("secret", secret)
	}

	if encrypted {
		p.Add("type", "encryption")
	}

	err := c.call("get_secrets", p, v)
	return v, err
}

// Returns preferences for the specified sync folder.
// http://[address]:[port]/api?method=get_folder_prefs&secret(secret)
// - secret (required) - must specify folder secret
func (c Client) FolderPreferences(secret string) {
	// TODO
}

// Sets preferences for the specified sync folder. Parameters are the same as in ‘Get folder preferences’. Returns current settings.
// http://[address]:[port]/api?method=set_folder_prefs&secret=(secret)&param1=value1&param2=value2,...
// - secret (required) - must specify folder secret
// - params - { use_dht, use_hosts, search_lan, use_relay_server, use_tracker, use_sync_trash }
func (c Client) SetFolderPreferences(secret string) {
	// TODO
}

// Returns list of predefined hosts for the folder, or error code if a secret is not specified.
// http://[address]:[port]/api?method=get_folder_hosts&secret=(secret)
// - secret (required) - must specify folder secret
func (c Client) FolderHosts(secret string) {
	// TODO
}

// Set folder hosts
// Sets one or several predefined hosts for the specified sync folder. Existing list of hosts will be replaced. Hosts should be added as values of the ‘host’ parameter and separated by commas.
// Returns current hosts if set successfully, error code otherwise.
// http://[address]:[port]/api?method=set_folder_hosts&secret=(secret)&hosts=host1:port1,host2:port2,...
// - secret (required) - must specify folder secret
// - hosts (required) - enter list of hosts separated by comma. Host should be represented as “[address]:[port]”
func (c Client) SetFolderHosts(secret string, hosts []string) {
	// TODO
}

// Returns BitTorrent Sync preferences. Contains dictionary with advanced preferences. Please see Sync user guide for description of each option.
// http://[address]:[port]/api?method=get_prefs
func (c Client) Preferences() (*Preferences, error) {
	v := &Preferences{}
	err := c.call("get_prefs", nil, v)
	return v, err
}

// Set preferences
// Sets BitTorrent Sync preferences. Parameters are the same as in ‘Get preferences’. Advanced preferences are set as general settings. Returns current settings.
// http://[address]:[port]/api?method=set_prefs&param1=value1&param2=value2,...
// - params - { device_name, download_limit, lang, listening_port, upload_limit, use_upnp } and advanced settings. You can find more information about advanced settings in user guide.
func (c Client) SetPreferences() {
	// TODO
}

// Returns OS name where BitTorrent Sync is running.
// http://[address]:[port]/api?method=get_os
func (c Client) OSName() (*OS, error) {
	v := &OS{}
	err := c.call("get_os", nil, v)
	return v, err
}

// Returns BitTorrent Sync version.
// http://[address]:[port]/api?method=get_version
func (c Client) Version() (*Version, error) {
	v := &Version{}
	err := c.call("get_version", nil, v)
	return v, err
}

// Returns current upload and download speed.
// http://[address]:[port]/api?method=get_speed
func (c Client) Speed() (*Speed, error) {
	v := &Speed{}
	err := c.call("get_speed", nil, v)
	return v, err
}

// Gracefully stops Sync.
// http://[address]:[port]/api?method=shutdown
func (c Client) Shutdown() (*OperationResult, error) {
	v := &OperationResult{}
	err := c.call("shutdown", nil, v)
	return v, err
}
