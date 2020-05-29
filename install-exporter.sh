#!/bin/sh

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