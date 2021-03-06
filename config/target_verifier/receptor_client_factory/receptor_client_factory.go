package receptor_client_factory

import (
	"github.com/cloudfoundry-incubator/receptor"
)

func MakeReceptorClient(target string) receptor.Client {
	return receptor.NewClient(target)
}
