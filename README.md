# Proxmox ZFS Exporter

## Description

Metrics exporter for Prometheus. Exposes metrics to help monitor the health of ZFS pools across a proxmox cluster. The default port is set to 9000.

The code was written quickly and needs to be cleaned up but it's functional. Lots of work could be done here and I always welcome contributions.

## Exported Metrics
| name  | Description |
| ------------- | ------------- |
| zfs_zpool_error | Is there a zpool error |
| zfs_zpool_online | Is the zpool online |
| zfs_zpool_free | Free space on zpool |
| zfs_zpool_allocated | Allocated space on zpool |
| zfs_zpool_size | Size of zpool |
| zfs_zpool_dedup | Is dedup enabled on zpool |
| zfs_zpool_last_scrub | Last zpool scrub |
| zfs_zpool_last_scrub_errors |Last scrub total errors on the zpool |
| zfs_zpool_parsing_error | Error when trying to parse the API data. |

## Installing

You'll need to create a proxmox user that can access the API. You can limit the permissions down to read only. Then create a config file and store it at `etc/proxmox-zfs-exporter/config.json` with content like:

``` json
{
  "User": "sampleuser@pam",
  "Pass": "samplepassword",
  "Host": "192.168.0.20",
  "Port": "8006"
}
```

NOTE: The port number is the port number to proxmox NOT the port for the metrics.

You can either compile the binary or download one the releases.

Example install using systemd (Make sure to get the current release from the release page and change the wget command):

``` bash
#Run this as root or add sudo to the commands
wget https://github.com/zbblanton/proxmox-zfs-exporter/releases/download/v0.1.1/proxmox-zfs-exporter -O /usr/bin/proxmox-zfs-exporter
chmod +x /usr/bin/proxmox-zfs-exporter
cat > /etc/systemd/system/proxmox-zfs-exporter.service <<EOF
[Unit]
Description=Proxmox ZFS Exporter
Documentation=https://github.com/zbblanton/proxmox-zfs-exporter
[Service]
ExecStart=/usr/bin/proxmox-zfs-exporter
Restart=on-failure
RestartSec=5
[Install]
WantedBy=multi-user.target
EOF
systemctl daemon-reload
systemctl enable proxmox-zfs-exporter
systemctl start proxmox-zfs-exporter
```
