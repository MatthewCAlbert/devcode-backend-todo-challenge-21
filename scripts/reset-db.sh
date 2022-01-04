#!/bin/bash

mysql -u root -p12345 << EOF
DROP DATABASE \`devcode-todo\`; 
CREATE DATABASE \`devcode-todo\`;
EOF

