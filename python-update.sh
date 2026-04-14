#!/bin/bash
set -e

# Define version to install
PY_VER="3.13"

echo "Adding deadsnakes PPA..."
sudo add-apt-repository -y ppa:deadsnakes/ppa
sudo apt update

echo "Installing Python $PY_VER..."
sudo apt install -y python$PY_VER python$PY_VER-venv python$PY_VER-full

echo "Verification:"
python$PY_VER --version
echo "Use 'python$PY_VER -m venv .venv' to start projects with this version."