package config

import (
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"time"
)

type KsniffSettings struct {
	UserSpecifiedPodName           string
	UserSpecifiedInterface         string
	UserSpecifiedFilter            string
	UserSpecifiedPodCreateTimeout  time.Duration
	UserSpecifiedContainer         string
	UserSpecifiedNamespace         string
	UserSpecifiedOutputFile        string
	UserSpecifiedLocalTcpdumpPath  string
	UserSpecifiedRemoteTcpdumpPath string
	UserSpecifiedVerboseMode       bool
	UserSpecifiedPrivilegedMode    bool
	UserSpecifiedImage             string
	DetectedPodNodeName            string
	DetectedContainerId            string
	DetectedContainerRuntime       string
	Image                          string
	UseDefaultImage                bool
	UserSpecifiedKubeContext       string
}

func NewKsniffSettings(streams genericclioptions.IOStreams) *KsniffSettings {
	return &KsniffSettings{}
}
