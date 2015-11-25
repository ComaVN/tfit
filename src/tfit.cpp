/*
 * main.cpp
 * Main function for TFIT executable
 */

#include <getopt.h>
#include <iostream>
#include "cmdline.h"

int main(int argc, char *argv[]) {
    gengetopt_args_info args_info;

    if (cmdline_parser(argc, argv, &args_info) != 0) {
        exit(1);
    }
}
