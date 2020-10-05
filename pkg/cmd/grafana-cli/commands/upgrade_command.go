package commands

import (
	"github.com/fatih/color"
	"gitlab.com/digitalizm/grafana/pkg/cmd/grafana-cli/logger"
	"gitlab.com/digitalizm/grafana/pkg/cmd/grafana-cli/services"
	"gitlab.com/digitalizm/grafana/pkg/cmd/grafana-cli/utils"
	"gitlab.com/digitalizm/grafana/pkg/util/errutil"
)

func (cmd Command) upgradeCommand(c utils.CommandLine) error {
	pluginsDir := c.PluginDirectory()
	pluginName := c.Args().First()

	localPlugin, err := services.ReadPlugin(pluginsDir, pluginName)

	if err != nil {
		return err
	}

	plugin, err2 := cmd.Client.GetPlugin(pluginName, c.RepoDirectory())
	if err2 != nil {
		return err2
	}

	if shouldUpgrade(localPlugin.Info.Version, &plugin) {
		if err := services.RemoveInstalledPlugin(pluginsDir, pluginName); err != nil {
			return errutil.Wrapf(err, "failed to remove plugin '%s'", pluginName)
		}

		return InstallPlugin(pluginName, "", c, cmd.Client)
	}

	logger.Infof("%s %s is up to date \n", color.GreenString("✔"), pluginName)
	return nil
}
