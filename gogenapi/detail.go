package gogenapi

type Detail struct {
	VCS       string
	User      string
	Project   string
	Namespace string
	Models    Models
	Model     *Model
	ImportDir string
	Database  string
}
