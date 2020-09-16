#!/bin/bash
cp *.service *.timer /etc/systemd/system/
systemctl daemon-reload
systemctl enable --now dc@5pmGebetServer.service
systemctl enable --now dc-reload@5pmGebetServer.timer
