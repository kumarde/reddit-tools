import json
import sys

def main():
    comment_f = sys.argv[1]
    count = 0
    total = 0
    for l in open(comment_f, 'r'):
        l = l.strip()
        data = json.loads(l)
        if data['toxicity']['attributeScores'] == None:
            count += 1
        total += 1
    print("Null frac from comments:")
    print(count, total, float(count)/total)

    thread_f = sys.argv[2]
    count = 0
    total = 0
    for l in open(thread_f, 'r'):
        l = l.strip()
        data = json.loads(l)
        if data['comments'] == None:
            continue
        for comment in data['comments']:
            if comment['toxicity']['attributeScores'] == None:
                count += 1
            total += 1
    print("Null frac from threads:")
    print(count, total, float(count)/total)

if __name__ == "__main__":
    main()
