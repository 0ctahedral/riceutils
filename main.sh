#!/bin/sh
#PALLET_PATH=${PALLET_PATH:-$HOME/dots/colors}
PALLET_PATH=${PALLET_PATH:-$HOME/code/ricemgr}
PALLET_FILE=${pallet:-source}

usage() {
	cat <<- ///
	ricemgr
	color pallets applied and stuff
	///

	exit
}

errx() { >&2 echo "$1"; exit 1; }
err() { >&2 echo "$1"; exit 1; }

parse_pallet() {
    file="$1"

    awk '
        /^$/ || /^#/ {next}
        match($0,/^(\w+)\s*:\s*"(#\w+)"$/,ma) {
            pallet[ma[1]]=ma[2]
        }

        END {
            for (v in pallet) {
                printf "%-20s%s\n", v, pallet[v]
            }
        }

    ' "$file"
}

main() {
    pallet=$(parse_pallet "$PALLET_PATH/$PALLET_FILE")
    echo "$pallet"
}

main "${@}"
