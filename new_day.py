#!/usr/bin/env python3
import argparse
import shutil
from pathlib import Path

template = 'day00'

def create_new_day(new_day):
    if new_day is None or new_day == '':
        print('new_day field required.')
        return

    source_dir = Path(template)
    destination_dir = Path(new_day)

    shutil.copytree(source_dir, destination_dir)

    for path in destination_dir.iterdir():
        if path.is_file() and template in path.name:
            # update package name
            content = path.read_text()
            modified_content = content.replace(template, new_day)
            path.write_text(modified_content)

            # update file name
            new_name = path.name.replace(template, new_day)
            path.rename(destination_dir/new_name)
    
    # create empty data files
    (Path(destination_dir) / 'sample').touch()
    (Path(destination_dir) / 'input').touch()


def main():
    parser = argparse.ArgumentParser(description='setup a new day for advent of code')
    parser.add_argument('new_day', type=str, help='name of the new puzzle')
    args = parser.parse_args()
    create_new_day(args.new_day)


if __name__ == '__main__':
    main()