# Containerlab VS Code Extension

The lab-as-code approach taken by Containerlab means labs are "written" in YAML in a text editor or IDE. It also means you have to manage your labs via the command-line.

[VS Code](https://code.visualstudio.com/download) is a powerful text editor used by many for this purpose, and with the YAML schema provided by Containerlab the topology writing experience is made even easier.

We decided to further improve the experience with VS Code with a Containerlab [VS Code extension](https://marketplace.visualstudio.com/items?itemName=srl-labs.vscode-containerlab).

The Containerlab VS Code extension aims to greatly simplify and improve the labbing workflow, allowing you to perform essential lab operations within VS Code.

If you prefer to sit back and watch a video about this extension, say no more!

-{{youtube(url='https://www.youtube.com/embed/NIw1PbfCyQ4')}}-

/// admonition | Support
If you have any questions about this extension, please join us in the [Containerlab Discord](https://discord.gg/vAyddtaEV9).
///

## Features

The Containerlab extension packs a lot of features hidden under the icons or context menus. Here is a list of the main features, and we are always looking for new feature requests:

### Explorer

In the activity bar of VS Code, you will notice a Containerlab icon. Clicking on this icon will open the explorer.

The explorer is similar to the VS Code file explorer but instead of finding files, it discovers Containerlab topologies, running labs, containers and their interfaces.

The explorer will discover all labs in your local directory (and subdirectories), as well as any running labs on the system.

The explorer is a Treeview and for running labs, you can expand the labs to see running containers and you can expand running containers to see their interfaces.

### Editor

In the editor title actions when a Containerlab topology is open, there is a 'run' action and a 'graph' button.

This allows easy deployment and graphing of the lab from within the editor context.

![editor-actions](https://gitlab.com/rdodin/pics/-/wikis/uploads/c095d277e032ad0754b5f61f4a9057f2/CleanShot_2025-02-04_at_13.56.21_2x.png)

#### Command palette

When you have a topology open and active in the editor you can execute common actions from the command palette that you can open up with `CTRL+SHIFT+P`/`CMD+SHIFT+P`.

![palette](https://gitlab.com/rdodin/pics/-/wikis/uploads/be8d2c2d7d806dedd76c2b7a924d153d/CleanShot_2025-02-04_at_13.58.39_2x.png)

#### Keybindings

We have also set some default keybindings you can use to interact with the lab when you are editing a topology.

| Keys         | Action   |
| ------------ | -------- |
| `CTRL+ALT+D` | Deploy   |
| `CTRL+ALT+R` | Redeploy |
| `CTRL+ALT+K` | Destroy  |
| `CTRL+ALT+G` | Graph    |

### TopoViewer

Integrated into the extension as a 'graph' action is the [TopoViewer](https://github.com/asadarafat/topoViewer) project by @asadarafat.

TopoViewer is an interactive way to visualize your containerlab topologies. The topology must be running for the TopoViewer to be able to visualize it.

### Packet capture

In the explorer you can expand running labs and containers to view all the discovered interfaces for that container. You can either right click on the interface and start a packet capture or click the shark icon on the interface label.

Packet capture relies on the [Edgeshark integration](wireshark.md#edgeshark-integration). Install Edgeshark first, and then use the context menu or the fin icon next to the interface name.

## Settings reference

Below is a reference to the available settings that can be modified in the Containerlab VS Code extension.

### `containerlab.defaultSshUser`

The default username used to connect to a node via SSH.

| Type     | Default |
| -------- | ------- |
| `string` | `admin` |

### `containerlab.sudoEnabledByDefault`

Whether or not to append `sudo` to any commands executed by the extension.

| Type      | Default |
| --------- | ------- |
| `boolean` | `true`  |

### `containerlab.refreshInterval`

The time interval (in milliseconds) for which the extension automatically refreshes.

On each refresh the extension will discover all local and running labs, as well as containers and interfaces of containers belonging to running labs.

By default this is 10 seconds.

| Type     | Default |
| -------- | ------- |
| `number` | `10000` |

### `containerlab.node.execCommandMapping`

The a mapping between the node kind and command executed on the 'node attach' action.

The 'node attach' action performs the `docker exec -it <node> <command>` command. The `execCommandMapping` allows you to change the command executed by default on a per-kind basis.

By default this setting is empty, it should be used to override the [default mappings](https://github.com/srl-labs/vscode-containerlab/blob/main/resources/exec_cmd.json).

| Type     | Default     |
| -------- | ----------- |
| `object` | `undefined` |

#### Example

```json
{
    "nokia_srl": "sr_cli",
}
```

In the settings UI, simply set the 'Item' field to the kind and the 'Value' field to `nokia_srl` and the command to `sr_cli`.
