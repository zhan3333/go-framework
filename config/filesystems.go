package config

type filesystems struct {
	Default string
	Cloud   string
	Disks   Disks
}

type Disks struct {
	Local Local
}

type Local struct {
	Driver string
	Root   string
}
