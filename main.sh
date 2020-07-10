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

apply_pallet() {
    pal="${1}"
    trg="$2"

    {
        echo "$pal"
        echo "__ENDPALLET"
        cat "$trg"
    } | awk '
        {print}
    '
}

print_term() {
    pal="${1}"

    [ -n "$TMUX" ] && {

        esc='\033Ptmux;\033\033]'
        cesc='\007\033\\'
    }

    : "${esc=\033]}"
    : "${cesc=\007}"

    clr="$2"
    echo "$pal" | grep "$clr" | awk -v esc="${esc}" -v cesc="${cesc}" '
        $1 == "bg" {
            printf "%s11;%s%s\n", esc, $2, cesc
            printf "%s4;0;%s%s\n", esc, $2, cesc
        }
        $1 == "bg_alt" {
            printf "%s;8;%s%s\n", esc, $2, cesc
        }
        $1 == "primay" {
            printf "%s4;1;%s%s\n", esc, $2, cesc
            printf "%s4;9;%s%s\n", esc, $2, cesc
        }
        $1 == "secondary" {
            printf "%s4;2;%s%s\n", esc, $2, cesc
            printf "%s4;10;%s%s\n", esc, $2, cesc
        }
        $1 == "alert" {
            printf "%s4;3;%s%s\n", esc, $2, cesc
            printf "%s4;11;%s%s\n", esc, $2, cesc
        }
        $1 == "cursor" {
            printf "%s4;4;%s%s\n", esc, $2, cesc
            printf "%s4;12;%s%s\n", esc, $2, cesc
            printf "%s12;%s%s\n", esc, $2, cesc
        }
        $1 == "fill" {
            printf "%s4;5;%s%s\n", esc, $2, cesc
            printf "%s4;6;%s%s\n", esc, $2, cesc
            printf "%s4;13;%s%s\n", esc, $2, cesc
            printf "%s4;14;%s%s\n", esc, $2, cesc
        }
        $1 == "fg" {
            printf "%s10;%s%s\n", esc, $2, cesc
            printf "%s4;7;%s%s\n",esc,  $2, cesc
        }
        $1 == "fg_alt" {
            printf "%s17;%s%s\n", esc, $2, cesc
            printf "%s4;15;%s%s\n", esc, $2, cesc
        }
    '
}

main() {
    pallet=$(parse_pallet "$PALLET_PATH/$PALLET_FILE")

    print_term "$pallet"
}

main "${@}"
