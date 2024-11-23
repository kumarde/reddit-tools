#!/bin/bash

# Directory where the zst files are located
target_dir=$1

# Function to uncompress zst files recursively
uncompress_recursive() {
  local dir="$1"

  # Find zst files in the directory
  zst_files=$(find "$dir" -name "RC_201[8-9]*.zst" -o -name "RC_202[0-4]*.zst")
  # zst_files=$(find "$dir" -name "RC_2018-09.zst")

  # Uncompress each file
  for zst_file in $zst_files; do
    echo "Uncompressing: $zst_file"
  
    # Uncompress the zst file
    if [ -f "${zst_file%.zst}.json" ]; then
      echo "${zst_file%.zst}.json exists"
    else
      zstd -d "$zst_file" -o "${zst_file%.zst}.json" --long=31
    fi
    if [ -f "${zst_file%.zst}.csv" ]; then
      echo "${zst_file%.zst}.csv exists"
	  else
      ./flatten_comment_json/flatten_reddit_json "${zst_file%.zst}"
    fi
	done
}

# Start the process from the target directory
uncompress_recursive "$target_dir"

echo "All zst files uncompressed."
