#!/usr/bin/python3

# http://www.cyberciti.biz/faq/python-execute-unix-linux-command-examples/

import subprocess

print("pass the argument")
subprocess.call(["ls", "-l", "/etc/resolv.conf"])

print("\nstore output to the output variable")
p = subprocess.Popen("date", stdout=subprocess.PIPE, shell=True)
(output, err) = p.communicate()
print("Today is", output)

# Good example for storing output
print("\nanother example (passing command line args)")
p = subprocess.Popen(["ls", "-l", "/etc/resolv.conf"], stdout=subprocess.PIPE)
output, err = p.communicate()
print("*** Running ls -l command ***\n", output)

print("\nto execute shell commands with pipe")
p=subprocess.getoutput("date | cut -d' ' -f4")
print(p)

# http://stackoverflow.com/questions/4368818/ 
#p=subprocess.call("date | cut -d' ' -f4")
# No such file or directory
p=subprocess.call("date | cut -d' ' -f4", shell=True)
print(p)
#subprocess.getoutput("date | cut -d' ' -f4 | xclip -i")
# This will block the program from finishing
subprocess.call("date | cut -d' ' -f4 | xclip -i", shell=True)
# This enable the program to finish correctly

# http://stackoverflow.com/questions/7353054/
task = subprocess.Popen("date | cut -d' ' -f4", shell=True, stdout=subprocess.PIPE)
data = task.stdout.read()
assert task.wait() == 0
print(data)

"""
Note that this does not capture stderr. And if you want to capture stderr as well, you'll need to use task.communicate(); calling task.stdout.read() and then task.stderr.read() can deadlock if the buffer for stderr fills. If you want them combined, you should be able to use 2>&1 as part of the shell command.
"""

print("end")

