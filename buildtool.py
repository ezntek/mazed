#!/usr/bin/python3
import sys, subprocess

ARCH = "linux-amd64"
BUILDCMD = f"go build -o out/{ARCH}/out src/main.go"
RUNCMD = f"go run src/main.go"
EXECCMD = f"out/{ARCH}/out"

res = []

RESPONSE_LOOKUP = {
    1 : "error",
    0 : "success"
}


def main() -> int:
    for task in sys.argv:
        if task == "build":
            res.append(subprocess.call(BUILDCMD, shell=True))
        if task == "run":
            res.append(subprocess.call(RUNCMD, shell=True))
        if task == "exec":
            res.append(subprocess.call(EXECCMD, shell=True))
            
    for count, response in enumerate(res):
        print(f"task {count+1}: {RESPONSE_LOOKUP[response]}")
    return 0

if __name__ == "__main__":
    sys.exit(main())


