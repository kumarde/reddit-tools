import csv

dup = set()

def count_rows(csv_file):
    row_count = 0
    with open(csv_file, newline='', encoding='utf-8') as file:  # Use newline='' to handle embedded newlines in fields
        reader = csv.reader(file)
        for row in reader:
            row_count += 1  # Count rows
            dup.add(row[12])
    print(len(dup))
    return row_count

# Usage
csv_file = "flatten_reddit_json/out/comments.csv"  # Replace with your CSV file path
total_rows = count_rows(csv_file)
print(f"The file has {total_rows} rows.")