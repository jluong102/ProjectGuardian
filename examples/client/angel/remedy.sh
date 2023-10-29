#!/bin/bash
filepath=$(find / -type f -name check.sh 2>/dev/null)
touch ${filepath%check.sh}/flag