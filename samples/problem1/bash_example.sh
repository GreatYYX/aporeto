#!/bin/bash

CONTENT="\
Alabama
Alaska
Arizona
Arkansas
California
Colorado
Connecticut
Delaware
Florida
Georgia
Hawaii
Idaho
Illinois
Indiana
Iowa
Kansas
Kentucky
Louisiana
Maine
Maryland
Massachusetts
Michigan
Minnesota
Mississippi
Missouri
Montana
Nebraska
Nevada
New Hampshire
New Jersey
New Mexico
New York
North Carolina
North Dakota
Ohio
Oklahoma
Oregon
Pennsylvania
Rhode Island
South Carolina
South Dakota
Tennessee
Texas
Utah
Vermont
Virginia
Washington
West Virginia
Wisconsin
Wyoming"

FILE_NAME=""
NO_PROMPT=false
VERBOSE=false

usage() {
	echo "Usage:"
	echo "./bash_example [--help|-h]"
	echo "./bash_example --create-file=<filename> [--no-prompt] [--verbose]"
	exit $1
}

OPT_PATTERN=`getopt -o h --long help,create-file:,no-prompt,verbose -n 'bash_example.sh' -- "$@"`
if [ $? -ne 0 ]; then usage 1; fi # invalid arguments
eval set -- "$OPT_PATTERN"
while true; do
    case "$1" in
		-h|--help)
			usage 0 ;;
		--create-file)
			FILE_NAME="$2"; shift 2 ;;
		--no-prompt)
			NO_PROMPT=true; shift ;;
		--verbose)
			VERBOSE=true; shift ;;
		--)
			shift; break ;;
        *)
            usage 1 ;;
    esac
done

if [ -z "$FILE_NAME" ]; then usage 1; fi
if [ -f "$FILE_NAME" ]; then
	# file exists
	if $VERBOSE; then echo "File already exists"; fi
	if ! $NO_PROMPT; then
		while true; do
			read -p "File exists. Overwrite (y/n) ?" yn
			case $yn in
				[Yy]*)
					break ;;
				[Nn]*)
					exit 1 ;;
				*) ;;
			esac
		done
	fi
	# remove file
	rm -rf $FILE_NAME
	if [ $? -ne 0 ]; then exit 1; fi
	if $VERBOSE; then echo "File removed"; fi
fi

# new file
echo -e "$CONTENT" > $FILE_NAME
if [ $? -ne 0 ]; then exit 1; fi
if $VERBOSE; then echo "File created"; fi
exit 0