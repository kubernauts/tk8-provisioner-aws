package provisioner

import (
	"github.com/kubernauts/tk8-provisioner-aws/internal/cluster"
	"github.com/kubernauts/tk8/pkg/provisioner"
)

type AWS struct {
}

func (p AWS) Init(args []string) {
	cluster.KubesprayInit()
	cluster.AWSCreate()
}

func (p AWS) Setup(args []string) {
	cluster.AWSInstall()

}

func (p AWS) Scale(args []string) {
	cluster.AWSScale()

}

func (p AWS) Reset(args []string) {
	cluster.AWSReset()

}

func (p AWS) Remove(args []string) {
	cluster.AWSRemove()

}

func (p AWS) Upgrade(args []string) {
	provisioner.NotImplemented()
}

func (p AWS) Destroy(args []string) {
	cluster.AWSDestroy()
}

func NewAWS() provisioner.Provisioner {
	cluster.SetClusterName()
	provisioner := new(AWS)
	return provisioner
}
