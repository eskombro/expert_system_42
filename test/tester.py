#!/usr/bin/env python2
import os
import subprocess

path = './tests/'
path_results = './test_results/'
sufix = '_res.test'

files = []
for r, d, f in os.walk(path):
    for file in f:
        if not '.py' in file:
            files.append(file)

for f in files:
    print("launching file: {}".format(f))
    a = subprocess.Popen(['go', 'run', 'expert_system', '-f', path+f],
            stdout = subprocess.PIPE,
            stderr=subprocess.STDOUT)
    stdout,stderr = a.communicate()
    file=open(path_results+f+sufix, "w")
    file.write(stdout)
    file.close()

    b = subprocess.Popen(['diff', path_results+f, path_results+f+sufix],
            stdout = subprocess.PIPE,
            stderr=subprocess.STDOUT)
    stdout,stderr = b.communicate()
    if (stdout == ''):
        print("OK")
    else:
        print("There is a diff:")
        print(stdout)

for f in files:
    c = subprocess.Popen(['rm', '-rf', path_results+f+sufix])
