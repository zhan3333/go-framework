package conf

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

var Filesystems = filesystems{
	Default: "local",
	Cloud:   "",
	Disks: Disks{
		Local: Local{
			Driver: "local",
			Root:   "app/public",
		},
	},
}
