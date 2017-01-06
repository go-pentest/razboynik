package kernelmodule

import (
	"errors"

	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboynik/services/kernel"
	"github.com/eatbytes/razboynik/services/lister"
	"github.com/eatbytes/razboynik/services/provider"
)

type Plugincmd struct{}

func (plugin *Plugincmd) Exec(kl *kernel.KernelLine, config *razboy.Config) kernel.KernelResponse {
	var (
		arr  []string
		err  error
		info *provider.Info
		args *provider.Args
		resp *provider.Response
	)

	arr = kl.GetArg()

	if len(arr) < 1 {
		return kernel.KernelResponse{Err: errors.New("Please specify a path for external plugin")}
	}

	args = new(provider.Args)
	args.Line = kl.GetStr()

	info = &provider.Info{
		Path:   arr[0],
		Method: provider.EXEC_FN,
	}

	resp, err = provider.CallProvider(info, args)

	if err != nil {
		return kernel.KernelResponse{Err: err}
	}

	kernel.WriteSuccess(kl.GetStdout(), resp.Content)

	return kernel.KernelResponse{Body: resp.Content}
}

func (plugin *Plugincmd) GetName() string {
	return "-plugin"
}

func (plugin *Plugincmd) GetCompleter() (kernel.CompleterFunction, bool) {
	return lister.Local, true
}
