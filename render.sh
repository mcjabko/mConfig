#!/bin/bash
template_file=$1
output_file=$2

if [[ -z $template_file ]] || [[ -z $output_file ]]; then
  echo "Usage: render <template_file> <output_file>"
fi

if [ ! -f "$template_file" ]; then
  echo "Template file doesn't exits"
fi

if [ ! -f "$output_file" ]; then
  echo >"$output_file"
  else
    rm "$output_file"
fi

function load_default_env() {
    set -o allexport
    source .env.default set
}

function render_file() {
    while read -r line; do
      sed -i '' "s/{$line}/${!line}/" "$output_file"
    done < <(cat .env.default | awk -F'=' '{print $1}')
}


# Load default env
load_default_env
cp -v "$template_file" "$output_file"
render_file
kl
echo "File rendered successfuly!"
