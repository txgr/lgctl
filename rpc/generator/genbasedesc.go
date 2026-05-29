package generator

import (
	"path/filepath"
	"strings"

	"github.com/iancoleman/strcase"

	conf "github.com/txgr/lgctl/config"
	"github.com/txgr/lgctl/rpc/parser"
	"github.com/txgr/lgctl/util"
	"github.com/txgr/lgctl/util/pathx"
)

func (g *Generator) GenBaseDesc(ctx DirContext, _ parser.Proto, cfg *conf.Config, c *ZRpcContext) error {
	protoFilename := filepath.Join(ctx.GetMain().Filename, "desc", "base.proto")
	if err := pathx.MkdirIfNotExist(filepath.Join(ctx.GetMain().Filename, "desc")); err != nil {
		return err
	}

	err := util.With("t").Parse(rpcTemplateText).SaveTo(map[string]string{
		"package":     strings.ToLower(c.RpcName),
		"serviceName": strcase.ToCamel(c.RpcName),
	}, protoFilename, false)
	return err
}
