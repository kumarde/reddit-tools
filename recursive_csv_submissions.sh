#!/bin/bash

# Directory where the zst files are located
target_dir=$1

# Function to uncompress zst files recursively
uncompress_recursive() {
  local dir="$1"

  # Find zst files in the directory
  zst_files=$(find "$dir" -name "RS_201[8-9]*.zst" -o -name "RS_202[0-4]*.zst")
  # zst_files=$(find "$dir" -name "RS_2018-09.zst")

  # Uncompress each file
  for zst_file in $zst_files; do
    echo "Uncompressing: $zst_file"
  
    # Uncompress the zst file
    if [ -f "${zst_file%.zst}.json" ]; then
      echo "${zst_file%.zst}.json exists"
    else
        if [ -f "${zst_file%.zst}.csv" ]; then
        echo "${zst_file%.zst}.zst already parsed"
        else
        zstd -d "$zst_file" -o "${zst_file%.zst}.json" --long=31
        fi
    fi
    if [ -f "${zst_file%.zst}.csv" ]; then
      echo "${zst_file%.zst}.csv exists"
	  else
      ./flatten_submission_json/flatten_submission_json "${zst_file%.zst}"
    fi
    rm -rf "${zst_file%.zst}.json"
	done
}

# Start the process from the target directory
uncompress_recursive "$target_dir"

echo "All zst files uncompressed."
