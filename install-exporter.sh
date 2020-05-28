#!/bin/sh

cat > /etc/systemd/system/proxmox-zfs-exporter.service <<EOF
[Unit]
Description=Proxmox ZFS Exporter
Documentation=https://github.com/zbblanton/proxmox-zfs-exporter
[Service]
ExecStart=/usr/local/bin/proxmox-zfs-exporter
Restart=on-failure
RestartSec=5
[Install]
WantedBy=multi-user.target
EOF

systemctl daemon-reload
sudo systemctl enable proxmox-zfs-exporter
sudo systemctl start proxmox-zfs-exporter