#!/usr/bin/env python

import os
import subprocess


README_TEMPLATE = """\
A place for experiments, learning, and random pieces of code.

{progress}"""

PROGRESS_TEMPLATE = (
    "[![{name}](http://img.shields.io/badge/"
    "{escaped}-{percent}%25_({complete}/{total})-{color}.svg?style=flat"
    ")](https://github.com/jasonkeene/playground/blob/master/{name}/"
    "progress.md)  \n"
)

REPO_ROOT = os.path.dirname(os.path.realpath(__file__))


def badge_color(percent):
    """Return sheild.io badge color."""
    if percent < 16.6:
        return 'red'
    elif percent < 33.2:
        return 'orange'
    elif percent < 50:
        return 'yellow'
    elif percent < 66.4:
        return 'yellowgreen'
    elif percent < 83.0:
        return 'green'
    else:
        return 'brightgreen'


def escape_shield_name(name):
    """Return the shield name escaped for shield.io URL."""
    return name.replace('-', '--')


def get_progres_status(name):
    """Return the percent, complete, and total progress."""
    with open(os.path.join(REPO_ROOT, name, 'progress.md')) as f:
        data = f.read()
    uncomplete = data.count('- [ ]')
    complete = data.count('- [x]')
    total = complete + uncomplete
    percent = float(complete) / total * 100
    return percent, complete, total


def generate_progress():
    """Return a string of progress information."""
    progress_dirs = find_progress_dirs()
    progress = ''
    for name in progress_dirs:
        percent, complete, total = get_progres_status(name)
        progress += PROGRESS_TEMPLATE.format(
            name=name,
            escaped=escape_shield_name(name),
            percent='{:.2f}'.format(percent).rstrip('0.') or '0',
            complete=complete,
            total=total,
            color=badge_color(percent),
        )
    return progress.strip()


def order_progress_dir(name):
    path = os.path.join(REPO_ROOT, name, 'progress.md')
    command = 'git log --format="%at" -- {}'.format(path)
    output = subprocess.check_output(command, shell=True)
    if output:
        return max(output.split('\n'))
    return '999999999999999'


def find_progress_dirs():
    """Find all top level directories that contain progress files."""
    progress_dirs = []
    for node in os.listdir(REPO_ROOT):
        if os.path.isdir(node):
            if 'progress.md' in os.listdir(node):
                progress_dirs.append(node)
    return reversed(sorted(progress_dirs, key=order_progress_dir))


def main():
    print README_TEMPLATE.format(progress=generate_progress())


if __name__ == '__main__':
    main()
