package gogen

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/txgr/lgctl/api/spec"
	"github.com/txgr/lgctl/config"
	"github.com/txgr/lgctl/util/format"
	"github.com/txgr/lgctl/util/pathx"
	"github.com/txgr/lgctl/vars"
)

//go:embed main.tpl
var mainTemplate string

func genMain(dir, rootPkg string, cfg *config.Config, api *spec.ApiSpec, g *GenContext) error {
	filename, err := format.FileNamingFormat(cfg.NamingFormat, api.Service.Name)
	if err != nil {
		return err
	}

	configName := filename
	if strings.HasSuffix(filename, "-api") {
		filename = strings.ReplaceAll(filename, "-api", "")
	}

	return genFile(fileGenConfig{
		dir:             dir,
		subdir:          "",
		filename:        filename + ".go",
		templateName:    "mainTemplate",
		category:        category,
		templateFile:    mainTemplateFile,
		builtinTemplate: mainTemplate,
		data: map[string]any{
			"importPackages": genMainImports(rootPkg),
			"serviceName":    configName,
			"port":           g.Port,
		},
	})
}

func genMainImports(parentPkg string) string {
	var imports []string
	imports = append(imports, fmt.Sprintf("\"%s\"", pathx.JoinPackages(parentPkg, configDir)))
	imports = append(imports, fmt.Sprintf("\"%s\"", pathx.JoinPackages(parentPkg, handlerDir)))
	imports = append(imports, fmt.Sprintf("\"%s\"\n", pathx.JoinPackages(parentPkg, contextDir)))
	imports = append(imports, fmt.Sprintf("\"%s/core/conf\"", vars.ProjectOpenSourceURL))
	imports = append(imports, fmt.Sprintf("\"%s/rest\"", vars.ProjectOpenSourceURL))
	return strings.Join(imports, "\n\t")
}
