# This is a simple script to import data from a JSON file into a SQLite database
# https://imudatascience.medium.com/importing-data-into-sqlite-via-python-f248cc23ebc2
import sys
import json
import sqlite3
import argparse

if __name__ == '__main__':
    parser = argparse.ArgumentParser(
        description='Import JSON data into SQLite')
    parser.add_argument('-f', '--filepath',
                        help='Path to JSON file', required=True)
    parser.add_argument(
        '-o', '--output', help='Path to SQLite database', required=True)

    args = parser.parse_args()

    if not args.filepath:
        print("Please provide a filepath")
        sys.exit(1)

    if not args.output:
        print("Please provide a database path")
        sys.exit(1)

    conn = sqlite3.connect(args.output)
    characters = json.load(open(args.filepath, 'r'))

    columns = []
    column = []

    for data in characters:
        for key in data.keys():
            if key not in columns:
                columns.append(key)

    # print(columns)
    # print(len(columns))

    values = []
    value = []

    for data in characters:
        for key in data.keys():
            value.append(data[key])
        values.append(value)
        value = []

    # print(len(values)) # test example has 25 users

    create_query = 'CREATE TABLE IF NOT EXISTS characters (id, name, blood, species, patronus, born, quote, imgUrl)'
    insert_query = 'INSERT INTO characters (id, name, blood, species, patronus, born, quote, imgUrl) VALUES (?, ?, ?, ?, ?, ?, ?, ?)'

    c = conn.cursor()
    c.execute(create_query)
    c.executemany(insert_query, values)
    conn.commit()
    c.close()

    print("Done")
